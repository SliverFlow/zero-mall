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

export const userFindByUUID = () => {
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
