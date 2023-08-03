import service from '@/utils/request.js'

export const menuTreeListApi = () => {
  return service({
    url: '/system/menu/treeList',
    method: 'post'
  })
}

export const menuTreeListAllApi = () => {
  return service({
    url: '/system/menu/treeListAll',
    method: 'post'
  })
}

export const menuChangeStatusApi = (data) => {
  return service({
    url: '/system/menu/changeStatus',
    method: 'post',
    data: data
  })
}

export const menuUpdateApi = (data) => {
  return service({
    url: '/system/menu/update',
    method: 'post',
    data: data
  })
}
