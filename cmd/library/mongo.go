package main

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/mgo.v2"
)

var (
	// CopySessionMaxRetry ...
	CopySessionMaxRetry = 5
)

// Safe ...
type Safe struct {
	W        int    `json:"w"`
	WMode    string `json:"wmode"`
	WTimeout int    `json:"wtimeoutms"`
	FSync    bool   `json:"fsync"`
	J        bool   `json:"j"`
}

// MgoConfig ...
type MgoConfig struct {
	Host           string `json:"host"`
	DB             string `json:"db"`
	Coll           string `json:"coll"`
	Mode           string `json:"mode"`
	SyncTimeoutInS int64  `json:"timeout"` // 以秒为单位
	Safe           *Safe  `json:"safe"`
}

// Session ...
type Session struct {
	*mgo.Session
	DB   *mgo.Database
	Coll *mgo.Collection
}

// Open ...
func Open(cfg *MgoConfig) *Session {

	session, _ := mgo.Dial(cfg.Host)
	EnsureSafe(session, cfg.Safe)
	db := session.DB(cfg.DB)
	c := db.C(cfg.Coll)

	return &Session{session, db, c}
}

// IsSessionClosed test whether session closed
//
// PS: sometimes it's not corrected
func IsSessionClosed(s *mgo.Session) (res bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("[MGO2_IS_SESSION_CLOSED] check session closed panic:", err)
		}
	}()
	res = true
	return s.Ping() != nil
}

func checkSession(s *mgo.Session) (err error) {
	return s.Ping()
}

func isServersFailed(err error) bool {
	return strings.Contains(err.Error(), "no reachable servers")
}

// CopySession ...
func CopySession(s *mgo.Session) *mgo.Session {
	for i := 0; i < CopySessionMaxRetry; i++ {
		res := s.Copy()
		err := checkSession(res)
		if err == nil {
			return res
		}
		CloseSession(res)
		log.Fatal("[MGO2_COPY_SESSION] copy session and check failed:", err)
		if isServersFailed(err) {
			panic("[MGO2_COPY_SESSION_FAILED] servers failed")
		}
	}
	msg := fmt.Sprintf("[MGO2_COPY_SESSION_FAILED] failed after %d retries", CopySessionMaxRetry)
	log.Fatal(msg)
	panic(msg)
}

// FastCopySession ...
func FastCopySession(s *mgo.Session) *mgo.Session {
	return s.Copy()
}

// CloseSession ...
func CloseSession(s *mgo.Session) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("[MGO2_CLOSE_SESSION_RECOVER] close session panic", err)
		}
	}()
	s.Close()
}

// CopyDatabase copy database's session, and re-create the database.
//
// you need call `CloseDatabase` after use this
func CopyDatabase(db *mgo.Database) *mgo.Database {
	return CopySession(db.Session).DB(db.Name)
}

// FastCopyDatabase ...
func FastCopyDatabase(db *mgo.Database) *mgo.Database {
	return FastCopySession(db.Session).DB(db.Name)
}

// CloseDatbase close the session of the database
func CloseDatbase(db *mgo.Database) {
	CloseSession(db.Session)
}

// CopyCollection copy collection's session, and re-create the collection
//
// you need call `CloseColletion` after use this
func CopyCollection(c *mgo.Collection) *mgo.Collection {
	return CopyDatabase(c.Database).C(c.Name)
}

// FastCopyCollection ...
func FastCopyCollection(c *mgo.Collection) *mgo.Collection {
	return FastCopyDatabase(c.Database).C(c.Name)
}

// CloseCollection close the session of the collection
func CloseCollection(c *mgo.Collection) {
	CloseDatbase(c.Database)
}

// CheckIndex ...
func CheckIndex(c *mgo.Collection, key []string, unique bool) error {
	originIndexs, err := c.Indexes()
	if err != nil {
		return fmt.Errorf("<CheckIndex> get indexes: %v", err)
	}
	for _, index := range originIndexs {
		if checkIndexKey(index.Key, key) && unique == index.Unique {
			return nil
		}
	}
	return fmt.Errorf("<CheckIndex> not found index: %v unique: %v", key, unique)
}

func checkIndexKey(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, k := range a {
		if k != b[i] {
			return false
		}
	}
	return true
}

// EnsureSafe W 和 WMode 只在 replset 模式下生效，非replset不能配置，否则会出错
// WMode只在2.0版本以上才生效
func EnsureSafe(session *mgo.Session, safe *Safe) {
	if safe == nil {
		return
	}
	session.EnsureSafe(&mgo.Safe{
		W:        safe.W,
		WMode:    safe.WMode,
		WTimeout: safe.WTimeout,
		FSync:    safe.FSync,
		J:        safe.J,
	})
}

// NewSession ...
func NewSession(cfg *MgoConfig, collection string, indexes []string) (*Session, error) {
	cfg.Coll = collection
	session := Open(cfg)
	if indexes != nil {
		err := session.Coll.EnsureIndex(mgo.Index{Key: indexes, Unique: true})
		if err != nil {
			return nil, fmt.Errorf("%s EnsureIndex error: %v", collection, err)
		}
	}
	return session, nil
}
