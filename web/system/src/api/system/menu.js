import service from '@/utils/request.js'

export const menuTreeListApi = () => {
  return service({
    url: '/system/menu/treeList',
    method: 'post'
  })
}

export const menuTreeListAllApi = (data) => {
  return service({
    url: '/system/menu/treeListAll',
    method: 'post',
    data: data
  })
}

export const menuChangeStatusApi = (data) => {
  return service({
    url: '/system/v1/menu/changeStatus',
    method: 'post',
    data: data
  })
}

export const menuUpdateApi = (data) => {
  return service({
    url: '/system/v1/menu/update',
    method: 'post',
    data: data
  })
}

export const menuCreateApi = (data) => {
  return service({
    url: '/system/v1/menu/create',
    method: 'post',
    data: data
  })
}

export const menuFindApi = (data) => {
  return service({
    url: '/system/v1/menu/find',
    method: 'post',
    data: data
  })
}
