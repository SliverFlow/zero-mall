<template>
  <div>
    <warning-bar title="子菜单会跟随父菜单的修改而修改，例如：父菜单的删除会造成父菜单下的所有菜单删除，修改角色同理。" />

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="Plus" @click="addMenu(0,0)">添加根菜单</el-button>
        <el-select
          v-model="roleData"
          placeholder="请选择角色"
          style="width: 130px;margin-left: 20px"
          @change="loadData"
        >
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
        :header-cell-style="{background:'#f7fbff'}"
      >
        <el-table-column align="left" label="ID" min-width="100" prop="ID" fixed="left" />
        <el-table-column align="left" label="菜单名称" min-width="150" prop="meta.title" />
        <el-table-column align="left" label="归属" min-width="100" prop="role">
          <template #default="scope">
            <el-text v-if="scope.row.role === 3" type="primary">管理员</el-text>
            <el-text v-if="scope.row.role === 2" type="warning">商家</el-text>
          </template>
        </el-table-column>
        <el-table-column align="left" label="组件位置" min-width="220" prop="component" />
        <el-table-column align="left" label="图标" min-width="100" prop="role">
          <template #default="scope">
            <component :is="scope.row.meta.icon" style="width: 18px;height: 18px;display: flex;align-items: center" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" min-width="100" prop="sorted" />
        <el-table-column align="left" label="状态" min-width="100">
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
        <el-table-column label="操作" min-width="255" fixed="right">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="plus"
              :disabled="scope.row.parentId !== 0"
              v-if="scope.row.parentId === 0"
              @click="addMenu(scope.row.ID, scope.row.role)"
            >添加子菜单
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >编辑
            </el-button>
            <el-button
              :disabled="scope.row.ID === 2 || scope.row.ID === 3 || scope.row.ID === 1 || scope.row.ID === 4"
              type="danger"
              link
              icon="delete"
              @click="deleteCategory(scope.row.ID)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-dialog
        v-model="dialogVisible"
        :title="dialogTitle"
        @close="closeDialog"
      >
        <el-form
          ref="formRef"
          :model="formData"
          :inline="true"
          label-position="top"
          label-width="85px"
        >
          <el-form-item prop="title" label="路由展示名称">
            <el-input v-model="formData.meta.title" placeholder="请输入展示名称" autocomplete="off" />
          </el-form-item>
          <el-form-item prop="title" label="父节点 ID">
            <el-cascader
              v-model="formData.parentId"
              style="width:100%"
              :options="menuOption"
              :disabled="!isEdit || (formData.ID === 2 || formData.ID=== 3 || formData.ID === 1 || formData.ID === 4)"
              :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
              :show-all-levels="false"
              filterable
            />
          </el-form-item>
          <el-form-item prop="path" label="唯一路由路径" style="width: 40%">
            <el-input v-model="formData.path" placeholder="请输入路由 path :/layout/xxx/xxx" autocomplete="off" />
          </el-form-item>
          <el-form-item prop="path" label="唯一 name">
            <el-input v-model="formData.name" placeholder="请输入唯一name" autocomplete="off" />
          </el-form-item>
          <el-form-item prop="status" label="是否隐藏">
            <el-select
              v-model="formData.status"
              placeholder="选择菜单状态"
              :disabled="formData.ID === 2 || formData.ID=== 3 || formData.ID === 1 || formData.ID === 4"
            >
              <el-option
                v-for="item in statusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="role" label="菜单角色" style="width: 140px;">
            <el-select v-model="formData.role" placeholder="选择菜单角色" :disabled="idAddChildrenMenu">
              <el-option
                v-for="item in roleList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                :disabled="item.value === 10"
              />
            </el-select>
          </el-form-item>
          <el-form-item prop="sorted" label="排序标记">
            <el-input-number v-model="formData.sorted" :min="0" />
          </el-form-item>
          <el-form-item prop="icon" label="图标">
            <icon :key="formData.ID" :meta="formData.meta" style="width:100%" />
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

  </div>

