import Vue from 'vue'
import Router from 'vue-router'
import signIn from '@/components/signIn'
import panelMain from '@/components/panel/panelMain'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'sign_in',
      component: signIn
    },
    {
      path: '/panel',
      name: 'panel_main',
      component: panelMain
    }
  ]
})
