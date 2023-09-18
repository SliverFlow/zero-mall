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

export const categoryUpdateApi = (data) => {
  return service({
    method: 'post',
    url: '/system/v1/category/update',
    data: data
  })
}
