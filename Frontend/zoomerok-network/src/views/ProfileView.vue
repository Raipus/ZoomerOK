<template>
  <profile-header v-if="user.id" :user="user" :friends="friends" :realLogin="realLogin" />
  <div class="flex justify-center" v-if="userPosts.length === 0">
    <div
      class="w-[610px] min-h-[83px] mt-[24px] bg-[#7500DB] rounded-[40px] grid items-center place-content-center p-[14px]"
    >
      <p
        class="w-[574px] font-bold place-content-center items-center grid px-[16px] min-h-[55px] bg-white rounded-[40px] top-[14px] left-[14px]"
      >
        {{
          user.login === realLogin
            ? 'У Вас ещё нет постов! 🐱‍👤 Пора завести ⤵'
            : 'У этого зумера ещё нет постов! 😁'
        }}
      </p>
    </div>
  </div>
  <post-composer v-if="user.login === realLogin" />
  <div>
    <post-item v-for="post in userPosts" :key="post.id" :post="post" />
    <div class="flex justify-center" v-if="loading">
      <v-progress-circular indeterminate class="my-5" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useProfileStore } from '@/stores/profileStore'
import ProfileHeader from '@/components/ProfileHeader.vue'
import PostComposer from '@/components/PostComposer.vue'
import PostItem from '@/components/PostItem.vue'
import { storeToRefs } from 'pinia'
import type { TokenPayload } from '@/types/tokenType'
import { jwtDecode } from 'jwt-decode'
import Cookies from 'js-cookie'
import { usePostsStore } from '@/stores/postStore'

const route = useRoute()
const login = route.params.login as string

const userStore = useProfileStore()
const postStore = usePostsStore()
let realLogin = ''

onMounted(async () => {
  const token = Cookies.get('access_token') as string
  try {
    const decoded = jwtDecode(token) as TokenPayload
    realLogin = decoded.login
  } catch (error) {
    Cookies.remove('access_token')
    window.location.replace(`${import.meta.env.VITE_AUTH_URL}/signin`)
    throw(error)
  }

  await userStore.fetchUser(login)
  await userStore.fetchFriends()
  await postStore.fetchUserPosts(userStore.user.id, 1)
})

watch(
  () => route.params.login,
  async (newLogin) => {
    if (newLogin) {
      await userStore.fetchUser(newLogin as string)
      await userStore.fetchFriends()
      await postStore.fetchUserPosts(userStore.user.id, 1)
    }
  },
)

const { user, friends } = storeToRefs(userStore)
const { userPosts, loading } = storeToRefs(postStore)
</script>
