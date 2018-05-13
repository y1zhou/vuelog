<template>
    <div class="search">
        <el-card class="box-card" shadow="hover">
            <div> {{ msg }} </div>
            <el-input prefix-icon="el-icon-search" clearable v-model="searchContent" class="searchbar">
            </el-input>
        </el-card>
    </div>
</template>

<script>
/* eslint-disable */
export default {
  name: 'Search',
  data() {
    return {
      msg: 'Looking for something?',
      searchContent: ''
    }
  }
}
</script>

<style scoped>
.search {
  padding: 5rem 10rem;
  display: flex;
  align-items: center;
  justify-content: center;
}
.box-card {
  width: 30rem;
  text-align: center;
  font-family: Raleway;
}
.searchbar {
  width: 20rem;
  margin-top: 2rem;
}
</style>
