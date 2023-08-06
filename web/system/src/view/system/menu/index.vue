<template>
  <div class="gva-table-box">
    <div class="gva-btn-list">
      <el-button type="primary" icon="Plus" @click="addMenu('添加根菜单')">添加根菜单</el-button>
      <el-select v-model="roleData" placeholder="请选择角色" style="width: 130px;margin-left: 20px" @change="loadData">
        <el-option
          v-for="item in roleList"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </div>
    <el-table
      :data="tableData"
      row-key="ID"
      :tree-props="{'children': 'children'}"
    >
      <el-table-column align="left" label="ID" min-width="100" prop="ID" fixed="left" />
      <el-table-column align="left" label="菜单名称" min-width="150" prop="title" />
      <el-table-column align="left" label="归属" min-width="100" prop="role">
        <template #default="scope">
          <el-text v-if="scope.row.role === 1" type="primary">管理员</el-text>
          <el-text v-if="scope.row.role === 2" type="warning">商家</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="组件位置" min-width="200" prop="component" />
      <el-table-column align="left" label="排序" min-width="100" prop="sorted" />
      <el-table-column align="left" label="状态" min-width="150">
        <template #default="scope">
          <el-switch
            v-model="scope.row.status"
            :disabled="scope.row.ID === 2 || scope.row.ID === 3 || scope.row.ID === 1 || scope.row.ID === 4"
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
            @click="editMenu(scope.row.ID)"
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

    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :inline="true"
        label-position="top"
        label-width="85px"
      >
        <el-form-item prop="title" label="路由展示名称">
          <el-input v-model="formData.title" placeholder="请输入展示名称" autocomplete="off" />
        </el-form-item>
        <el-form-item prop="title" label="父节点 ID">
          <el-cascader
            v-model="formData.parentId"
            style="width:100%"
            :options="menuOption"
            :disabled="!isEdit"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item prop="path" label="唯一路由名称">
          <el-input v-model="formData.path" placeholder="请输入路由 path" autocomplete="off" />
        </el-form-item>
        <el-form-item prop="status" label="是否隐藏">
          <el-select v-model="formData.status" placeholder="选择菜单状态">
            <el-option
              v-for="item in statusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item prop="role" label="菜单角色" style="width: 140px;">
          <el-select v-model="formData.role" placeholder="选择菜单角色">
            <el-option
              v-for="item in roleList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
              :disabled="item.value === 0"
            />
          </el-select>
        </el-form-item>
        <el-form-item prop="sorted" label="排序标记">
          <el-input-number v-model="formData.sorted" :min="0" />
        </el-form-item>
        <el-form-item prop="icon" label="图标">
          <icon :meta="formData.meta" style="width:100%" />
        </el-form-item>
        <el-form-item prop="component" label="文件路径">
          <el-input v-model="formData.component" placeholder="请输入文件路径 :xxx/xxx.vue" autocomplete="off" />
          <span style="font-size:12px;margin-right:12px;">如果菜单包含子菜单，请创建router-view二级路由页面或者</span>
          <el-button style="margin-top:4px" @click="formData.component = 'routerHolder.vue'">点我设置</el-button>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="submitForm">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { formatDate } from '@/utils/format.js'
import { menuChangeStatusApi, menuCreateApi, menuFindApi, menuTreeListAllApi } from '@/api/system/menu.js'
import { ElMessage } from 'element-plus'
import { useRouterStore } from '@/store/model/router.js'
import router from '@/router/index.js'
import { useMenuStore } from '@/store/model/menu.js'
import Icon from '@/view/system/menu/component/icon.vue'
import { firstToUpper } from '@/utils/string.js'

const routerStore = useRouterStore()
const menuStore = useMenuStore()
// 表格数据
const tableData = ref([])
// 角色列表
const roleList = ref([
  { label: '系统管理员', value: 1 },
  { label: '系统商家', value: 2 },
  { label: '全部菜单', value: 0 },
])
// dialog
const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const formData = ref({
  title: '',
  path: '',
  icon: '',
  status: 0,
  parentId: 0,
  component: '',
  sorted: 0,
  meta: {
    icon: ''
  },
  name: ''
})
// 级联选择器
const menuOption = ref([
  {
    ID: 0,
    title: '根菜单'
  }
])
// 路由状态选择器
const statusOptions = ref([
  {
    value: 0,
    label: '隐藏'
  }, {
    value: 1,
    label: '激活'
  }
])
// 是否为编辑状态
const isEdit = ref(false)
const roleData = ref(0)

// 加载表格数据
const loadData = async() => {
  const res = await menuTreeListAllApi({ role: roleData.value })
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
  if (val.status === 0) {
    const index = menuStore.tabList.findIndex(i => i.path === val.path)
    menuStore.tabList.splice(index, 1)
  }

  // 重新设置 router
  setTimeout(async() => {
    await routerStore.setAsyncRouter()
    routerStore.asyncRouterList.forEach(i => router.addRoute('Layout', i))
  }, 500)
}

// 打开弹出层
const openDialog = (val) => {
  dialogTitle.value = val
  dialogVisible.value = true
}

// 关闭弹出层
const closeDialog = () => {
  dialogVisible.value = false
  // 清空表单
  formRef.value.resetFields()
}

// 添加菜单
const addMenu = (val) => {
  formData.value.parentId = 0
  openDialog(val)
}

// 提交表单
const submitForm = async() => {
  // TODO 表单验证
  formData.value.icon = formData.value.meta.icon
  formData.value.name = firstToUpper(formData.value.path)
  let res
  if (isEdit.value) { // 编辑状态
    console.log('编辑')
  } else { // 创建状态
    res = await menuCreateApi(formData.value)
  }
  if (res['code'] === 0) {
    ElMessage({
      message: '添加根菜单成功',
      type: 'success',
      showClose: true,
    })
  }
  closeDialog()
  await loadData()

  // 重新设置 router
  setTimeout(async() => {
    await routerStore.setAsyncRouter()
    routerStore.asyncRouterList.forEach(i => router.addRoute('Layout', i))
  }, 500)
}

// 编辑菜单
const editMenu = async(val) => {
  const res = await menuFindApi({ id: val })
  formData.value = res.data.menu
  openDialog('编辑菜单')
}
</script>
<style lang="scss" scoped>
:deep(.el-input-number__increase), :deep(.el-input-number__decrease) {
  .el-icon {
    color: #4D70FF;
  }
}

</style>
