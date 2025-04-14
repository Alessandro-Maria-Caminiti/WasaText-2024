import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ChatView from '../views/ChatView.vue'
import GroupInfo from '../views/GroupInfo.vue'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', component: LoginView },
    { path: '/home', component: HomeView },
    { path: '/chats', component: ChatView },
    { path: '/chats/:user/:convId', component: ChatView },
    { path: '/chats/:user/:groupId/GetGroup', component: GroupInfo },
  ]
})

export default router
