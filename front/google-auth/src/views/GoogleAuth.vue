<script setup>
import {loginAPI, logoutAPI, userInfoAPI} from '@/apis/user'
import {useUserStore} from "@/stores/userStore";

const userStore = useUserStore()
const doLogin = async (response) => {

  const res = await loginAPI(response.credential)
  userStore.token = res.data.token


  const userInfo = await userInfoAPI()

  userStore.setUserInfo(userInfo.data)
}
const doLogout = async () => {
  await logoutAPI()

  userStore.clearUserInfo()

}
</script>

<template>
  <!--    popup mode-->
  <GoogleLogin :callback="doLogin"/>

  <!--   redirect mode-->
  <!--  <GoogleLogin-->
  <!--      :id-configuration="{-->
  <!--      login_uri: 'https://localhost:8888/login',-->
  <!--      ux_mode: 'redirect',-->
  <!--    }"-->
  <!--      :callback="callback"-->
  <!--  />-->
  <br>
  <el-button type="info" @click="doLogout">logout</el-button>
</template>

<style scoped>

</style>