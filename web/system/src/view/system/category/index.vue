<template>
  <div>
    <warning-bar title="子分类会跟随父菜单的修改而修改，例如：父菜单的删除会造成父菜单下的所有菜单删除，修改状态同理。"/>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addCategory('0')"
        >添加主分类
        </el-button>
        <el-button
          type="primary"
          icon="delete"
          @click="batchDeleteCategory"
        >批量删除
        </el-button>
      </div>
      <el-table
        :data="tableData"
        style="z-index: 0"
        row-key="ID"
        :tree-props="{'children': 'children'}"
        @selection-change="selectionChange"
      >
        <el-table-column type="selection" width="30"/>
        <el-table-column align="left" label="分类编号" min-width="220" prop="categoryId" fixed="left"/>
        <el-table-column align="left" label="名称" min-width="140" prop="name">
          <template #default="scope">
            <el-text type="primary" v-if="scope.row.parentId === '0'">{{ scope.row.name }}</el-text>
            <el-text type="success" v-if="scope.row.parentId !== '0'">{{ scope.row.name }}</el-text>
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序标记" min-width="120" prop="type">
          <template #default="scope">
            <el-text truncated style="position: relative;top: 4px">{{ scope.row.sorted }}</el-text>
          </template>
        </el-table-column>
        <el-table-column align="left" label="所属" min-width="120" prop="type">
          <template #default="scope">
            <el-text v-if="scope.row.type === 0" truncated style="position: relative;top: 4px">系统分类</el-text>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" min-width="120" prop="status">
          <template #default="scope">
            <el-select v-model="scope.row.status"
                       size="small"
                       @change="changeStatus(scope.row.categoryId,scope.row.status)"
                       placeholder="选择分类状态">
              <el-option
                v-for="item in statusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </template>
        </el-table-column>

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
        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="plus"
              @click="addCategory(scope.row.categoryId)"
              v-if="scope.row.parentId === '0'"
            >添加子分类
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updateCategory(scope.row.categoryId)"
            >修改
            </el-button>
            <el-button
              type="danger"
              link
              icon="delete"
              @click="deleteCategory(scope.row.categoryId)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <category-form ref="formRef" @success="loadTableData"/>
  </div>

</template>

<script setup>
import { formatTimestamp } from '@/utils/date.js'
import { ref } from 'vue'
import {
  categoryBatchDeleteApi,
  categoryChangeStatus,
  categoryListAllApi
} from '@/api/system/category.js'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage } from 'element-plus'
import CategoryForm from '@/view/system/category/component/categoryForm.vue'


// 批量删除结构
const kvs = ref([])
// 弹出层 ref
const formRef = ref(null)
// 表格数据
const tableData = ref([])
// 选择框 option
const statusOptions = ref([
  {
    label: '暂存',
    value: 0
  },
  {
    label: '使用',
    value: 1
  }
])

const loadTableData = async() => {
  const res = await categoryListAllApi()
  if (res.code === 0) {
    tableData.value = res.data.list
  }
}
loadTableData()

const changeStatus = async(id, status) => {
  const res = await categoryChangeStatus({ categoryId: id, status: status })
  if (res.code === 0) {
    ElMessage({
      message: '更新状态成功',
      type: 'success',
    })
  }
  await loadTableData()
}

// 添加根分类
const addCategory = (id) => {
  formRef.value.setFormParentId(id)
  formRef.value.isEdit = false
  if (id !== '0') {
    formRef.value.openDialog('添加子级分类')
  } else {
    formRef.value.openDialog('添加主分类')
  }
}

// 更新分类
const updateCategory = (id) => {
  // 设置表单数据 ID
  formRef.value.setFormCategoryId(id)
  formRef.value.isEdit = true
  formRef.value.openDialog('修改分类信息')
}

// 多选框变化
const selectionChange = (val) => {
  kvs.value = val.map(i => {
    return {
      parentId: i.parentId,
      categoryId: i.categoryId
    }
  })
}

// 批量删除
const batchDeleteCategory = async() => {
  const res = await categoryBatchDeleteApi({ kvs: kvs.value })
  if (res.code === 0) {
    ElMessage({
      message: '删除菜单成功',
      type: 'success',
    })
    await loadTableData()
  }
}
</script>

<style scoped lang="scss">

</style>
