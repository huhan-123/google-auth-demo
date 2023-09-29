import {defineStore} from 'pinia'
import {ref} from "vue";

export const useUserStore = defineStore('user', () => {
    const token = ref("")

    const userInfo = ref({})

    const setUserInfo = (user) => {
        userInfo.value = user
    }

    const clearUserInfo = () => {
        token.value = ""
        userInfo.value = {}
    }

    return {
        token,
        userInfo,
        setUserInfo,
        clearUserInfo
    }
}, {
    persist: {
        key: 'user_info',
        paths: ['token', 'userInfo'],
        storage: localStorage
    }
})