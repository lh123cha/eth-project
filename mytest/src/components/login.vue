<template>

  <div  class="login-wrap">
    <el-form  class="login-container" :rules="rules">
      <h1 class="title">用户登陆</h1>
      <el-form-item label="" prop="username_val">
        <el-input type="text" placeholder="姓名" v-model="Name" autocomplete="off"></el-input>
      </el-form-item>
<!--      <el-form-item label="" prop="password_val">-->
<!--        <el-input type="password" placeholder="学号" v-model="stuid" autocomplete="off"></el-input>-->
<!--      </el-form-item>-->
      <el-form-item>
        <el-button type="primary" @click="doLogin()" style="width: 100%;">用户登录</el-button>
      </el-form-item>
    </el-form>
  </div>

</template>

<script>
export default{
  data:function(){
    return {
      Name: '',
      msg:'',
      rules: {
        username_val: [
          {required: true, message: 'xingming不可为空', trigger: 'blur'}
        ]
      },
    }
  },
  methods:{
    doLogin:function(){
      let params={
        Name:this.Name,
        methodName:'userLogin',
        msg:''
      };
      console.log(params);
      this.$axios.post(this.HOST+'/valid'
        , params).then(result=>{
             console.log(result.data)
              console.log(result.data.msg)
             this.msg = result.data.msg
              if(this.msg.length==0){
                return
              }
              else{
                this.$router.push('/');
             }
      }).catch(resp =>{
        console.log(resp);
      });

    }
  }
}
</script>

<style scoped>

</style>
