<template>
  <div class="gva-table-box">
    <div class="gva-btn-list" style="float: right">
      <el-input
        v-model="keyWord"
        placeholder="请输入店铺名称关键字"
        style="width: 183px;margin-right: 16px"
        suffix-icon="FirstAidKit"
      />
      <el-button
        type="primary"
        icon="search"
        @click="loadTableData"
      >搜索
      </el-button>
    </div>

    <el-table
      :data="tableData"
      style="z-index: 0"
      row-key="ID"
      :tree-props="{'children': 'children'}"
      v-loading="loading"
    >
      <el-table-column align="left" label="商户编号" min-width="180" prop="businessId" fixed="left"/>
      <el-table-column align="left" label="封面" min-width="150" prop="image">
        <template #default="scope">
          <el-image
            style="width: 100px; height: 100px;z-index: 100;"
            :src="scope.row.image[0]"
            :zoom-rate="1.2"
            close-on-press-escape
            preview-teleported
            lazy
            :preview-src-list="scope.row.image"
            :initial-index="4"
            fit="cover"
          />
        </template>
      </el-table-column>
      <el-table-column align="left" label="名称" min-width="170" prop="name">
        <template #default="scope">
          <el-text type="primary">{{ scope.row.name }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="状态" min-width="110" prop="status">
        <template #default="scope">
          <el-text v-if="scope.row.status === 0" type="info">暂存</el-text>
          <el-text v-if="scope.row.status === 1" type="success">正常</el-text>
          <el-text v-if="scope.row.status === 2" type="danger">禁用</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="详细介绍" min-width="240" prop="detail">
        <template #default="scope">
          <el-text truncated>{{ scope.row.detail }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="评分" min-width="200" prop="score">
        <template #default="scope">
          <el-rate
            v-model="scope.row.score"
            disabled
            show-score
            score-template=""
          />
        </template>
      </el-table-column>
      <el-table-column align="left" label="创建时间" min-width="200">
        <template #default="scope">
          {{ formatTimestamp(scope.row.createdAt) }}
        </template>
      </el-table-column>
      <el-table-column align="left" label="更新时间" min-width="200">
        <template #default="scope">
          {{ formatTimestamp(scope.row.updatedAt) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="220" fixed="right">
        <template #default="scope">
          <el-button
            type="primary"
            link
            icon="edit"
            @click="findBusinessUser(scope.row.uuid)"
          >查看管理员信息
          </el-button>
          <el-button
            type="danger"
            link
            icon="delete"
            v-if="scope.row.status === 0 || scope.row.status === 1"
            @click="changeStatus(scope.row.businessId,2)"
          >禁用
          </el-button>
          <el-button
            type="success"
            link
            icon="pointer"
            v-if="scope.row.status === 2"
            @click="changeStatus(scope.row.businessId,0)"
          >解封
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

    <!--抽屉 管理员信息-->
    <el-drawer
      v-model="drawer"
      title="管理员信息"
      direction="rtl"
      size="400px"
      @close="handleClose"

    >
      <el-form :model="formData" :disabled="true">
        <el-form-item label="用户头像：" label-width="90px">
          <el-image
            style="width: 200px;height: 200px"
            :src="formData.avatar"
            :zoom-rate="1.2"
            close-on-press-escape
            preview-teleported
            fit="cover"
          />
        </el-form-item>
        <el-form-item label="用户名：" label-width="90px">
          <el-input v-model="formData.username"/>
        </el-form-item>
        <el-form-item label="昵称：" label-width="90px">
          <el-input v-model="formData.nickname"/>
        </el-form-item>
        <el-form-item label="邮箱：" label-width="90px">
          <el-input v-model="formData.email"/>
        </el-form-item>
        <el-form-item label="手机号码：" label-width="90px">
          <el-input v-model="formData.phone"/>
        </el-form-item>
        <el-form-item label="注册时间：" label-width="90px">
          <el-text type="success">{{ formatTimestamp(formData.createdAt) }}</el-text>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { formatTimestamp } from '@/utils/date.js'
import { ref } from 'vue'
import { businessChangeStatusApi, businessListApi } from '@/api/system/business.js'
import { userFindByUUIDApi } from '@/api/system/user.js'
import { ElMessage } from 'element-plus'

// 分页相关
const page = ref(1)
const pageSize = ref(10)
const keyWord = ref('')
const loading = ref(false)
// 表格相关
const total = ref(0)
const tableData = ref([])
// 抽屉
const drawer = ref(false)
// 表单数据结构
const formData = ref({
  ID: 0,
  email: '',
  nickname: '',
  phone: '',
  username: '',
  avatar: '',
  createdAt: 0
})

// 当前页发生变化
const handleCurrentChange = (val) => {
  page.value = val
  loadTableData()
}
// 页面大小发生变化
const handleSizeChange = (val) => {
  pageSize.value = val
  loadTableData()
}
const loadTableData = async() => {
  loading.value = true
  const res = await businessListApi({ page: page.value, pageSize: pageSize.value, keyWord: keyWord.value })
  if (res['code'] === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
  loading.value = false
}
loadTableData()

// 查看管理员信息
const findBusinessUser = async(uuid) => {
  const res = await userFindByUUIDApi({ uuid: uuid })
  if (res.code === 0) {
    formData.value = res.data.user
    drawer.value = true
  }
}

// 设置表单初始值
const initFormData = () => {
  formData.value = {
    ID: 0,
    email: '',
    nickname: '',
    phone: '',
    username: '',
    avatar: '',
    createdAt: 0
  }
}

// 抽屉关闭
const handleClose = () => {
  drawer.value = false
  initFormData()
}

// 修改商户状态
const changeStatus = async(id, status) => {
  const res = await businessChangeStatusApi({ businessId: id, status: status })
  if (res.code === 0) {
    ElMessage({
      message: '更新商户状态成功',
      type: 'success',
    })
    await loadTableData()
  }
}
</script>

<style scoped lang="scss">

</style>
