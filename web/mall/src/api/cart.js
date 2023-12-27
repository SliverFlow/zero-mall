import service from '@/utils/request.js'

export const cartCreateApi = (data) => {
  return service({
    method: 'post',
    url: '/cart/create',
    data: data
  })
}

export const cartListApi = (data) => {
  return service({
    method: 'post',
    url: '/cart/list',
    data: data
  })
}

export const cartDeleteApi = (data) => {
  return service({
    method: 'post',
    url: '/cart/delete',
    data: data
  })
}
