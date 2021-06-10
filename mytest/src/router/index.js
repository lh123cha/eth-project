import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Registe from '@/components/registe'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Registe',
      component: Registe
    }
  ]
})
