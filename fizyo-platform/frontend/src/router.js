import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import List from '@/views/List.vue'
import Detail from '@/views/Detail.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/list/:city',
    name: 'List',
    component: List
  },
  {
    path: '/detail/:id',
    name: 'Detail',
    component: Detail
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
