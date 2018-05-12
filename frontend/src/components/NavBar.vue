<template>
  <el-row :class="navbarClass">
    <el-col :span="1" :offset="5" class="navbar-col">
      <router-link to="/">
        <img src="@/assets/logo-small.png" alt="logo" class="logo">
      </router-link>
    </el-col>
    <el-col :span="13" pull="5" class="navbar-col _right">
      <el-menu mode="horizontal" :router="true" :default-active="activeIndex" class="el-menu">
        <el-menu-item index="/search">
          <i class="el-icon-search" />
        </el-menu-item>
        <el-menu-item index="/login">Login</el-menu-item>
        <el-menu-item index="/about">About</el-menu-item>
        <el-menu-item index="/blog">Blog</el-menu-item>
      </el-menu>
    </el-col>
  </el-row>
</template>

<script>
/* eslint-disable */
export default {
  data() {
    return {
      isScrolled: false,
      searchInput: ''
    }
  },
  computed: {
    activeIndex() {
      if (this.$route.matched.length > 1) {
        return
      }
      return '/' + this.$route.matched[0].path.split('/')[1]
    },
    navbarClass() {
      return {
        navbar: true,
        scrolled: this.isScrolled
      }
    }
  },
  methods: {
    updateScroll() {
      this.isScrolled = window.scrollY > 0
    }
  },
  created() {
    window.addEventListener('scroll', this.updateScroll)
  },
  destroy() {
    window.removeEventListener('scroll', this.updateScroll)
  }
}
</script>

<style lang="scss" scoped>
$kinda-white: rgba(255, 255, 255, 0.8);
$almost-black: rgba(0, 0, 0, 0.15);
.navbar {
  position: fixed;
  width: 100%;
  top: 0;
  z-index: 900;
  background-color: $kinda-white;
  transition: all 0.15s linear;
  &.scrolled {
    box-shadow: 0 0.05rem 0.25rem $almost-black !important;
  }
  height: 3rem;
}
.navbar-col {
  height: inherit;
  line-height: 3rem;
  /deep/ & .el-input__inner {
    height: 2rem;
    background-color: $kinda-white !important;
  }
  &._right {
    float: right;
  }
}
.logo {
  width: 3rem;
  transition: background 0.15s ease-in-out;
  &:hover {
    background: #f8f8f9;
  }
}
.navbar-col > .el-menu {
  height: inherit;
  background-color: transparent;
  border: none;
  & > .el-menu-item {
    height: inherit;
    line-height: 3rem;
    font-family: Raleway;
    float: right;
    &:hover {
      background: #f8f8f9;
    }
  }
}
</style>
