import service from '@/utils/request.js'

export const userLoginApi = (data) => {
  return service({
    url: '/user/login',
    method: 'post',
    data
  })
}

const userPhoneCaptchaApi = (data) => {
  return service({
    url: '/user/phoneCaptcha',
    method: 'post',
    data: data
  })
}
