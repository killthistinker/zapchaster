import { createRouter, createWebHistory, createMemoryHistory, createWebHashHistory } from 'vue-router'
import Home from '../components/Home.vue'
import About from '../components/About.vue'
import Contacts from '../components/Contacts.vue'
import Delivery from '../components/Delivery.vue'

const routes = [
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
    path: '/contacts',
    name: 'Contacts',
    component: Contacts
  },
  {
    path: '/delivery',
    name: 'Delivery',
    component: Delivery
  }
]
const createHistory = process.env.SERVER ? createMemoryHistory : (process.env.VUE_ROUTER_MODE === 'history' ? createWebHistory : createWebHashHistory)
const router = createRouter({
  routes,
  history: createHistory(process.env.MODE === 'ssr' ? void 0 : process.env.VUE_ROUTER_BASE)
})
export default router