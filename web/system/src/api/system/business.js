import service from '@/utils/request.js'

export const businessListApi = (data) => {
  return service({
    url: '/system/v1/business/list',
    method: 'post',
    data: data
  })
}
