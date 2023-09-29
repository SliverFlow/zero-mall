import service from '@/utils/request.js'

// 商品分页 携带分类
export const productListApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/product/list',
    data: data
  })
}

// 商品删除
export const productDeleteApi = (data) =>  {
  return service({
    method: 'post',
    url: '/system/v1/product/delete',
    data: data
  })
}