</template>
<script setup>
import { ref } from 'vue'
import {
  menuChangeStatusApi,
  menuCreateApi, menuDeleteApi,
  menuFindApi, menuTreeListAllApi,
  menuUpdateApi
} from '@/api/system/menu.js'
import { ElMessage } from 'element-plus'
import { useRouterStore } from '@/store/model/router.js'
import router from '@/router/index.js'
import { useMenuStore } from '@/store/model/menu.js'
import Icon from '@/view/system/menu/component/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { formatTimestamp } from '@/utils/date.js'

const routerStore = useRouterStore()
const menuStore = useMenuStore()
// 表格数据
const tableData = ref([])
// 角色列表
const roleList = ref([
  { label: '系统管理员', value: 3 },
  { label: '系统商家', value: 2 },
  { label: '全部菜单', value: 10 },
])
// 是否添加子菜单窗台
const idAddChildrenMenu = ref(false)
// dialog
const dialogVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const formData = ref({
  ID: 0,
  name: '',
  path: '',
  status: 0,
  parentId: 0,
  component: '',
  sorted: 0,
  meta: {
    icon: '',
    title: '',
  },
  role: 3,
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
// 按角色搜索
const roleData = ref(10)

// 加载表格数据
const loadData = async() => {
  const res = await menuTreeListAllApi({ ID: roleData.value })
  tableData.value = res.data.list
}
loadData()

// 更新状态
const switchEnable = async(val) => {
  const res = await menuChangeStatusApi({ ID: val.ID, pid: val.parentId, status: val.status })
  if (res['code'] === 0) {
    ElMessage({
      message: '更新状态成功',
      type: 'success',
    })
  }
  await loadData()
  // 当将菜单修改为未激活时，删除 tab 栏打开的当前标签页
  if (val.status === 0) {
    const index = menuStore.tabList.findIndex(i => i.path === val.path)
    if (index >= 0) {
      menuStore.tabList.splice(index, 1)
    }
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
  isEdit.value = false
  dialogVisible.value = false
  idAddChildrenMenu.value = false
  // 清空表单
  formRef.value.resetFields()
}

// 添加菜单
const addMenu = (val, role) => {
  setMenuOption()
  formData.value.parentId = val
  if (val > 0) { // 添加子菜单的情况
    formData.value.role = role
    idAddChildrenMenu.value = true
  }
  openDialog('添加菜单')
}

// 提交表单
const submitForm = async() => {
  // TODO 表单验证
  let res
  if (isEdit.value) { // 编辑状态
    res = await menuUpdateApi(formData.value)
    if (res['code'] === 0) {
      ElMessage({
        message: '更新菜单成功',
        type: 'success',
        showClose: true,
      })
    }
  } else { // 创建状态
    res = await menuCreateApi(formData.value)
    if (res['code'] === 0) {
      ElMessage({
        message: '创建菜单成功',
        type: 'success',
        showClose: true,
      })
    }
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
  const res = await menuFindApi({ ID: val })
  isEdit.value = true
  formData.value = res.data
  setMenuOption()
  openDialog('编辑菜单')
}

// 设置编辑菜单时的级联选择框数据
const setMenuOption = () => {
  menuOption.value = [{ ID: 0, title: '根菜单' }]
  asyncMenuOptions(tableData.value, menuOption.value, false)
}

// 通过数据列表获取级联选择器的数据
const asyncMenuOptions = (tableData, optionData, disabled) => {
  tableData && tableData.forEach(i => {
    if (i.children && i.children.length) {
      const option = {
        title: i.meta.title,
        ID: i.ID,
        disabled: disabled || i.ID === formData.value.ID,
        children: []
      }
      asyncMenuOptions(i.children, option.children, disabled || i.ID === formData.value.ID)
      optionData.push(option)
    } else {
      const option = {
        title: i.meta.title,
        ID: i.ID,
        disabled: disabled || i.ID === formData.value.ID,
      }
      optionData.push(option)
    }
  })
}

// 删除分类
const deleteCategory = async (val) => {
  const res = await menuDeleteApi({ID: val})
  if (res.code === 0) {
    ElMessage({
      message: '删除菜单成功',
      type: 'success',
      showClose: true,
    })
    await loadData()
  }
}
</script>
<style lang="scss" scoped>
:deep(.el-input-number__increase), :deep(.el-input-number__decrease) {
  .el-icon {
    color: #4D70FF;
  }
}
:deep(.el-table__row>.el-table__cell) {
  padding: 8px 0;
}

</style>
