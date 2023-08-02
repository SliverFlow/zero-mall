import service from '@/utils/request.js'

export const menuTreeListApi = () => {
  return service({
    url: '/system/menu/treeList',
    method: 'post'
  })
}
