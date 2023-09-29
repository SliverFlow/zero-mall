import service from '@/utils/request.js'

// 头像上传
export const avatarFileTokenApi = () => {
  return service({
    method: 'post',
    url: '/system/v1/file/token/avatar'
  })
}

// 商品图片上传
export const productFileTokenApi = () => {
  return service({
    method: 'post',
    url: '/system/v1/file/token/product'
  })
}
