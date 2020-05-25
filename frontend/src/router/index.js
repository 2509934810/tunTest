import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

const createRouter = [
  {
    path: '/',
    name: 'index',
    component: () => import('@/views/index.vue')
  }
]

const router =  new Router({
  routes: createRouter
})

export default router

