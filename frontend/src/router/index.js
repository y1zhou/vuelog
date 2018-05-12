import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

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
      path: '/blog',
      name: 'Blog',
      component: HelloWorld
    },
    {
      path: '/about',
      name: 'About',
      component: HelloWorld
    }
  ]
})
