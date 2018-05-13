import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Page404 from '@/components/Page404'
import Login from '@/components/Login'
import Search from '@/components/Search'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HelloWorld
    },
    {
      path: '/tags',
      name: 'Tags',
      component: HelloWorld
    },
    {
      path: '/about',
      name: 'About',
      component: HelloWorld
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/search',
      name: 'Search',
      component: Search
    },
    {
      path: '/404',
      name: 'Page404',
      component: Page404
    }
  ]
})
