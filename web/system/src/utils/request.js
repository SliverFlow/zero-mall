import axios from 'axios'
import { useUserStore } from '@/store/model/user.js'
import { ElMessage, ElMessageBox } from 'element-plus'
import router from '@/router/index.js'

const service = axios.create({
  timeout: 10000,
  baseURL: import.meta.env.VITE_BASE_API
})

// 请求拦截
service.interceptors.request.use(
  config => {
    const userStore = useUserStore()
    config.headers = {
      'Authorization': userStore.token,
      ...config.headers
    }
    return config
  },
  error => {
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// 响应拦截
service.interceptors.response.use(
  response => {
    const userStore = useUserStore()
    // 头信息里面有 token 就替换原来的 token
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    return response.data
  },
  error => {
    // 请求出错
    if (!error.response) {
      ElMessageBox.confirm(`
        <p>检测到请求错误</p>
        <p>${error}</p>
        `, '请求报错', {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: '稍后重试',
        cancelButtonText: '取消'
      })
      return
    }
    // 500 和 404
    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(`
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `, '接口报错', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '清理缓存',
          cancelButtonText: '取消'
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.setToken('')
            localStorage.clear()
            router.push({ name: 'Login', replace: true })
          })
        break
      case 404:
        ElMessageBox.confirm(`
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `, '接口报错', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '我知道了',
          cancelButtonText: '取消'
        })
        break
    }
    return error
  }
)

export default service
