<script setup>
import { userLoginApi } from '@/api/user.js'

const login = async() => {
  const res = await userLoginApi({ username: '1212', password: '1212' })
  console.log(res)
}
login()
</script>

<template>
  <div>主页面</div>
</template>

<style scoped>

</style>
