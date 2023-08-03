<template>
  <div class="gva-table-box">
    <div class="gva-btn-list" style="display: flex;justify-content: space-between">
      <el-button type="primary" icon="Plus">添加根菜单</el-button>
      <div>
        <el-select v-model="roleType" placeholder="请选择角色" style="width: 130px;margin-left: 50px">
          <el-option
            v-for="item in roleList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-button type="primary" icon="Search" style="margin-left: 10px">按角色查询</el-button>
      </div>
    </div>
    <el-table
      :data="tableData"
      row-key="ID"
      :tree-props="{'children': 'children'}"
    >
      <el-table-column align="left" label="ID" min-width="100" prop="ID" fixed="left" />
      <el-table-column align="left" label="菜单名称" min-width="150" prop="title" />
      <el-table-column align="left" label="组件位置" min-width="200" prop="component" />
      <el-table-column align="left" label="排序" min-width="100" prop="sorted" />
      <el-table-column align="left" label="状态" min-width="150">
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :disabled="scope.row.ID === 2 || scope.row.ID === 3 || scope.row.ID === 1 "
            inline-prompt
            :active-value="1"
            :inactive-value="0"
            active-text="已激活"
            inactive-text="未激活"
            @change="()=>{switchEnable(scope.row)}"
          />
        </template>
      </el-table-column>
      <el-table-column align="left" label="创建时间" min-width="170">
        <template #default="scope">
          {{ formatDate(scope.row.createdAt) }}
        </template>
      </el-table-column>
      <el-table-column align="left" label="更新时间" min-width="170">
        <template #default="scope">
          {{ formatDate(scope.row.updatedAt) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="250" fixed="right">
        <template #default="scope">
          <el-button
            type="primary"
            link
            icon="plus"
            :disabled="scope.row.parentId !== 0"
            @click="addVideoCategory(scope.row.ID)"
          >添加子分类
          </el-button>
          <el-button
            type="primary"
            link
            icon="edit"
            @click="editVideoCategory(scope.row.ID)"
          >编辑
          </el-button>
          <el-button
            :disabled="scope.row.ID === 2 || scope.row.ID === 3"
            type="primary"
            link
            icon="delete"
            @click="deleteVideoCategory(scope.row.ID)"
          >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

  </div>
</template>
<script setup>
import { ref } from 'vue'
import { formatDate } from '@/utils/format.js'
import { menuChangeStatusApi, menuTreeListAllApi } from '@/api/system/menu.js'
import { ElMessage } from 'element-plus'
import { useRouterStore } from '@/store/model/router.js'
import router from '@/router/index.js'

const routerStore = useRouterStore()
// 表格数据
const tableData = ref([])
// 角色类型
const roleType = ref(null)
// 角色列表
const roleList = ref([
  { label: '系统管理员', value: 1 },
  { label: '系统商家', value: 2 },
])
// 加载表格数据
const loadData = async() => {
  const res = await menuTreeListAllApi()
  tableData.value = res.data.list
}
loadData()

// 更新状态
const switchEnable = async(val) => {
  const res = await menuChangeStatusApi({ id: val.ID, pid: val.parentId, status: val.status })
  if (res['code'] === 0) {
    ElMessage({
      message: '更新状态成功',
      type: 'success',
    })
  }
  await loadData()
  // TODO 标签页删除
  // 重新设置 router
  setTimeout(async() => {
    await routerStore.setAsyncRouter()
    routerStore.asyncRouterList.forEach(i => router.addRoute('Layout', i))
  }, 500)
}

</script>
