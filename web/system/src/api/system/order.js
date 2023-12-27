import request from '@/utils/request.js'

export const orderListApi = (data) => {
  return request({
    url: '/system/v1/order/list',
    method: 'post',
    data
  })
}
