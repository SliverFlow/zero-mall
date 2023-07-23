import service from '@/utils/request.js'

export const loginApi = (data) => {
  return service({
    url: '/system/base/login',
    data: data,
    method: 'post'
  })
}

export const captchaApi = () => {
  return service({
    url: '/system/base/captcha',
    method: 'post'
  })
}
