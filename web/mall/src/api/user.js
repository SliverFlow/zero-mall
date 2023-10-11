import service from '@/utils/request.js'

export const userFindApi = () => {
  return service({
    url: '/user/find',
    method: 'post'
  })
}

export const userLoginByPhoneOrUsernameApi = (data) => {
  return service({
    url: '/user/login/phoneOrUsername',
    method: 'post',
    data: data
  })
}

export const userLoginByPhoneApi = (data) => {
  return service({
    url: '/user/login/phone',
    method: 'post',
    data: data
  })
}

export const phoneCaptchaApi = (data) => {
  return service({
    url: '/user/phoneCaptcha',
    method: 'post',
    data: data
  })
}
