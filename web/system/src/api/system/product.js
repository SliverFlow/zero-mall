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
export const productDeleteApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/product/delete',
    data: data
  })
}

// 商品更新状态
export const productChangeStatusApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/product/changeStatus',
    data: data
  })
}

// 商品创建
export const productCreateApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/product/create',
    data: data
  })
}

// 商品查询
export const productFindApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/product/find',
    data: data
  })
}

// 商品更新
export const productUpdateApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/product/update',
    data: data
  })
}
