<template>
  <div class="gva-table-box">
    <div class="gva-btn-list">
      <div style="width: 300px;margin-right:40px;display: flex;align-items: center">
        <span style="width: 80px">用户名：</span>
        <el-input
            placeholder="请输入"
            v-model="username"
            width="200px"
            @change="loadTableData"/>
      </div>
      <div style="margin-right:40px;display: flex;align-items: center">
        <span style="width: 100px">请求方式：</span>
        <el-select
            placeholder="请选择"
            v-model="method"
            @change="loadTableData"
        >
          <el-option
              v-for="item in  methodDict"
              :key="item.value"
              :label="item.label"
              :value="item.value"/>
        </el-select>
      </div>
      <div style="width: 300px;margin-right:40px;display: flex;align-items: center">
        <span style="width: 50px">路径：</span>
        <el-input
            placeholder="请输入"
            v-model="path"
            width="200px"
            @change="loadTableData"/>
      </div>
      <div style="width: 350px;margin-right:40px;display: flex;align-items: center">
        <span>时间：</span>
        <el-date-picker
            @change="loadTableData"
            v-model="timeArr"
            type="daterange"
            range-separator="-"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
            value-format="YYYY-MM-DD HH:mm:ss"
            style="width: 240px;"
        />
      </div>
      <el-button
          type="primary"
          @click="resetQuery"
          :icon="Refresh">
        重置
      </el-button>
    </div>
    <el-table
        :data="tableData"
        v-loading="loading"
    >
      <el-table-column type="selection" width="50"/>
      <el-table-column
          prop="id"
          label="ID"
          width="80"/>
      <el-table-column
          prop="username"
          label="用户名"
          width="300"/>
      <el-table-column
          prop="ip"
          label="IP"
      />
      <el-table-column
          prop="method"
          label="请求方式"
          width="120"/>
      <el-table-column
          prop="path"
          label="请求路径"
          width="300"
      />
      <el-table-column
          prop="latency"
          label="请求时长">
        <template #default="scope">
          <el-text type="success" v-if="scope.row.latency < 200">{{ scope.row.latency }}ms</el-text>
          <el-text type="warning" v-if="scope.row.latency >= 200">{{ scope.row.latency }}ms</el-text>

        </template>
      </el-table-column>
      <el-table-column
          prop="time"
          label="请求时间"
          width="180">
        <template #default="scope">
          {{ formatTimestamp(scope.row.time) }}
        </template>
      </el-table-column>
      <el-table-column
          align="right"
          label="操作"
          width="200"
      >
        <template #default="scope">
          <el-button type="text" size="small">删除</el-button>
          <el-button type="text" size="small">详情</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="gva-pagination">
      <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
      />
    </div>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {pageListLogApi} from "../../api/system/log.js";
import {formatTimestamp} from "../../utils/date.js";
import {Refresh} from "@element-plus/icons-vue";

const tableData = ref([])
const loading = ref(false)
const total = ref(0)
const page = ref(1)
const pageSize = ref(15)
const timeArr = ref([])
const path = ref('')
const method = ref('')
const username = ref('')

const methodDict = ref([
  {label: 'GET', value: 'GET'},
  {label: 'POST', value: 'POST'},
])

const resetQuery = () => {
  page.value = 1
  pageSize.value = 15
  timeArr.value = []
  path.value = ''
  method.value = ''
  username.value = ''
  loadTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  loadTableData()
}

const loadTableData = async () => {
  loading.value = true

  const data = {
    page: page.value,
    pageSize: pageSize.value,
    method: method.value,
    path: path.value,
    username: username.value,
    startTime: '',
    endTime: '',
    statusCode: 0
  }

  if (timeArr.value.length > 0) {
    data.startTime = timeArr.value[0]
    data.endTime = timeArr.value[1]
  } else {
    data.startTime = ''
    data.endTime = ''
  }

  const res = await pageListLogApi(data)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
  loading.value = false
}
loadTableData()

</script>
