export const getOssClient = (params) => {
  return new OSS({
    ...params
  })
}
