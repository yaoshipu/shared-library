<template>
  <div class="dashboard-container">
    <div class="search-wrapper">
      <input type="text" v-model="search" placeholder="Search book name.." />
      <label>Search book name:</label>
    </div>
    <el-row :gutter="20">
      <el-col v-for="(item,index) in filteredList" :key="index" :span="4">
        <el-card :body-style="{ padding: '0px' }">
          <el-image style="width: 200px; height: 300px" :src="item.image"></el-image>
          <div style="padding: 14px;">
            <div class="clearfix">
              <small class="name">{{ item.name }}</small>
            </div>
            <div class="clearfix">
              <small class="author">by {{ item.author }}</small>
            </div>
            <div class="bottom clearfix">
              <small class="contributor">donator {{ item.donator }}</small>
              <!-- <el-button type="text" class="button" icon="el-icon-plus"></el-button> -->
              <el-popover placement="top" width="160" v-model="visible">
                <p>Do you want to borrow the book?</p>
                <div style="text-align: right; margin: 0">
                  <el-button size="mini" type="text" @click="visible = false">Cancel</el-button>
                  <el-button type="primary" size="mini" @click="visible = false">Yes</el-button>
                </div>
                <el-button
                  slot="reference"
                  class="button"
                  icon="el-icon-plus"
                  @click="visible = true"
                ></el-button>
              </el-popover>
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

export default {
  name: "Dashboard",
  data() {
    return {
      visible: false,
      list: [
        {
          id: 1,
          name: "Site Reliability Engineering",
          author: "Chris Jones",
          publisher: "xxx",
          donator: "Spark",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/51XswOmuLqL._SX379_BO1,204,203,200_.jpg",
        },
        {
          id: 2,
          name: "Effective Java",
          author: "Joshua Bloch",
          publisher: "yyy",
          donator: "Steven",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/41JLgmt8MlL._SX402_BO1,204,203,200_.jpg",
        },
        {
          id: 3,
          name: "Think Python",
          author: "Allen B Downey ",
          publisher: "yyy",
          donator: "Sun",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/51-54ZrGSZL._SX379_BO1,204,203,200_.jpg",
        },
        {
          id: 4,
          name: "Design Patterns",
          author: "Erich Gamma",
          publisher: "yyy",
          donator: "Coco",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/41JLgmt8MlL._SX402_BO1,204,203,200_.jpg",
        },
        {
          id: 5,
          name: "Hacking",
          author: "Jon Erickson",
          publisher: "yyy",
          donator: "Mei",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/61VBaAS4IbL._SX383_BO1,204,203,200_.jpg",
        },
        {
          id: 6,
          name: "Introduction to Algorithms",
          author: "Thomas H. Cormen",
          publisher: "yyy",
          donator: "Max",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/41T0iBxY8FL._SX440_BO1,204,203,200_.jpg",
        },
        {
          id: 7,
          name: "Thinking in Systems",
          author: "Donella H. Meadows",
          publisher: "yyy",
          donator: "Steven",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/51frZKhRiIL._SX330_BO1,204,203,200_.jpg",
        },
        {
          id: 8,
          name: "Life 3.0",
          author: "Max Tegmark",
          publisher: "yyy",
          donator: "Spark",
          image:
            "https://images-na.ssl-images-amazon.com/images/I/41mc05UgX8L._SX329_BO1,204,203,200_.jpg",
        },
      ],
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
    //this.list;
    this.fetchData();
  },
  methods: {
    fetchData() {
      //   let response = [{
      //     id: 1,
      //     name: 'Vue.js',
      //     author: 'abc',
      //     press: 'xxx',
      //     isbn: 'xxx',
      //     contributor: 'Spark'
      //   }]
      //   for (let index = 0; index < 6; index++) {
      //     response.push({
      //       id: 1,
      //       name: 'Vue.js',
      //       author: 'abc',
      //       press: 'xxx',
      //       isbn: 'xxx',
      //       contributor: 'Steven'
      //     })
      //   }
      getBooks().then((response) => {
        console.log(response.data.items);
        this.list = response.data.items;
      });
    },
    // editBook(item) {
    //   console.log(item);
    //   this.dialogFormVisible = true;
    //   this.selectBook = item;
    // },
    // saveBook() {
    //   const payload = this.selectBook;
    //   contributeBook(payload).then((response) => {
    //     console.log(response.data.items);
    //   });
    // },
    // contribute() {
    //   const payload = {
    //     name: "test",
    //     author: "abc",
    //     press: "xxx",
    //     isbn: "xxx",
    //     contributor: "",
    //   };
    //   contributeBook(payload).then((response) => {
    //     console.log(response.data.items);
    //   });
    // },
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
