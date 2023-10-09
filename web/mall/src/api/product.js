import service from '@/utils/request.js'

export const productFindApi = (data) => {
  return service({
    url: '/product/find',
    method: 'post',
    data: data
  })
}
