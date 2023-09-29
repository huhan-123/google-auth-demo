import axios from 'axios'
import {useUserStore} from "@/stores/userStore";

// 创建axios实例
const httpInstance = axios.create({
    baseURL: 'http://localhost:8888',
    timeout: 5000
})

// axios请求拦截器
httpInstance.interceptors.request.use(config => {
    const userStore = useUserStore()
    const token = userStore.token
    console.log(token)
    if (token) {
        config.headers.Authorization = `Bearer ${token}`
    }
    return config
}, e => Promise.reject(e))

// axios响应式拦截器
httpInstance.interceptors.response.use(res => res.data, e => {
    return Promise.reject(e)
})


export default httpInstance