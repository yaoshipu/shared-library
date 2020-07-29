package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2" // imports as package "cli"
)

var (
	// build number set at compile-time
	build = "0"

	mgoConfig *MgoConfig

	bookColl *BookColl
)

func main() {
	app := cli.NewApp()
	app.Name = "library"
	app.Usage = "shared library"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "library service port number",
			EnvVars: []string{"PORT"},
			Value:   "20987",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	// Echo instance
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	logFile, err := os.OpenFile("library.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: "${time_rfc3339_nano},${remote_ip},${method},${uri},${status},${latency_human},${error}\n",
	}))

	// Routes
	e.GET("/ping", ping)

	books := e.Group("/api/books")
	{
		books.GET("", listBooks)
		books.GET("/:username", myBooks)
		books.POST("", createBook)
	}

	user := e.Group("/api/user")
	{
		user.POST("/login", userLogin)
		user.GET("/info", userInfo)
	}

	if err := initService(); err != nil {
		return err
	}

	// Start server
	return e.Start(fmt.Sprintf(":%s", c.String("port")))
}

func initService() error {
	var err error
	mgoConfig = &MgoConfig{Host: "40.112.174.85:27017", DB: "sharedlibrary_test", Coll: "book", Mode: "strong"}
	bookColl, err = NewBookColl(mgoConfig)
	if err != nil {
		return err
	}
	return nil
}
