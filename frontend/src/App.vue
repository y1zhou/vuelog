<template>
  <el-container id="app">
    <el-header id="header">
      <nav-bar/>
    </el-header>
    <el-main>
      <img src="./assets/logo.png">
      <router-view/>
    </el-main>
  </el-container>
</template>

<script>
import NavBar from './components/NavBar'
export default {
  name: 'App',
  components: {
    NavBar
  }
}
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Raleway');
$quotemark: url('assets/quotemark.png');
html,
body {
  margin: 0;
  padding: 0;
  font-size: 1rem;
  line-height: 1.75rem;
  background-color: #f8f8f9;
  text-align: center;
}
#header {
  height: 3rem !important;
  padding: 0;
}

blockquote {
  position: relative;
  padding: 2.5rem 0 1rem 1rem;
  color: #bbbec4;
  margin: 1rem 0 0 0;
  font-size: 1rem;
  line-height: 1.25rem;
  &::before {
    position: absolute;
    top: 0;
    left: 0;
    content: '';
    background: $quotemark;
    width: 2rem;
    height: 2rem;
    background-size: contain;
    background-repeat: no-repeat;
  }
}
.mobile .katex-display > .katex {
  white-space: nowrap;
  max-width: 100%;
  overflow-x: scroll;
  overflow-y: visible;
}
</style>
