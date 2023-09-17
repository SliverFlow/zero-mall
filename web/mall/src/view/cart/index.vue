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
          <el-dropdown style="position: relative;top:-1px">
            <span>
              你脑子挖掉了
              <el-icon style="margin-left: 2px;position: relative;top: 2px">
                <arrow-down />
              </el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>个人中心</el-dropdown-item>
                <el-dropdown-item>评价晒单</el-dropdown-item>
                <el-dropdown-item>我的喜欢</el-dropdown-item>
                <el-dropdown-item divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
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
</script>

<style scoped lang="scss">
nav {
  width: 100%;
  background-color: white;
  border-bottom: 2px solid #ff6700;

  .content {
    width: 1220px;
    height: 100px;
    margin: 0 auto;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .left {
      display: flex;
      align-items: center;

      p {
        font-size: 28px;
        line-height: 48px;
      }

      span {
        height: 20px;
        line-height: 20px;
        margin-top: 20px;
        margin-left: 15px;
        color: #757575;
        font-size: 12px;
      }

      img {
        width: 48px;
        height: 48px;
        margin-right: 40px;
      }
    }

    .right {
      display: flex;
      align-items: center;

      .fg {
        margin: 0 10px;
        color: rgba(117, 117, 117, 0.58);
        font-weight: 100;
      }

      p {
        font-size: 14px;
        color: #757575;
        cursor: pointer;
      }
    }
  }
}

.cart_content {
  width: 100%;
  background-color: #f5f5f5;
  display: flex;
  align-items: center;
  flex-direction: column;

  .cart_table {
    width: 1220px;
    background-color: white;
    margin-top: 40px;

    .delete_icon {

      &:hover {
        cursor: pointer;
        color: #ff6700;
      }
    }
  }

  .settlement {
    width: 1220px;
    margin-top: 20px;
    margin-bottom: 40px;
    background-color: white;
    display: flex;
    height: 50px;
    align-items: center;
    justify-content: space-between;

    .left {
      display: flex;
      align-items: center;
      margin-left: 30px;

      p {
        font-size: 14px;
        color: #757575;
      }

      span {
        margin: 0 20px;
        color: rgba(117, 117, 117, 0.62);
        font-weight: 100;
      }
    }

    .right {
      height:100%;
      display: flex;
      align-items: center;
      color: #ff6700;
      p {
        font-size: 14px;
        position: relative;
        top: 4px;
      }

      span {
        font-size: 24px;
      }

      a {
        height: 100%;
        width: 200px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: white;
        background-color: rgba(255, 103, 0, 0.89);
        margin-left: 20px;
        cursor: pointer;
        transition: all 0.3s;
      }
      a:hover {
        background-color: #ff6700;
      }
    }
  }
}

// 下拉框黑边问题
:deep(:focus-visible) {
  outline: none;
}

:deep(.el-dropdown-menu__item) {
  font-size: 12px;
}

:deep(.el-dropdown-menu__item):hover {
  color: #ff6700;
  background-color: rgba(204, 204, 204, 0.17);
}
:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0  1px  #ff6700 inset;
}
:deep(.el-input-number__decrease) {
  background-color: white;
  border-right: none;
}
:deep(.el-input-number__increase) {
  background-color: white;
  border-left: none;
}
:deep(.el-table tbody tr):hover>td {
  background-color: white !important;
}
:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  background-color: #ff6700;
  border-color: #ff6700;
}
:deep(.el-checkbox__inner):hover {
  border-color: #ff6700;
}
:deep(.el-checkbox__input.is-indeterminate .el-checkbox__inner) {
  background-color: #ff6700;
  border-color: #ff6700;
}
:deep(.el-table__header) {
  height: 70px
}
:deep(.cell) {
  font-size: 14px;
  font-weight: 500;
  color: #424242;
}
</style>
