<template>
  <el-row>
    <el-col :span="12">
      <div class="header">My donated books</div>
      <div class="wrapper">
        <div class="card" v-for="(item,index) in mybooks.donated" :key="index">
          <a v-bind:href="item.link" target="_blank">
            <img v-bind:src="item.image" />
            <small>by {{ item.author }}</small>
          </a>
        </div>
      </div>
    </el-col>
    <el-col :span="12">
      <div class="header">My borrowed books</div>
      <div class="wrapper">
        <div class="card" v-for="(item,index) in mybooks.borrowed" :key="index">
          <a v-bind:href="item.link" target="_blank">
            <img v-bind:src="item.image" />
            <small>by {{ item.author }}</small>
          </a>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import { mapGetters } from "vuex";
import { contributeBook } from '@/api/book'

export default {
  name: "MyBooks",
  data() {
    return {
      mybooks: {
        donated: [
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
        ],
        borrowed: [
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
      },

      dialogFormVisible: false,
      search: "",
    };
  },
  computed: {
    ...mapGetters(["name"]),
  },
  created() {
    this.list;
  },
  methods: {
    fetchData() {
      let response = [{
        id: 1,
        name: 'Vue.js',
        author: 'abc',
        press: 'xxx',
        isbn: 'xxx',
        contributor: 'Spark'
      }]
      for (let index = 0; index < 6; index++) {
        response.push({
          id: 1,
          name: 'Vue.js',
          author: 'abc',
          press: 'xxx',
          isbn: 'xxx',
          contributor: 'Steven'
        })
      }
      this.list = response
    getList().then(response => {
      this.list = response.data.items
      console.log(response.data.items)
    })
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
