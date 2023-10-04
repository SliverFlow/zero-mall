import service from '@/utils/request.js'

export const categoryListApi = () => {
  return service({
    method: 'post',
    url: '/category/list',
  })
}

export const categoryIDByProductListApi = (data) => {
  return service({
    method: 'post',
    url: '/category/productList',
    data: data
  })
}
