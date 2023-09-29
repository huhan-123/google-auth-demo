import {createApp} from 'vue'
import App from './App.vue'
import router from './router'

import vue3GoogleLogin from 'vue3-google-login'
import {createPinia} from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

const app = createApp(App)


const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(router)
app.use(pinia)
app.use(vue3GoogleLogin, {
    clientId: 'your client_id'
})

app.mount('#app')