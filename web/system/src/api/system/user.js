import service from '@/utils/request.js'

export const logoutApi = () => {
  return service({
    url: '/system/user/logout',
    method: 'post',
  })
}
