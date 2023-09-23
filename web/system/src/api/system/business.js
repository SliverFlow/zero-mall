import service from '@/utils/request.js'

export const businessListApi = (data) => {
  return service({
    url: '/system/v1/business/list',
    method: 'post',
    data: data
  })
}

// 修改商户状态
export const businessChangeStatusApi = (data) => {
  return service({
    url: '/system/v1/business/changeStatus',
    method: 'post',
    data: data
  })
}
