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
          <p>我的订单</p>
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
          <el-table-column type="selection" width="55" />
          <el-table-column label="封面" align="center">
            <template #default="scope">
              <el-image style="width: 100px; height: 100px" :src="scope.row.url" :fit="fit" />
            </template>
          </el-table-column>
          <el-table-column property="name" label="商品名称" width="380px" align="center">
            <template #default="scope">
              <el-text style="width: 380px" size="large" truncated>{{ scope.row.name }}</el-text>
            </template>
          </el-table-column>
          <el-table-column property="price" label="价格" align="center">
            <template #default="scope">
              <el-text style="width: 380px" size="large" truncated>{{ scope.row.price }}元</el-text>
            </template>
          </el-table-column>
          <el-table-column property="quantity" label="数量" width="200px" align="center">
            <template #default="scope">
              <el-input-number v-model="scope.row.quantity" :min="1" :max="10" style="height: 38px" />
            </template>
          </el-table-column>
          <el-table-column property="xj" label="小计" align="center">
            <template #default="scope">
              <el-text style="color: #ff6700;" size="large">{{ scope.row.xj }}元</el-text>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="80px" align="center">
            <template #default="scope">
              <span class="delete_icon" @click="deleteCartItem(scope.row.id)">
                <el-icon size="22"><CircleClose /></el-icon>
              </span>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="settlement">
        <div class="left">
          <p>继续购物</p>
          <span>|</span>
          <p>已选择 <font style="color: #ff6700">0</font> 件商品</p>
        </div>
        <div class="right" >
          <p>合计：</p>
          <span>9.9</span>
          <p>&nbsp;元</p>
          <a>去结算</a>
        </div>
      </div>
    </div>
    <!--尾部-->
    <mall-footer />
  </div>
</template>

<script setup>
import MallFooter from '@/view/layout/footer/mallFooter.vue'
import { ref } from 'vue'
import UserInfoCom from '@/components/userInfo/UserInfoCom.vue'
import { useUserStore } from '@/pinia/model/user.js'

const tableData = ref([
  {
    name: '小米巨能写中性笔10支装 黑色 10支装',
    price: 9.9,
    url: 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-shopapi-pms/pms_1559616366.16874615.jpg?thumb=1&w=120&h=120',
    quantity: 1,
    xj: 9.9
  },
  {
    name: '小米巨能写中性笔10支装 黑色 10支装',
    price: 9.9,
    url: 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-shopapi-pms/pms_1559616366.16874615.jpg?thumb=1&w=120&h=120',
    quantity: 1,
    xj: 9.9
  },
  {
    name: '小米巨能写中性笔10支装 黑色 10支装',
    price: 9.9,
    url: 'https://cdn.cnbj0.fds.api.mi-img.com/b2c-shopapi-pms/pms_1559616366.16874615.jpg?thumb=1&w=120&h=120',
    quantity: 1,
    xj: 9.9
  }
])
// 多选框回调
const handleSelectionChange = (e) => {

}

// 删除购物车商品
const deleteCartItem = (id) => {
}

//
const userStore = useUserStore()
const isLogin = userStore.isLogin
</script>

<style scoped lang="scss">
@import '@/style/cart.scss';
</style>
