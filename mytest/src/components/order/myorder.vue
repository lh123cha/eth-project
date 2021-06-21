<template>
  <div class="basetable" v-loading="loading" element-loading-text="拼命加载中">
    <div class="selectMenu">
      <el-date-picker v-model="value6" type="daterange" placeholder="选择日期范围">
      </el-date-picker>
    </div>


    <div class="tableMain">
      <el-table :data="tableData" style="width: 100%">

        <el-table-column prop="id" label="订单号" width="180">
        </el-table-column>
        <el-table-column prop="username" label="订单发起方" width="180">
        </el-table-column>
        <el-table-column prop="money" label="金额" width="180" sortable>
        </el-table-column>

        <el-table-column prop="missin" label="任务" width="180">
        </el-table-column>

        <el-table-column prop="tip" label="备注">
        </el-table-column>

        <el-table-column prop="time" label="时间限制" sortable>
        </el-table-column>

        <el-table-column label="操作">

          <template slot-scope="scope">
            <el-button size="small" type="success" @click="handleFinish(scope.$index, scope.row)">完成订单
            </el-button>
            <el-button size="small" type="danger" @click="handleCancle(scope.$index, scope.row)">取消接单
            </el-button>
          </template>
        </el-table-column>

      </el-table>
    </div>
    <div class="page">
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page.sync="currentPage3" :page-size="100" layout="prev, pager, next, jumper" :total="1000">
      </el-pagination>
    </div>

<!--    <el-dialog title="用户信息" :visible.sync="dialogFormVisible">-->
<!--      <el-form :model="form">-->
<!--        <el-form-item label="地址" :label-width="formLabelWidth">-->
<!--          <el-input v-model="form.address" auto-complete="off"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="姓名" :label-width="formLabelWidth">-->
<!--          <el-input v-model="form.name" auto-complete="off"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="日期" :label-width="formLabelWidth">-->
<!--          <el-date-picker v-model="form.date" type="date" placeholder="选择日期">-->
<!--          </el-date-picker>-->
<!--        </el-form-item>-->

<!--        <el-form-item label="性别" :label-width="formLabelWidth">-->
<!--          <el-select v-model="form.region" placeholder="性别">-->
<!--            <el-option label="男" value="男"></el-option>-->
<!--            <el-option label="女" value="女"></el-option>-->
<!--          </el-select>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
<!--      <div slot="footer" class="dialog-footer">-->
<!--        <el-button @click="cancel">取 消</el-button>-->
<!--        <el-button type="primary" @click="update">确 定</el-button>-->
<!--      </div>-->
<!--    </el-dialog>-->
  </div>
</template>

<script type="text/ecmascript-6">
export default {
  data() {
    return {
      loading: true,
      tableData: [{
        id:'2021061701',
        username:'lh',
        money:'25',
        mission:'带饭',
        tip:'带瓶快乐水，有小费',
        time:'25min'
      },{
        id:'2021061702',
        username:'czr',
        money:'2',
        mission:'大作业问卷调查帮答',
        tip:'全部作答，需要认真作答，作答完毕有奖励！！！',
        time:'明天之前完成'
      }],
      dialogFormVisible: false,
      formLabelWidth: '80px',
      form: {},
      value6: '',
      currentPage3: 1,
      currentIndex: '',
    }
  },
  created() {
    setTimeout(() => {
      this.loading = false
    }, 1500)
    this.$nextTick(() => {
      // 在此处执行你要执行的函数
      this.$axios.post(this.HOST+'/myorder').then(result=>{
        this.tableData=result.data
      })
    });

  },
  methods: {
    showTime() {
      this.$alert(this.value6, '起止时间', {
        confirmButtonText: '确定',
        callback: action => {
          this.$message({
            type: 'info',
            message: '已显示'
          })
        }
      })
    },
    update() {
      this.form.date = this.form.date
      this.tableData.push(this.form)
      this.dialogFormVisible = false
    },


    handleFinish(index, row) {
      this.$confirm('确定完成该订单', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.form = this.tableData[index]
        this.currentIndex = index
        let params={
          Name:this.form.id
        }

        this.$axios.post(this.HOST+'/finish_deal',params).then(result=>{
          console.log(result.data)
        }).catch(resp =>{
          console.log(resp);
        });
        this.$message({
          type: 'success',
          message: '订单完成成功!'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        })
      })
      this.dialogFormVisible = false
    },

    handleCancle(index, row) {
      this.$confirm('确定取消接收该订单', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.form = this.tableData[index]
        this.currentIndex = index
        let params={
          Name:this.form.id
        }

        this.$axios.post(this.HOST+'/cancel_deal',params).then(result=>{
          console.log(result.data)
        }).catch(resp =>{
          console.log(resp);
        });
        this.$message({
          type: 'success',
          message: '订单取消成功!'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消失败'
        })
      })
      this.dialogFormVisible = false
    },


    cancel() {
      this.dialogFormVisible = false
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`)
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`)
    }
  },
}
</script>
<style lang="scss">
.basetable {
  .selectMenu {}
  .tableMain {
    margin: {
      top: 10px;
    }
  }
  .page {
    float: right;
    margin: {
      top: 10px;
    }
  }
}
</style>

