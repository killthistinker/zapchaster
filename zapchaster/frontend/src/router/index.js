import { createRouter, createWebHistory } from 'vue-router'
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

const router = createRouter({
  routes,
  history: createWebHistory()
})
export default router