<template>
  <div class="gva-table-box">
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
      />
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

const tableData = ref([])
const loading = ref(false)
const total = ref(0)
const page = ref(1)
const pageSize = ref(15)

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
  const res = await pageListLogApi({
    page: page.value,
    pageSize: pageSize.value,
    username: '',
    method: '',
    path: '',
    endTime: '',
    startTime: '',
    statusCode: 0
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
  loading.value = false
}
loadTableData()

</script>
