import service from '@/utils/request.js'

export const logoutApi = () => {
  return service({
    url: '/system/user/logout',
    method: 'post',
  })
}

export const userChangeRoleApi = (data) => {
  return service({
    url: '/system/user/changeRole',
    method: 'post',
    data: data
  })
}
