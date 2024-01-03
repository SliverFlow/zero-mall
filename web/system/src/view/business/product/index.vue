<template>
  <div class="gva-table-box">
    <div class="gva-btn-list" style="display: flex;justify-content: space-between;margin-bottom: 20px">
      <el-button
        type="primary"
        icon="plus"
        @click="addProduct"
      >添加商品
      </el-button>
      <div>
        <el-input
          placeholder="请输入商品名"
          v-model="keyWord" style="width: 183px;margin-right: 16px"
          suffix-icon="MilkTea"/>
        <el-button
          type="primary"
          icon="search"
          @click="loadTableData"
        >搜索
        </el-button>
      </div>
    </div>
    <el-table
      :data="tableData"
      style="z-index: 0;height: calc(100vh - 284px)"
      row-key="ID"
      :tree-props="{'children': 'children'}"
    >
      <el-table-column align="left" label="商品编号" min-width="180" prop="productId" fixed="left"/>
      <el-table-column align="left" label="图片" min-width="150" prop="image">
        <template #default="scope">
          <el-carousel :interval="3000" arrow="never" indicator-position="none" style="height: 100px;width: 100px">
            <el-carousel-item v-for="(v,k) in scope.row.image">
              <el-image
                style="width: 100px; height: 100px;z-index: 100;"
                :zoom-rate="1.2"
                close-on-press-escape
                preview-teleported
                :src="v.url"
                lazy
                :initial-index="4"
                fit="cover"
              />
            </el-carousel-item>
          </el-carousel>
        </template>
      </el-table-column>
      <el-table-column align="left" label="名称" min-width="170" prop="name">
        <template #default="scope">
          <el-text type="primary">{{ scope.row.name }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="副标题" min-width="170" prop="subtitle">
        <template #default="scope">
          <el-text>{{ scope.row.subtitle }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="状态" min-width="110" prop="status">
        <template #default="scope">
          <el-button v-if="scope.row.status === 0" type="info" plain disabled>暂存</el-button>
          <el-button v-if="scope.row.status === 1" type="success" plain disabled>在售</el-button>
          <el-button v-if="scope.row.status === 2" type="danger" plain disabled>下架</el-button>
        </template>
      </el-table-column>
      <el-table-column align="left" label="价格" min-width="170" prop="price">
        <template #default="scope">
          <el-text type="primary">{{ scope.row.price.toFixed(2) }} 元</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="库存" min-width="170" prop="stock">
        <template #default="scope">
          <el-text type="primary">{{ scope.row.stock }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="系统分类" min-width="170" prop="categories">
        <template #default="scope">
          <template v-for="(v,k) in scope.row.categories">
            <el-text type="primary" v-if="v.type === 0">{{ v.name }}</el-text>
          </template>

        </template>
      </el-table-column>
      <el-table-column align="left" label="详细介绍" min-width="240" prop="detail">
        <template #default="scope">
          <el-text truncated>{{ scope.row.detail }}</el-text>
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
      <el-table-column label="操作" min-width="210" fixed="right">
        <template #default="scope">
          <el-button
            type="primary"
            link
            icon="edit"
            @click="() => editProduct(scope.row.productId)"
          >编辑
          </el-button>
          <el-button
            type="success"
            link
            icon="pointer"
            v-if="scope.row.status === 0 || scope.row.status === 2"
            @click="changeStatus(scope.row.productId,1)"
          >上架
          </el-button>
          <el-button
            type="warning"
            link
            icon="bottom"
            v-if=" scope.row.status === 1"
            @click="changeStatus(scope.row.productId,2)"
          >下架
          </el-button>
          <el-button
            type="danger"
            link
            icon="delete"
            @click="deleteProduct(scope.row.productId)"
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
        :page-sizes="[5, 10, 15]"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>

    <product-form ref="productForm" @success="loadTableData"/>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { productChangeStatusApi, productDeleteApi, productListApi } from '@/api/system/product.js'
import { formatTimestamp } from '@/utils/date.js'
import { ElMessage } from 'element-plus'
import ProductForm from '@/view/business/product/component/productForm.vue'

// 分页相关
const page = ref(1)
const pageSize = ref(10)
const keyWord = ref('')
const total = ref(0)
// 表格数据
const tableData = ref([])
// 弹出层
const productForm = ref(null)


// 加载表格数据
const loadTableData = async() => {
  const res = await productListApi({ page: page.value, pageSize: pageSize.value, keyWord: keyWord.value ,businessId: ''})
  tableData.value = res.data.list
  total.value = res.data.total
}
loadTableData()
// 页数发生变化
const handleCurrentChange = async(val) => {
  page.value = val
  await loadTableData()
}
// 页面大小发生变化
const handleSizeChange = async(val) => {
  pageSize.value = val
  await loadTableData()
}
// 添加商品
const addProduct = () => {
  productForm.value.isEdit = false
  productForm.value.openDialog('添加商品')
}

// 编辑商品
const editProduct = (id) => {
  productForm.value.isEdit = true
  productForm.value.setProductID(id)
  productForm.value.openDialog('编辑商品信息')
}

// 删除商品
const deleteProduct = async(id) => {
  const res = await productDeleteApi({ productId: id })
  if (res.code === 0) {
    ElMessage({
      message: '删除商品成功',
      type: 'success',
    })
    await loadTableData()
  }
}

// 修改商户状态
const changeStatus = async(id, status) => {
  const res = await productChangeStatusApi({ productId: id, status: status })
  if (res.code === 0) {
    ElMessage({
      message: '操作成功',
      type: 'success',
    })
    await loadTableData()
  }
}
</script>

