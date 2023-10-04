<template>
  <div>
    <!--轮播图-->
    <div style="width: 1220px;margin: 0 auto" class="mall_carousel_category" @mouseleave="is_hover_category = false">
      <el-carousel :interval="2000" arrow="never" indicator-position="none" style="height: 100%;width: 100%">
        <el-carousel-item v-for="item in carouselItemList" :key="item">
          <el-image style="width: 100%;height: 100%;" :src="item.url" />
        </el-carousel-item>
      </el-carousel>
      <!--侧面分类-->
      <div class="mall_category">
        <template v-for="(v,k) in categoryList">
          <div @click="toProductList(v.categoryId)">
            <p>{{ v.name }}</p>
            <svg
              t="1694692637616"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="3973"
              width="200"
              height="200"
            >
              <path
                d="M761.055557 532.128047c0.512619-0.992555 1.343475-1.823411 1.792447-2.848649 8.800538-18.304636 5.919204-40.703346-9.664077-55.424808L399.935923 139.743798c-19.264507-18.208305-49.631179-17.344765-67.872168 1.888778-18.208305 19.264507-17.375729 49.631179 1.888778 67.872168l316.960409 299.839269L335.199677 813.631716c-19.071845 18.399247-19.648112 48.767639-1.247144 67.872168 9.407768 9.791372 21.984142 14.688778 34.560516 14.688778 12.000108 0 24.000215-4.479398 33.311652-13.439914l350.048434-337.375729c0.672598-0.672598 0.927187-1.599785 1.599785-2.303346 0.512619-0.479935 1.056202-0.832576 1.567101-1.343475C757.759656 538.879828 759.199462 535.391265 761.055557 532.128047z"
                fill="#ffffff"
                p-id="3974"
              />
            </svg>
          </div>
        </template>
      </div>
    </div>
    <!--图片信息-->
    <div style="width: 1220px;margin: 10px auto" class="image-info">
      <div class="info">
        <div>
          <span />
          <span />
        </div>
      </div>
      <div class="image">
        <template v-for="item in infoItemList">
          <el-image style="width: 318px;height: 170px;" :src="item.url" />
        </template>
      </div>
    </div>

    <!--详细商品-->
    <div class="product-detail">
      <img
        src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/1a7f2dc81ede03fdb143e84e11eb550e.png?thumb=1&w=1839&h=180&f=webp&q=90"
        alt=""
      >

      <template v-for="(v,k) in productList">
        <div class="product-info">
          <div class="title">
            <p>{{ v.name }}</p>
            <span @click="toProductList(v.categoryId)">查看更多 >></span>
          </div>
          <div class="detail">
            <img
              v-if="k === 0"
              src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/9613bd4127fede053f8ba5049eff0392.jpeg?thumb=1&w=351&h=921&f=webp&q=90"
            >
            <img
              v-else
              src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/c9550c95ee33a22dc8db677dada00f09.jpg?thumb=1&w=351&h=921&f=webp&q=90"
              alt=""
            >
            <div class="product_list">
              <template v-for="(i,k) in v.productList">
                <div class="product">
                  <img
                    :src="i.image[0].url"
                    style="margin-top: 14px"
                  >
                  <p class="name">{{ i.name }}</p>
                  <p class="desc">{{ i.subtitle }}</p>
                  <span> {{ i.price }} 元起 </span>
                </div>
              </template>
            </div>
          </div>
        </div>
      </template>

      <img
        src="https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/4d1e0c3e71628cf1df6c941d7a908637.jpeg?thumb=1&w=1839&h=180&f=webp&q=90"
        alt=""
      >
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { categoryIDByProductListApi, categoryListApi } from '@/api/category.js'

const router = useRouter()
// 轮播图列表
const carouselItemList = ref([
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7bfda2a88bc0ce1f17b44cce285be2d4.jpg?thumb=1&w=1839&h=690&f=webp&q=90' },
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/da7c749aa3ce4f675ad44d7ca80ad749.jpg?thumb=1&w=1839&h=690&f=webp&q=90' },
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/309d66406483fd2da682e014cc772ef3.jpg?thumb=1&w=1839&h=690&f=webp&q=90' }
])

const infoItemList = ref([
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/8dede2520f8dfff9c9b690af498cafe8.jpg?w=632&h=340' },
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/1ac77590368ff636d0b4f6a988133f55.png?w=632&h=340' },
  { url: 'https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/029f7790f2a2816aa7c0d205b25dcef0.jpeg?w=632&h=340' }
])
// 是否在分类上浮动
const is_hover_category = ref(false)
// 分类列表
const categoryList = ref([])
// productList
const productList = ref([])
const page = ref(1)
const pageSize = ref(8)

const loadCategoryList = async() => {
  const res = await categoryListApi()
  categoryList.value = res.data.list

  const list1 = await categoryIDByProductListApi({
    categoryId: categoryList.value[0].categoryId,
    page: page.value,
    pageSize: pageSize.value,
    keyWord: ''
  })
  const list2 = await categoryIDByProductListApi({
    categoryId: categoryList.value[1].categoryId,
    page: page.value,
    pageSize: pageSize.value,
    keyWord: ''
  })
  productList.value.push(list1.data, list2.data)
}
loadCategoryList()
// 去往商品详情页
const toProductList = (id) => {
  console.log(id)
  router.push({ name: 'ProductList', query: { id: id }})
}
</script>

<style scoped lang="scss">
@import "@/style/index.scss";
</style>
