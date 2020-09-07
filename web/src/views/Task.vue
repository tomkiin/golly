<template>
  <div>
    <el-card shadow="never">
      <span style="font-size: 14px">任务名称</span>
      <el-input v-model="addTaskData.name" suffix-icon="el-icon-edit" placeholder="请输入任务名称" size="small"
                style="width: 200px; margin-left: 20px"></el-input>
      <span style="font-size: 14px; margin-left: 20px">种子地址</span>
      <el-input v-model="addTaskData.seed" suffix-icon="el-icon-monitor" placeholder="请输入种子地址" size="small"
                style="width: 200px; margin-left: 20px"></el-input>
      <el-button @click="addTask" type="primary" icon="el-icon-circle-plus-outline" size="small"
                 style="margin-left: 20px">添加任务
      </el-button>
      <el-button @click="getTask" type="info" icon="el-icon-refresh" size="small" style="margin-left: 20px">刷新状态
      </el-button>
    </el-card>
    <el-table :data="taskData" size="small" stripe border style="margin-top: 20px">
      <el-table-column prop="id" label="ID" align="center"/>
      <el-table-column prop="name" label="任务名称" align="center"/>
      <el-table-column prop="seed" label="种子地址" align="center"/>
      <el-table-column prop="base_url" label="源地址" align="center"/>
      <el-table-column prop="status" label="运行状态" align="center">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === '就绪'" effect="plain">{{ scope.row.status }}</el-tag>
          <el-tag v-if="scope.row.status === '运行'" type="success" effect="plain">{{ scope.row.status }}</el-tag>
          <el-tag v-if="scope.row.status === '结束'" type="danger" effect="plain">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="edit" label="编辑" align="center">
        <template slot-scope="scope">
          <el-button @click="getOneTask(scope.row.id)" type="text"
                     size="small">配置
          </el-button>
          <el-button @click="deleteTask(scope.row.id)" type="text" size="small">删除</el-button>
          <el-button type="text" @click="runTask(scope.row.id)" size="small" :loading="scope.row.status === '运行'">运行
          </el-button>
          <el-button type="text" @click="abortTask(scope.row.id)" size="small" :disabled="scope.row.status !== '运行'">终止
          </el-button>
        </template>
      </el-table-column>
      <el-table-column prop="edit" label="日志" align="center">
        <template slot-scope="scope">
          <el-button @click="showLog(scope.row.id)" :type="scope.row.status !== '运行' ? 'info' : 'success'" size="small"
                     circle icon="el-icon-search"></el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination layout="prev, pager, next" :total="this.pageInfo.total" @current-change="handleCurrentChange"
                   style="float: right; margin-top: 10px">
    </el-pagination>
    <el-drawer :visible.sync="isTaskDrawerShow" :with-header="false" size="40%">
      <el-card shadow="never">
        <el-form :model="updateTaskData" :label-position="labelPosition">
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="任务名称">
                <el-input v-model="updateTaskData.name" size="small"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="种子地址">
                <el-input v-model="updateTaskData.seed" size="small"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="源地址">
                <el-input v-model="updateTaskData.base_url" size="small"/>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-card>
                <div slot="header">
                  <span>收集器规则</span>
                  <el-button @click="updateTaskData.rules.pop()" style="float: right; padding: 3px 0" type="text">删除规则
                  </el-button>
                  <el-button @click="addTaskRule" style="float: right; padding: 3px 0; margin-right: 10px" type="text">
                    添加规则
                  </el-button>
                </div>
                <div v-for="(item, index) in updateTaskData.rules" :key="index">
                  <el-col :span="24">
                    <el-divider>收集器 {{ index + 1 }}</el-divider>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="收集器名称">
                      <el-input v-model="item.name" size="small"/>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="解析规则">
                      <el-input v-model="item.pattern" size="small"/>
                    </el-form-item>
                  </el-col>
                  <el-col :span="8">
                    <el-form-item label="运行线程数">
                      <el-select v-model.number="item.runtime" size="small">
                        <el-option label="1" value="1"/>
                        <el-option label="2" value="2"/>
                        <el-option label="3" value="3"/>
                        <el-option label="4" value="4"/>
                        <el-option label="5" value="5"/>
                      </el-select>
                    </el-form-item>
                  </el-col>
                  <el-col :span="8">
                    <el-form-item label="数据类型">
                      <el-select v-model="item.find" size="small">
                        <el-option label="attr" value="attr"/>
                        <el-option label="text" value="text"/>
                      </el-select>
                    </el-form-item>
                  </el-col>
                  <el-col :span="8" v-show="item.find === 'attr'">
                    <el-form-item label="参数名称">
                      <el-input v-model="item.attr" size="small"/>
                    </el-form-item>
                  </el-col>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
      <div style="margin: 20px 0 0 20px">
        <el-button type="primary" size="small" @click="updateTask">更新</el-button>
        <el-button size="small" @click="isTaskDrawerShow = false">取消</el-button>
      </div>
    </el-drawer>
    <el-drawer :visible.sync="isLogDrawerShow" :before-close="handleLogClose" :with-header="false" size="40%">
      <el-timeline :reverse="true" style="margin: 20px 0 20px 0">
        <el-timeline-item
            v-for="(item, index) in logs"
            :key="index"
            :timestamp="item.timestamp">
          {{ item.data }}
        </el-timeline-item>
      </el-timeline>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: "Task",
  data() {
    return {
      taskData: [],
      addTaskData: {},
      updateTaskData: {},
      isTaskDrawerShow: false,
      pageInfo: {total: 0, page: 1, size: 10},
      labelPosition: "top",
      isLogDrawerShow: false,
      logs: [],
      timer: {}
    }
  },
  methods: {
    getTask() {
      this.$router.push("?page=" + this.pageInfo.page, () => {
      })
      let args = {
        page: this.$route.params.page ? this.$route.params.page : this.pageInfo.page,
        size: this.pageInfo.size
      }
      this.axios.get("/tasks", {params: args}).then(resp => {
        let data = resp.data
        if (data.code !== 0) {
          this.$message.error(data.message)
        } else {
          this.pageInfo.total = data.total
          this.taskData = data.data
          if (this.taskData.length === 0 && this.pageInfo.page > 1) {
            this.pageInfo.page -= 1
            this.getTask()
          }
        }
      }, e => {
        this.$message.error(e)
      })
    },
    getOneTask(id) {
      this.isTaskDrawerShow = true
      this.axios.get("/task/" + id).then(resp => {
        let data = resp.data
        if (data.code !== 0) {
          this.$message.error(data.message)
        } else {
          this.updateTaskData = data.data
        }
      }, e => {
        this.$message.error(e)
      })
    },
    addTask() {
      if (this.addTaskData.name === undefined) {
        this.$message.warning("请输入任务名称")
        return
      }
      if (this.addTaskData.seed === undefined) {
        this.$message.warning("请输入种子地址")
        return
      }
      this.axios.post("/task", this.addTaskData).then(resp => {
        let data = resp.data
        if (data.code !== 0) {
          this.$message.error(data.message)
        }
        this.addTaskData = {}
        this.getTask()
      }, e => {
        this.$message.error(e)
      })
    },
    updateTask() {
      this.axios.put("/task", this.updateTaskData).then(resp => {
        let data = resp.data
        if (data.code !== 0) {
          this.$message.error(data.message)
        }
        this.getTask()
      }, e => {
        this.$message.error(e)
      })
      this.isTaskDrawerShow = false
    },
    deleteTask(id) {
      this.axios.delete("/task/" + id).then(resp => {
        let data = resp.data
        if (data.code !== 0) {
          this.$message.error(data.message)
        }
        this.getTask()
      }, e => {
        this.$message.error(e)
      })
    },
    runTask(id) {
      this.axios.post("/task/" + id + "/run").then(resp => {
        let data = resp.data
        if (data.code !== 0) {
          this.$message.error(data.message)
        } else {
          this.getTask()
        }
      }, e => {
        this.$message.error(e)
      })
    },
    abortTask(id) {
      this.axios.post("/task/" + id + "/abort").then(resp => {
        let data = resp.data
        if (data.code !== 0) {
          this.$message.error(data.message)
        }
        this.getTask()
      }, e => {
        this.$message.error(e)
      })
    },
    logTask(id) {
      return () => {
        this.axios.get("/task/" + id + "/log").then(resp => {
          let data = resp.data
          if (data.code !== 0) {
            this.logs.push({
              timestamp: "now",
              data: "运行结束"
            })
          } else {
            data.data.forEach(e => {
              this.logs.push(e)
            })
          }
        }, e => {
          this.$message.error(e)
        })
      }
    },
    showLog(id) {
      this.isLogDrawerShow = true
      this.timer = setInterval(this.logTask(id), 1000)
    },
    handleLogClose() {
      this.isLogDrawerShow = false
      this.logs = []
      clearInterval(this.timer);
    },
    addTaskRule() {
      let taskID = this.updateTaskData.id
      let taskRule = {
        task_id: taskID,
        pattern: "",
        runtime: 1,
        find: "attr",
        attr: "href",
      }
      if (this.updateTaskData.rules === null) {
        this.updateTaskData.rules = []
      }
      this.updateTaskData.rules.push(taskRule)
    },
    handleCurrentChange(page) {
      this.pageInfo.page = page
      this.getTask()
    }
  },
  created() {
    this.isLogDrawerShow = false
    this.getTask()
  }
}
</script>

<style>
.el-drawer__body {
  overflow-y: auto;
}
</style>
