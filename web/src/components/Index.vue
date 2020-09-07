<template>
  <el-container>
    <el-header>
      <div class="title">Golly</div>
    </el-header>
    <el-container>
      <el-aside width="140px">
        <el-menu active-text-color="#2980b9" router>
          <el-menu-item index="/task">
            <i class="el-icon-document-copy"></i>
            <span slot="title">任务管理</span>
          </el-menu-item>
          <el-menu-item index="/system">
            <i class="el-icon-setting"></i>
            <span slot="title">系统设置</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>

export default {
  name: 'Index',
  data() {
    const item = {
      date: '2016-05-02',
      name: '王小虎',
      address: '上海市普陀区金沙江路 1518 弄'
    };
    return {
      tableData: Array(20).fill(item)
    }
  }
}
</script>

<style scoped>
.el-header {
  background-color: #304156;
}


.title {
  color: #fff;
  font-size: 28px;
  font-weight: inherit;
  line-height: 60px;
  margin-left: 20px;
}

.el-aside {
  background-color: #fff;
}

.el-main {
  background-color: #ecf0f1;
  /*padding-top: 0;*/
}
</style>
