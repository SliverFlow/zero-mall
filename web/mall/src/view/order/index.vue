<template>
  <div>
    <nav>
      <div class="content">
        <div class="left">
          <img src="https://cdn.cnbj1.fds.api.mi-img.com/mi.com-assets/shop/img/logo-mi2.png">
          <p>我的订单</p>
          <span>温馨提示：与你相关的订单将全部展示在这里哦</span>
        </div>
        <div class="right">
          <user-info-com v-if="isLogin" />
          <span class="fg">|</span>
          <p @click="toCart">购物车</p>
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
              <el-text style="width: 380px" size="large" truncated>{{ scope.row.price }}元</el-text>
            </template>
          </el-table-column>
          <el-table-column property="quantity" label="数量" width="200px" align="center">
            <template #default="scope">
              {{ scope.row.quantity }} 件
            </template>
          </el-table-column>
          <el-table-column property="xj" label="支付金额" align="center">
            <template #default="scope">
              <el-text style="color: #ff6700;" size="large">{{ scope.row.payment }}元</el-text>
            </template>
          </el-table-column>
          <el-table-column property="xj" label="状态" align="center">
            <template #default="scope">
              <el-tag v-if="scope.row.orderStatus === 70" type="info">已取消</el-tag>
              <el-tag v-if="scope.row.orderStatus === 10" type="danger">未付款</el-tag>
              <el-tag v-if="scope.row.orderStatus === 20" type="info">已付款</el-tag>
              <el-tag v-if="scope.row.orderStatus === 30" type="info">代发货</el-tag>
              <el-tag v-if="scope.row.orderStatus === 40" type="info">待收货</el-tag>
              <el-tag v-if="scope.row.orderStatus === 50" type="success">交易完成</el-tag>
              <el-tag v-if="scope.row.orderStatus === 60" type="success">交易关闭</el-tag>
              <el-tag v-if="scope.row.orderStatus === 80" type="success">订单完成</el-tag>
              <el-button
                v-if="scope.row.orderStatus === 10"
                link
                class="qx_btn"
                @click="disableOrder(scope.row.orderId)"
              >取消订单
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="settlement" style="margin-bottom: 0">
        <div class="left">
          <p style="cursor: pointer;" @click="() => {router.push({name: 'Index'})}">继续购物</p>
          <span>|</span>
          <p>已选择 <span style="color: #ff6700;margin: 0 4px">0</span> 个订单</p>
        </div>
        <div class="right">
          <span>温馨提示：此操作不会删除以付款但未完成的订单</span>
          <a :class="{isDelete: selectOrderList.length === 0}" @click="deleteOrderItem">删除所选订单</a>
        </div>
      </div>
      <!--分页-->
      <div class="pagination-order">
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
    <!--尾部-->
    <mall-footer />
  </div>
</template>

<script setup>
import MallFooter from '@/view/layout/footer/mallFooter.vue'
import { ref } from 'vue'
import UserInfoCom from '@/components/userInfo/UserInfoCom.vue'
import { useUserStore } from '@/pinia/model/user.js'
import { useRouter } from 'vue-router'
import { orderDeleteApi, orderDisableApi, orderListApi } from '@/api/order.js'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()
const isLogin = userStore.isLogin
const tableData = ref([])

const router = useRouter()
const page = ref(1)
const pageSize = ref(8)
const total = ref(0)
const selectOrderList = ref([])
const orderIds = ref([])

const toCart = () => {
  router.push({ name: 'Cart' })
}

// 多选框回调
const handleSelectionChange = (e) => {
  selectOrderList.value = e
}

// 删除购物车商品
const deleteOrderItem = async () => {
  if (selectOrderList.value.length === 0) {
    ElMessage.warning('请选择要删除的订单')
    return
  }
  selectOrderList.value.forEach((v) => {
    orderIds.value.push(v.orderId)
  })
  const res = await orderDeleteApi({
    ids: orderIds.value
  })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    await loadData()
  }
}

// 取消订单
const disableOrder = async(val) => {
  const res = await orderDisableApi({
    orderId: val,
  })
  if (res.code === 0) {
    ElMessage.success('取消成功')
    await loadData()
  }
}

const handleCurrentChange = (val) => {
  page.value = val
  loadData()
}
const handleSizeChange = (val) => {
  pageSize.value = val
  loadData()
}
const loadData = async() => {
  const res = await orderListApi({
    page: page.value,
    pageSize: pageSize.value,
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

loadData()
</script>

<style scoped lang="scss">
@import '@/style/cart.scss';

.right {

  span {
    color: #757575;
    font-size: 12px !important;
    margin-top: 14px !important;
  }
}

.pagination-order {
  box-sizing: border-box;
  padding: 20px 0;
  margin: 20px 0;
  background-color: white;
  width: 1220px;
  display: flex;
  align-items: center;
  justify-content: center;

}

:deep(.el-collapse-item__content) {
  padding-bottom: 8px;
}

:deep(.el-pager li.is-active) {
  color: #ff6700;
}

:deep(.el-select .el-input.is-focus .el-input__wrapper) {
  box-shadow: 0 0 0 1px #ff6700 inset !important;
}

:deep(.el-input__wrapper) {
  background-color: white;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #ff6700 inset !important;
}

:deep(.el-pager li) {
  background-color: white;
  margin-left: 0.25rem;
  margin-right: 0.25rem;
  border-radius: 0.25rem;
  border-width: 1px;
  border-style: solid;
  --tw-border-opacity: 1;
  border-color: rgb(209 213 219 / var(--tw-border-opacity));
  font-size: .875rem;
  line-height: 1.25rem;
  --tw-text-opacity: 1;
  color: rgb(75 85 99 / var(--tw-text-opacity));
  transition: all 0.2s;
}

:deep(.el-pager li):hover {
  color: #ff6700;
}

:deep(.btn-prev), :deep(.btn-next) {
  background-color: white !important;
  margin-left: 0.25rem;
  margin-right: 0.25rem;
  border-radius: 0.25rem;
  border-width: 1px;
  border-style: solid;
  --tw-border-opacity: 1;
  border-color: rgb(209 213 219 / var(--tw-border-opacity));
  font-size: .875rem;
  line-height: 1.25rem;
  --tw-text-opacity: 1;
  color: rgb(75 85 99 / var(--tw-text-opacity));
  transition: all 0.2s;
}

:deep(.btn-prev):hover, :deep(.btn-next):hover {
  border-color: #ff6700;
  color: #ff6700;
}

.qx_btn {
  color: #ff6700 !important;
  margin-left: 4px;

  &:hover {
    color: rgba(255, 103, 0, 0.76) !important;
  }
}

.isDelete {
  background-color: #9f9d9d !important;
  cursor: text !important;
  pointer-events: none;
  color: white !important;
}
</style>
