<script lang="ts" setup>
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useProfileStore } from '@/stores/profileStore'

const profileStore = useProfileStore()

onMounted(async () => {
  await profileStore.fetchFriends()
  await profileStore.fetchUnacceptedFriends()
})

const { friends, unacceptedFriends } = storeToRefs(profileStore)

const acceptFriend = async (userId: number) => {
  try {
    const payload = {
      friend_user_id: userId,
    }
    await profileStore.acceptFriend(payload)
  } catch (error) {
    console.error('Ошибка принятия запроса в друзья:', error)
  }
}

const deleteFriend = async (userId: number) => {
  try {
    const payload = {
      friend_user_id: userId,
    }
    await profileStore.deleteFriend(payload)
  } catch (error) {
    console.error('Ошибка принятия запроса в друзья:', error)
  }
}
</script>

<template>
  <p class="flex justify-center m-7 text-2xl font-bold">Заявки в друзья:</p>
  <div class="flex justify-center">
    <div class="grid grid-cols-3 w-[1118px] justify-items-center gap-y-10">
      <div v-for="friend in unacceptedFriends" :key="friend.friend.id">
        <div class="w-[220px] h-max p-4 bg-[#7500DB] rounded-[15px] hover:scale-110 duration-200">
          <div class="grid justify-items-center bg-white rounded-[15px] p-3 gap-2">
            <router-link
              :to="`/profile/${friend.friend.login}`"
              class="grid gap-2 justify-items-center hover:scale-110 duration-200"
            >
              <img
                :src="`data:image/jpg;base64,${friend.friend.image}`"
                alt="avatar"
                class="rounded-full size-[100px]"
              />
              <p>{{ friend.friend.name }}</p>
            </router-link>
            <button
              type="button"
              @click="acceptFriend(friend.friend.id)"
              class="hover:scale-110 duration-200 h-max w-[160px] bg-green-700 text-white rounded-[15px] p-2"
            >
              Принять запрос дружбы
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <p class="flex justify-center m-7 text-2xl font-bold">Ваши друзья:</p>
  <div class="flex justify-center">
    <div class="grid grid-cols-3 w-[1118px] justify-items-center gap-y-10">
      <div v-for="friend in friends" :key="friend.friend.id">
        <div class="w-[220px] h-max p-4 bg-[#7500DB] rounded-[15px] hover:scale-110 duration-200">
          <div class="grid justify-items-center bg-white rounded-[15px] p-3 gap-2">
            <router-link
              :to="`/profile/${friend.friend.login}`"
              class="grid gap-2 justify-items-center hover:scale-110 duration-200"
            >
              <img
                :src="`data:image/jpg;base64,${friend.friend.image}`"
                alt="avatar"
                class="rounded-full size-[100px]"
              />
              <p>{{ friend.friend.name }}</p>
            </router-link>
            <button
              type="button"
              @click="deleteFriend(friend.friend.id)"
              class="hover:scale-110 duration-200 h-max w-[160px] bg-red-500 text-white rounded-[15px] p-2"
            >
              Удалить друга
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
