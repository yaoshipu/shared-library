<template>
  <el-row>
    <el-col :span="12">
      <div class="header">My donated books</div>
      <div class="wrapper">
        <div class="card" v-for="(item,index) in mybooks.donated" :key="index">
          <a v-bind:href="item.link" target="_blank">
            <img v-bind:src="item.image_url" />
            <small>by {{ item.author }}</small>
          </a>
        </div>
      </div>
    </el-col>
    <el-col :span="12">
      <div class="header">My borrowed books</div>
      <div class="wrapper">
        <div class="card" v-for="(item,index) in mybooks.borrowed" :key="index">
          <a>
            <img v-bind:src="item.image_url" />
            <small class="author-name">by {{ item.author }}</small>
            <el-popover placement="top" trigger="click" width="250" v-model="item.visible">
              <span>Do you want to return the book?</span>
              <div style="text-align: right; margin: 0">
                <el-button size="mini" type="text" @click="cancelReturn(index)">Cancel</el-button>
                <el-button type="primary" size="mini" @click="confirmReturn(index)">Yes</el-button>
              </div>
              <el-button slot="reference" size="mini" class="button" icon="el-icon-minus" circle></el-button>
            </el-popover>
          </a>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import { mapGetters } from "vuex";
import { getMyBooks } from "@/api/book";
import { returnBook } from "@/api/book";

export default {
  name: "MyBooks",
  data() {
    return {
      visible: false, // return pophover visibility

      mybooks: {
        donated: [],
        borrowed: [],
      },
      dialogFormVisible: false,
      search: "",
    };
  },
  computed: {
    ...mapGetters(["name"]),
  },
  created() {
    this.fetchData();
  },
  methods: {
    cancelReturn(index) {
      this.$set(this.mybooks.borrowed[index], "visible", false);
    },
    confirmReturn(index) {
      this.$set(this.mybooks.borrowed[index], "visible", false);
      let returnOpt = {
        id: this.mybooks.borrowed[index].id,
        name:this.mybooks.borrowed[index].name,
        borrower: "spark", // hard code user name
      };
      returnBook(returnOpt).then((response) => {
        this.$message({
          message: `${returnOpt.name} 归还成功`,
          type: 'success'
        });
        this.fetchData();
      });
    },
    fetchData() {
      getMyBooks("spark").then((response) => {
        this.mybooks = response;
      });
    },
  },
};
</script>

<style lang="scss" scoped>
html,
body {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  margin-top: 16px;
  margin-bottom: 16px;
}

div#app {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  .wrapper {
    display: flex;
    max-width: 444px;
    flex-wrap: wrap;
    padding-top: 12px;
  }

  .card {
    box-shadow: rgba(0, 0, 0, 0.117647) 0px 1px 6px,
      rgba(0, 0, 0, 0.117647) 0px 1px 4px;
    width: 124px;
    margin: 12px;
    transition: 0.15s all ease-in-out;
    &:hover {
      transform: scale(1.1);
    }
    a {
      text-decoration: none;
      padding: 12px;
      color: #03a9f4;
      font-size: 24px;
      display: flex;
      flex-direction: column;
      align-items: center;
      img {
        height: 100px;
      }
      small {
        font-size: 10px;
        padding: 4px;
        &.author-name{
          height: 30px;
        }
      }
    }
  }

  .hotpink {
    background: hotpink;
  }

  .green {
    background: green;
  }

  .box {
    width: 100px;
    height: 100px;
    border: 1px solid rgba(0, 0, 0, 0.12);
  }
}

.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #99a9bf;
}
.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}

.header {
  margin-top: 20px;
  margin-left: 20px;
}
</style>
