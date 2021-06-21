<template>
  <div class="myself-wrap">
    <span>addr:{{form.addr}}</span>
    <br/>
    <span >id:{{form.id}}</span>
    <br/>
    <span>name:{{form.name}}</span>
    <br/>
    <span>telephone:{{form.tel}}</span>
    <br/>
    <span>money:{{form.money}}</span>
    <br/>
    <span>department:{{form.dept}}</span>
  </div>
</template>

<script>
export default {
  name: "myself",
  data(){
    return{
      form:{}
    }
  },
  created() {
    this.$nextTick(()=> {
      this.$axios.post(this.HOST + '/myself').then(result => {
        this.form = result.data;
      }).catch(excp => {
        console.log(excp);
      })
    });
  }
}
</script>

<style scoped>
.myself-wrap {
  border-radius: 15px;
  background-clip: padding-box;
  margin: 90px auto;
  width: 450px;
  padding: 35px 35px 15px 35px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}
.span{
  font-size:12px;
  text-align: center;
  font-family: 楷体;
}

</style>
