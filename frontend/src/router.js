import Vue from 'vue'
import Router from 'vue-router'
const Home = () => import('./views/Home')
const About = () => import('./views/About')
const Login = () => import('./views/Login')
const Search = () => import('./views/Search')
const Page404 = () => import('./views/Page404')

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/about',
      name: 'About',
      component: About
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
      path: '*',
      component: Page404
    }
  ]
})
