import service from '@/utils/request.js'

export const logoutApi = () => {
  return service({
    url: '/system/v1/user/logout',
    method: 'post',
  })
}

export const userChangeRoleApi = (data) => {
  return service({
    url: '/system/v1/user/changeRole',
    method: 'post',
    data: data
  })
}

export const userFindByUUIDApi = () => {
  return service({
    url: '/system/v1/user/findByUUID',
    method: 'post',
  })
}

export const userListApi = (data) => {
  return service({
    url: '/system/v1/user/list',
    method: 'post',
    data: data
  })
}

// 用户更新状态
export const userChangeStatusApi = (data) => {
  return service({
    url: '/system/v1/user/changeStatus',
    method: 'post',
    data: data
  })
}

export const userDeleteApi = (data) => {
  return service({
    url: '/system/v1/user/delete',
    method: 'post',
    data: data
  })
}
