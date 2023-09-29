// 封装所有和用户相关的接口函数
import request from '@/utils/http'

export const loginAPI = (credential) => {
    return request({
        url: '/login',
        method: 'POST',
        data: {
            credential
        }
    })
}

export const userInfoAPI = () => {
    return request({
        url: '/user/info'
    })
}

export const logoutAPI = () => {
    return request({
        url: '/logout',
        method: 'POST'
    })
}