<template>
  <div class="dashboard-container">
    <el-dialog title="收货地址"
               :visible.sync="dialogFormVisible">
      <el-form :model="selectBook">
        <el-form-item label="书名">
          <el-input v-model="selectBook.name"
                    autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="作者">
          <el-input v-model="selectBook.author"
                    autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer"
           class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary"
                   @click="saveBook">确 定</el-button>
      </div>
    </el-dialog>
    <el-row :gutter="50">
      <el-col v-for="(item,index) in list"
              :key="index"
              :span="6">
        <div class="book-container">
          <div class="pic">
          </div>
          <div class="info">
            <h3 class="name">{{item.name}}</h3>
            <span class="author">{{item.author}}</span>
            <span class="press">{{item.press}}</span>
            <span class="contributor">{{item.contributor}}</span>
            <el-button @click="editBook(item)"
                       type="primary"
                       size="small"
                       round>+</el-button>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getList } from '@/api/table'
import { contributeBook } from '@/api/book'

export default {
  name: 'Dashboard',
  data() {
    return {
      list: null,
      dialogFormVisible: false,
      selectBook: {
        name: '',
        author: '',
        press: '',
        isbn: '',
        contributor: ''
      }
    }
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      // let response = [{
      //   id: 1,
      //   name: 'test',
      //   author: 'abc',
      //   press: 'xxx',
      //   isbn: 'xxx',
      //   contributor: ''
      // }]
      // for (let index = 0; index < 20; index++) {
      //   response.push({
      //     id: 1,
      //     name: 'test',
      //     author: 'abc',
      //     press: 'xxx',
      //     isbn: 'xxx',
      //     contributor: ''
      //   })
      // }
      // this.list = response
      getList().then(response => {
        this.list = response.data.items
        console.log(response.data.items)
      })
    },
    editBook(item) {
      console.log(item);
      this.dialogFormVisible = true;
      this.selectBook = item;
    },
    saveBook() {
      const payload = this.selectBook;
      contributeBook(payload).then(response => {
        console.log(response.data.items)
      })
    },
    contribute() {
      const payload = {
        name: 'test',
        author: 'abc',
        press: 'xxx',
        isbn: 'xxx',
        contributor: ''
      }
      contributeBook(payload).then(response => {
        console.log(response.data.items)
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
.book-container {
  width: 150px;
  height: 200px;
  background: #d3dce6;
}
</style>
