<template>
  <div class="message">
    <el-card class="box-card" shadow="hover">
      <div slot="header">
        <h2>404 Not Found</h2>
      </div>
      <div> {{ msg }} </div>
      <el-input prefix-icon="el-icon-search" clearable v-model="searchContent" class="searchbar">
      </el-input>
    </el-card>
  </div>
</template>

<script>
/* eslint-disable */
export default {
  name: 'Page404',
  data() {
    return {
      msg:
        "The requested URL wasn't found on this server. Try searching for it maybe?",
      searchContent: ''
    }
  }
}
</script>

<style scoped>
.message {
  padding: 5rem 10rem;
  display: flex;
  align-items: center;
  justify-content: center;
}
.box-card {
  width: 40rem;
  text-align: center;
  font-family: Raleway;
}
.searchbar {
  width: 20rem;
  margin-top: 2rem;
}
</style>
