import service from '@/utils/request.js'

// 分类列表 分页查询
export const categoryListAllApi = () => {
  return service({
    method: 'post',
    url: '/system/v1/category/listAll',
  })
}

// 添加分类
export const categoryCreateApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/category/create',
    data: data
  })
}

// 更新分类状态
export const categoryChangeStatus = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/category/changeStatus',
    data: data
  })
}

// 更新分类
export const categoryUpdateApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/category/update',
    data: data
  })
}

// 查询分类信息
export const categoryFindApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/category/find',
    data: data
  })
}

// 批量删除
export const categoryBatchDeleteApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/category/batchDelete',
    data: data
  })
}

// 指定删除
export const categoryDeleteApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/category/delete',
    data: data
  })
}
