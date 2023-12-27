import service from '@/utils/request.js'

export const orderListApi = (data) => {
  return service({
    method: 'post',
    url: '/order/list',
    data: data
  })
}

export const orderCreateApi = (data) => {
  return service({
    method: 'post',
    url: '/order/create',
    data: data
  })
}

export const orderDisableApi = (data) => {
  return service({
    method: 'post',
    url: '/order/disable',
    data: data
  })
}

export const orderDeleteApi = (data) => {
  return service({
    method: 'post',
    url: '/order/delete',
    data: data
  })
}
