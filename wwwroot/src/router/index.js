import Vue from 'vue'
import Router from 'vue-router'
import SignIn from '@/components/sign_in'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'sign_in',
      component: SignIn
    }
  ]
})
