<template>
  <div class="dashboard-container">
    <div class="search-wrapper">
      <input type="text" v-model="search" placeholder="Search book name.." />
      <label>Search book name:</label>
    </div>
    <el-row :gutter="20">
      <el-col v-for="(item,index) in filteredList" :key="index" :span="4">
        <el-card :body-style="{ padding: '0px' }">
          <el-image style="width: 200px; height: 300px" :src="item.image_url"></el-image>
          <div style="padding: 14px;">
            <div class="clearfix">
              <small class="name">{{ item.name }}</small>
            </div>
            <div class="clearfix">
              <small class="author">by {{ item.author }}</small>
            </div>
            <div class="bottom clearfix">
              <small style="text-align: left;">donator {{ item.donator }}</small>
              <!-- <el-button type="text" class="button" icon="el-icon-plus"></el-button> -->
              <!-- 绑定到设置好的 visible -->
              <el-popover
                placement="top"
                trigger="click"
                width="100"
                v-model="item.visible"
                v-if="item.current_borrower ===''"
              >
                <span>Borrow the book?</span>
                <div style="text-align: right; margin-top: 10px">
                  <el-button size="mini" type="text" @click="cancelBorrow(index)">Cancel</el-button>
                  <el-button type="primary" size="mini" @click="confirmBorrow(index)">Yes</el-button>
                </div>
                <el-button slot="reference" class="button" icon="el-icon-plus"></el-button>
              </el-popover>
              <el-tooltip v-else class="item" effect="dark" :content="`Borrowed by ${ item.current_borrower}`" placement="top-start">
                <el-button class="button" icon="el-icon-warning"></el-button>
              </el-tooltip>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { getBooks } from "@/api/book";
import { borrowBook } from "@/api/book";

export default {
  name: "Dashboard",
  data() {
    return {
      list: [],
      dialogFormVisible: false,
      search: "",
      selectBook: {
        name: "",
        author: "",
        press: "",
        contributor: "",
      },
    };
  },
  computed: {
    ...mapGetters(["name"]),
    filteredList() {
      return this.list.filter((item) => {
        return item.name.toLowerCase().includes(this.search.toLowerCase());
      });
    },
  },
  created() {
    this.fetchData();
  },
  methods: {
    cancelBorrow(index) {
      this.$set(this.list[index], "visible", false)
    },
    confirmBorrow(index) {
      this.$set(this.list[index], "visible", false)
      let borrowOpt = {
        id: this.list[index].id,
        name:this.list[index].name,
        borrower: "spark",  // hard code user name
      };

      borrowBook(borrowOpt).then((response) => {
        this.$message({
          message: `${borrowOpt.name} 借阅成功`,
          type: 'success'
        });
        this.fetchData()
      });
    },
    fetchData() {
      getBooks().then((response) => {
        this.list = response;
        this.list.map((element) => {
          this.$set(element, "visible", false)
        });
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 10px;
  }
}
.el-col {
  border-radius: 4px;
}
.el-card {
  margin-bottom: 20px;
}
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
.name {
  display: inline-block; /*Display inline and maintain block characteristics.*/
  vertical-align: top; /*Makes sure all the divs are correctly aligned.*/
  white-space: normal; /*Prevents child elements from inheriting nowrap.*/
}

.time {
  font-size: 13px;
  color: #999;
}
.bottom {
  margin-top: 13px;
  line-height: 12px;
}
.button {
  padding: 0;
  float: right;
}
.image {
  width: 100%;
  display: block;
}
.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}

.search-wrapper {
  position: relative;
  margin-bottom: 10px;
  label {
    position: absolute;
    font-size: 12px;
    color: rgba(0, 0, 0, 0.5);
    top: 8px;
    left: 12px;
    z-index: -1;
    transition: 0.15s all ease-in-out;
  }
  input {
    padding: 4px 12px;
    color: rgba(0, 0, 0, 0.7);
    border: 1px solid rgba(0, 0, 0, 0.12);
    transition: 0.15s all ease-in-out;
    background: white;
    &:focus {
      outline: none;
      transform: scale(1.05);
      & + label {
        font-size: 10px;
        transform: translateY(-24px) translateX(-12px);
      }
    }
    &::-webkit-input-placeholder {
      font-size: 12px;
      color: rgba(0, 0, 0, 0.5);
      font-weight: 100;
    }
  }
}
</style>
