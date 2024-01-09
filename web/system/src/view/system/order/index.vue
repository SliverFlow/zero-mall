<template>
  <div class="gva-table-box">
    <div class="gva-btn-list" style="margin-bottom: 24px">
      <div style="width: 300px;margin-right:40px;display: flex;align-items: center">
        <span>店铺：</span>
        <el-select
          placeholder="请选择"
          v-model="businessId"
          @change="loadTableData"
        >
          <el-option
            v-for="item in  businessDict"
            :key="item.value"
            :label="item.label"
            :value="item.value"/>
        </el-select>
      </div>
      <div style="width: 300px;margin-right:40px;display: flex;align-items: center">
        <span>订单状态：</span>
        <el-select
          placeholder="请选择"
          v-model="status"
          @change="loadTableData"
        >
          <el-option
            v-for="item in statusList"
            :key="item.value"
            :label="item.label"
            :value="item.value"/>
        </el-select>
      </div>
      <div style="width: 350px;margin-right:40px;display: flex;align-items: center">
        <span>时间：</span>
        <el-date-picker
          @change="loadTableData"
          v-model="timeArr"
          type="daterange"
          range-separator="-"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          value-format="YYYY-MM-DD HH:mm:ss"
          style="width: 240px;"
        />
      </div>
      <el-button
        type="primary"
        @click="resetQuery"
        :icon="Refresh">
        重置
      </el-button>
    </div>
    <el-table
      :data="tableData"
      style="z-index: 0;height: calc(100vh - 284px)"
      row-key="orderId"
    >
      <el-table-column align="left" label="订单编号" min-width="220" prop="orderId" fixed="left"/>
      <el-table-column align="left" label="购买者" min-width="150" prop="username"/>
      <el-table-column align="left" label="商品名称" min-width="220" prop="orderId">
        <template #default="scope">
          <el-text type="primary">{{ scope.row.productName }}</el-text>
        </template>
      </el-table-column>
      <el-table-column align="left" label="商家名称" min-width="220" prop="businessName"/>
      <el-table-column align="left" label="商品图片" min-width="150">
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

      <el-table-column align="left" label="商品数量" min-width="150" prop="quantity">
        <template #default="scope">
          {{ scope.row.quantity + ' 件' }}
        </template>
      </el-table-column>
      <el-table-column align="left" label="支付金额" min-width="150" prop="payment">
        <template #default="scope">
          {{ scope.row.payment.toFixed(2) + ' 元' }}
        </template>
      </el-table-column>
      <el-table-column align="left" label="订单状态" min-width="150" prop="orderStatus">
        <template #default="scope">
          <el-tag v-if="scope.row.orderStatus === 70" type="info">已取消</el-tag>
          <el-tag v-if="scope.row.orderStatus === 10" type="danger">未付款</el-tag>
          <el-tag v-if="scope.row.orderStatus === 20" type="info">已付款</el-tag>
          <el-tag v-if="scope.row.orderStatus === 30" type="info">代发货</el-tag>
          <el-tag v-if="scope.row.orderStatus === 40" type="info">待收货</el-tag>
          <el-tag v-if="scope.row.orderStatus === 50" type="success">交易完成</el-tag>
          <el-tag v-if="scope.row.orderStatus === 60" type="success">交易关闭</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="left" label="订单创建时间" min-width="220">
        <template #default="scope">
          {{ formatTimestamp(scope.row.createTime) }}
        </template>
      </el-table-column>
      <el-table-column align="left" label="订单完成时间" min-width="220" prop="endTime">
        <template #default="scope">
          {{
            formatTimestamp(scope.row.endTime) === '1970-01-01 08:00:00' ? '未完成' : formatTimestamp(scope.row.endTime)
          }}
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="160" fixed="right">
        <template #default="scope">
          <el-button
            type="primary"
            link
          >商家信息
          </el-button>
          <el-button
            type="primary"
            link
          >用户信息
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
import { onMounted, ref } from 'vue'
import { orderListApi } from '@/api/system/order.js'
import { formatTimestamp } from '@/utils/date.js'
import { Delete, Refresh, Search } from '@element-plus/icons-vue'
import { businessDictApi, businessFindApi } from '@/api/system/business.js'

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const status = ref('0')
const businessId = ref('')
const timeArr = ref([])
const businessDict = ref([])
const statusList = ref([
  { label: '全部', value: '0' },
  { label: '已取消', value: '70' },
  { label: '未付款', value: '10' },
  { label: '已付款', value: '20' },
  { label: '待发货', value: '30' },
  { label: '待收货', value: '40' },
  { label: '交易完成', value: '50' },
  { label: '交易关闭', value: '60' },
])



const resetQuery = () => {
  page.value = 1
  pageSize.value = 10
  status.value = '0'
  businessId.value = ''
  timeArr.value = []
  loadTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  loadTableData()
}
const handleSizeChange = (val) => {
  pageSize.value = val
  loadTableData()
}

const loadTableData = async() => {

  const data = {
    page: page.value,
    pageSize: pageSize.value,
    businessId: businessId.value,
    status: parseInt(status.value),
  }

  if (timeArr.value.length > 0) {
    data.startTime = timeArr.value[0]
    data.endTime = timeArr.value[1]
  } else {
    data.startTime = ''
    data.endTime = ''
  }

  const res = await orderListApi(data)
  if (res.code !== 0) return
  tableData.value = res.data.list
  total.value = res.data.total
}

const loadBusinessDict = async() => {
  const res = await businessDictApi()
  if (res.code !== 0) return
  businessDict.value = res.data.list
}


onMounted(async() => {
  await loadBusinessDict()
  await loadTableData()
})


</script>

<style scoped lang="scss">

</style>
