import axios from 'axios'

const service = axios.create()

// 获取 github 仓库的提交记录
export const comments = () => {
  return service({
    url: 'https://api.github.com/repos/SliverFlow/zpshop/commits?page=1',
    method: 'get'
  })
}
