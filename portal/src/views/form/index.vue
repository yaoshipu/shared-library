<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Book name">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="Author">
        <el-input v-model="form.author" />
      </el-form-item>
      <el-form-item label="Publisher">
        <el-input v-model="form.publisher" />
      </el-form-item>
      <el-form-item label="Donator">
        <el-input v-model="form.donator" />
      </el-form-item>
      <el-form-item label="Image URL">
        <el-input v-model="form.image" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="saveBook">Donate</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { donateBook } from "@/api/book";

export default {
  data() {
    return {
      form: {
        name: "",
        author: "",
        publisher: "",
        donator: "",
        image: "",
      },
    };
  },
  methods: {
    onSubmit() {
      this.$message("submit!");
    },
    onCancel() {
      this.$message({
        message: "cancel!",
        type: "warning",
      });
    },
    saveBook() {
      donateBook(this.form).then((response) => {
        console.log(response.data.items);
      });
    },
  },
};
</script>

<style scoped>
.line {
  text-align: center;
}
</style>

