<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavBar from './components/NavBar.vue'
import Cookies from 'js-cookie'
import { onMounted } from 'vue'
import { jwtDecode } from 'jwt-decode'
import type { TokenPayload } from '@/types/tokenType'
import { useAuthStore } from './stores/authStore'

const authStore = useAuthStore()

onMounted(async () => {
  const token = Cookies.get('access_token') as string
  try {
    const decoded = jwtDecode(token) as TokenPayload
    authStore.setLogin(decoded.login)
    authStore.fetchUserData()
  } catch (error) {
    Cookies.remove('access_token')
    window.location.replace(`${import.meta.env.VITE_AUTH_URL}/signin`)
    throw(error)
  }
})
</script>

<template>
  <NavBar :userStore="authStore" />
  <RouterView />
  <div class="flex justify-center my-10">© Zoomers Team 2025</div>
</template>
