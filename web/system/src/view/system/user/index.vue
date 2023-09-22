<template>
  <div class="gva-table-box">
    <div class="gva-btn-list" style="float: right">
      <el-input placeholder="请输入用户昵称关键字" v-model="keyWord" style="width: 183px;margin-right: 16px" suffix-icon="user" />
      <el-button
        type="primary"
        icon="search"
        @click="loadTableData"
      >搜索
      </el-button>
    </div>

    <el-table
      :data="tableData"
      row-key="ID"
      :tree-props="{'children': 'children'}"
    >
      <el-table-column align="left" label="ID" min-width="50" prop="ID" fixed="left" />
      <el-table-column align="left" label="头像" min-width="70" prop="avatar">
        <template #default="scope">
          <el-avatar
            :size="30"
            style="position: relative;top:3px;"
            :src="scope.row.avatar"
          />
        </template>
      </el-table-column>
      <el-table-column align="left" label="登录名" min-width="140" prop="username" />
      <el-table-column align="left" label="昵称" min-width="140" prop="nickname" />
      <el-table-column align="left" label="角色" min-width="100" prop="role">
        <template #default="scope">
          <el-text v-if="scope.row.role === 3" type="primary">超级管理员</el-text>
          <el-text v-if="scope.row.role === 2" type="warning">商家</el-text>
          <el-text v-if="scope.row.role === 1" type="success">普通用户</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="状态" min-width="110" prop="status">
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :disabled="scope.row.ID === 1"
            inline-prompt
            :active-value="1"
            :inactive-value="0"
            inactive-color="#f56c6c"
            active-text="正常"
            inactive-text="禁用"
            @change="()=>{switchEnable(scope.row)}"
          />
        </template>
      </el-table-column>
      <el-table-column align="left" label="电话号码" min-width="120" prop="phone" />
      <el-table-column align="left" label="邮箱" min-width="170" prop="email" />
      <el-table-column align="left" label="创建时间" min-width="170">
        <template #default="scope">
          {{ formatTimestamp(scope.row.createdAt) }}
        </template>
      </el-table-column>
      <el-table-column align="left" label="更新时间" min-width="170">
        <template #default="scope">
          {{ formatTimestamp(scope.row.updatedAt) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="110" fixed="right">
        <template #default="scope">
          <el-button
            :disabled="scope.row.ID === 1"
            type="danger"
            link
            icon="delete"
            @click="deleteVideoCategory(scope.row.ID)"
          >删除
          </el-button>
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
import { ref } from 'vue'
import { formatTimestamp } from '@/utils/date.js'
import { userListApi } from '@/api/system/user.js'

// 分页相关
const page = ref(1)
const pageSize = ref(10)
const keyWord = ref('')
// 表格相关
const total = ref(0)
const tableData = ref([])

const handleCurrentChange = (val) => {
  console.log(val)
}
const handleSizeChange = () => {
  console.log(val)
}
const loadTableData = async() => {
  const res = await userListApi({ page: page.value, pageSize: pageSize.value, keyWord: keyWord.value })
  if (res['code'] === 0) {
    tableData.value = res.data.user
    total.value = res.data.total
  }
}
loadTableData()
</script>
