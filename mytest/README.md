# chainfront

> A Vue.js project

## 安装依赖包并启动前端

首先我们需要在虚拟机中下载好前端代码

```bash
git clone git@github.com:lh123cha/eth-project.git
```

之后进入mytest文件安装依赖

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

#可以选择build也可以不进行build
# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).

可能有时会出现安装包版本过高的报错，这时需要卸载对应的安装包，再下载低版本的安装包。可以参考的解决方案：

[vue-loader版本报错]: https://blog.csdn.net/qq_43329216/article/details/108670797



成功启动之后前端会运行再8080端口，进入http://localhost:8080，进入登录界面说明安装启动成功。

## 与后端交互

在mytest/src/main.js文件内的

```python
Vue.prototype.HOST = 'http://localhost:5000'
```

为后端服务的端口号，可以修改它来与指定的后端端口号进行交互。

前端所有页面的路由信息在文件mytest/src/router/index.js文件内，可以看到整个项目的路由信息。

```python
export default new Router({
  mode:"history",
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },{
    path:'/registe',
      name:'Registe',
      component:Registe
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
        },{
          path:'/order_home/myorder',
          component:Myorder
        },{
          path:'/order_home/myself',
          component:Myself
        },{
          path:'/order_home/myreceiveorder',
          component:Myreceiveorder
        }
      ]
    }
  ]
})
```

