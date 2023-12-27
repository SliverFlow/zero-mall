<template>
  <div>
    <nav>
      <div class="content">
        <div class="left">
          <img src="https://cdn.cnbj1.fds.api.mi-img.com/mi.com-assets/shop/img/logo-mi2.png">
          <p>我的购物车</p>
          <span>温馨提示：产品是否购买成功，以最终下单为准哦，请尽快结算</span>
        </div>
        <div class="right">
          <user-info-com v-if="isLogin"/>
          <span class="fg">|</span>
          <p @click="toOrder">我的订单</p>
        </div>
      </div>
    </nav>
    <!--购物车内容主题-->
    <div class="cart_content">
      <div class="cart_table">
        <el-table
          ref="multipleTableRef"
          :data="tableData"
          style="width: 100%"

          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="封面" align="center" width="120">
            <template #default="scope">
              <el-carousel :interval="3000" arrow="never" indicator-position="none" style="height: 100px;width: 100px">
                <el-carousel-item v-for="(v,k) in scope.row.productImage">
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
                  {{ v.url }}
                </el-carousel-item>
              </el-carousel>

            </template>
          </el-table-column>
          <el-table-column property="name" label="商品名称" width="380px" align="center">
            <template #default="scope">
              <el-text style="width: 380px" size="large" truncated>{{ scope.row.productName }}</el-text>
            </template>
          </el-table-column>
          <el-table-column property="price" label="价格" align="center">
            <template #default="scope">
              <el-text style="width: 380px" size="large" truncated>{{ scope.row.productPrice }}元</el-text>
            </template>
          </el-table-column>
          <el-table-column property="quantity" label="数量" width="200px" align="center">
            <template #default="scope">
              <el-input-number v-model="scope.row.quantity" :min="1" :max="10" style="height: 38px"/>
            </template>
          </el-table-column>
          <el-table-column property="xj" label="小计" align="center">
            <template #default="scope">
              <el-text style="color: #ff6700;" size="large">{{ scope.row.price }}元</el-text>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="80px" align="center">
            <template #default="scope">
              <span class="delete_icon" @click="deleteCartItem(scope.row.cartId)">
                <el-icon size="22"><CircleClose/></el-icon>
              </span>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="settlement">
        <div class="left">
          <p style="cursor: pointer;" @click="() => {router.push({name: 'Index'})}">继续购物</p>
          <span>|</span>
          <p>已选择 <span style="color: #ff6700;margin: 0 4px">0</span> 件商品</p>
        </div>
        <div class="right">
          <p>合计：</p>
          <span>{{ totalPrice }}</span>
          <p>&nbsp;元</p>
          <a>去结算</a>
        </div>
      </div>
    </div>
    <!--尾部-->
    <mall-footer/>
  </div>
</template>

<script setup>
import MallFooter from '@/view/layout/footer/mallFooter.vue'
import { ref } from 'vue'
import UserInfoCom from '@/components/userInfo/UserInfoCom.vue'
import { useUserStore } from '@/pinia/model/user.js'
import { useRouter } from 'vue-router'
import { cartDeleteApi, cartListApi } from '@/api/cart.js'
import { ElMessage, ElMessageBox } from 'element-plus'

const tableData = ref([])

const router = useRouter()
const userStore = useUserStore()
const isLogin = userStore.isLogin
const totalPrice = ref(0)

const toOrder = () => {
  router.push({ name: 'Order' })
}

// 多选框回调
const handleSelectionChange = (e) => {
  if (e.length === 0) {
    totalPrice.value = 0
    return
  }
  e.forEach((v) => {
    totalPrice.value += v.price
  })
}

// 删除购物车商品
const deleteCartItem = async(id) => {
  await ElMessageBox.confirm(
    '确定将这个商品从购物车中移除吗?',
    'Warning',
    {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
  // 删除购物车信息
  const [res] = await Promise.all([cartDeleteApi({ cartId: id })])
  if (res.code === 0) {
    ElMessage.success('商品移除成功')
    await loadTableData()
  }
}

const loadTableData = async() => {
  const [res] = await Promise.all([cartListApi({
    page: 1,
    pageSize: 9999,
    keyWord: ''
  })])
  if (res.code === 0) {
    tableData.value = res.data.list
  }
}
loadTableData()
</script>

<style scoped lang="scss">
@import '@/style/cart.scss';
</style>
