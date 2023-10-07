<template>
  <div class="content_container">
    <div class="title"/>
    <!--分类显示-->
    <div class="category_container">
      <div class="category_content">
        <el-collapse accordion>
          <el-collapse-item name="active" class="title-info">
            <template #title>
              <span>首页</span>
              <span>></span>
              <span>全部商品</span>
              <span>></span>
              <p>展开分类 : </p>
            </template>
            <template v-for="(v,k) in categoryList">
              <div class="category_item" style="margin: 2px 0 10px 0">
                <p :class="{active: v.categoryId === categoryId}" @click="reload(v.categoryId)">{{ v.name }}</p>
                <template v-for="(i,k) in v.children">
                  <a
                    :class="{active: i.categoryId === categoryId}"
                    class="secondary"
                    @click="reload(i.categoryId)"
                  >{{ i.name }}</a>
                </template>
              </div>
            </template>
          </el-collapse-item>
        </el-collapse>
      </div>
    </div>
    <!--筛选条件-->
    <div class="screen">
      <div class="content">
        <p>综合</p>
        <span>|</span>
        <p>销量</p>
        <span>|</span>
        <p>价格</p>
      </div>
    </div>
    <!--商品列表-->
    <div class="product_list">
      <div class="content">
        <template v-for="(v) in productList">
          <div class="prod" @click="toProductDetail(v.productId)">
            <img
              :src="v.image[0].url"
              alt=""
            >
            <el-text style="width: 240px;display: flex;align-items: center;justify-content: center" truncated>{{
                v.name
              }}
            </el-text>
            <span class="price">{{ v.price }}元</span>
            <ul>
              <template v-for="(i,k) in v.image">
                <li>
                  <img
                    :src="i.url"
                    alt=""
                  >
                </li>
              </template>
            </ul>
            <div class="tag">
              <span>分期</span>
            </div>
          </div>
        </template>
      </div>
    </div>
    <!--分页-->
    <div class="pagination">
      <el-pagination
        class="page_con"
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
import { inject, onBeforeMount, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { categoryIDByProductListApi, categoryListApi } from '@/api/category.js'
import { useActiveStore } from '@/pinia/model/active.js'

const keyWord = inject('keyWord')
const activeStore = useActiveStore()
const router = useRouter()
const route = useRoute()
const categoryId = ref(route.query.id)
const pageSize = ref(8)
const page = ref(1)
const total = ref(0)
const productList = ref([])
const categoryList = ref([])

const handleCurrentChange = (val) => {
  page.value = val
  loadData()
}
const handleSizeChange = (val) => {
  pageSize.value = val
  loadData()
}

const toProductDetail = (id) => {
  router.push({ name: 'ProductDetail', query: { id: id } })
}

const loadData = async() => {
  if (!categoryId.value) {
    categoryId.value = ''
  }
  const res = await categoryIDByProductListApi({
    categoryId: categoryId.value,
    page: page.value,
    pageSize: pageSize.value,
    keyWord: keyWord.value
  })
  productList.value = res.data.productList
  total.value = res.data.total
}
loadData()
const loadCategoryList = async() => {
  const res = await categoryListApi()
  categoryList.value = res.data.list
}

const reload = async(id) => {
  categoryId.value = id
  await loadData()
}

onBeforeMount(() => {
  loadCategoryList()
})

watch(() => activeStore.active, async(val, old) => {
  if (val) {
    await loadData()
    activeStore.active = false
  }
})
</script>

<style scoped lang="scss">
@import "@/style/product.scss";

.title-info {
  span {
    color: #5a5d5a;
    margin-right: 14px;
  }
}

.active {
  color: #ff6700;
}
</style>
