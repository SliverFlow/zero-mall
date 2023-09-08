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
        @click="editMenu(scope.row.ID)"
      >搜索
      </el-button>
    </div>

    <el-table
      :data="tableData"
      style="z-index: 0"
      row-key="ID"
      :tree-props="{'children': 'children'}"
    >
      <el-table-column align="left" label="商户编号" min-width="180" prop="businessId" fixed="left" />
      <el-table-column align="left" label="封面" min-width="180" prop="image">
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
      <el-table-column align="left" label="状态" min-width="110" prop="status">
        <template #default="scope">
          <el-text v-if="scope.row.status === 0" type="info">暂存</el-text>
          <el-text v-if="scope.row.status === 1" type="success">正常</el-text>
          <el-text v-if="scope.row.status === 2" type="danger">禁用</el-text>
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
      <el-table-column label="操作" min-width="270" fixed="right">
        <template #default="scope">
          <el-button
            :disabled="scope.row.ID === 1"
            type="primary"
            link
            icon="delete"
            @click="deleteVideoCategory(scope.row.ID)"
          >查看管理员信息
          </el-button>
          <el-button
            :disabled="scope.row.ID === 1"
            type="danger"
            link
            icon="delete"
            @click="deleteVideoCategory(scope.row.ID)"
          >暂停商户使用
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
import { formatTimestamp } from '@/utils/date.js'
import { ref } from 'vue'
import { businessListApi } from '@/api/system/business.js'

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
  const res = await businessListApi({ page: page.value, pageSize: pageSize.value, keyWord: keyWord.value })
  if (res['code'] === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}
loadTableData()
</script>

<style scoped lang="scss">

</style>
