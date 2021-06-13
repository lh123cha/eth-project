import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Registe from '@/components/registe'
import Login from '@/components/login'
import Allorder from '@/components/allorder'
import Home from '@/components/home'
import Sendorder from '@/components/order/sendorder'
Vue.use(Router)

export default new Router({
  mode:"history",
  routes: [
    {
      path: '/',
      name: 'Registe',
      component: Registe
    },{
    path:'/login',
      name:'login',
      component:Login
    },{
      path:'/order_home',
      component:Home,
      children:[
        {
          path:'/order_home/basetable',
          component:Allorder,
      },{
          path:'/order_home/sendorder',
          component:Sendorder
        }
      ]
    }
  ]
})
