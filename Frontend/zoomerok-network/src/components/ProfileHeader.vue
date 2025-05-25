<script setup lang="ts">
import { ref } from 'vue'
import ChangeUserForm from '@/components/ChangeUserForm.vue'
import { computed, defineProps } from 'vue'
import { useProfileStore, type Friend, type UserProfile } from '@/stores/profileStore'

const profileStore = useProfileStore()

const props = defineProps<{ user: UserProfile; realLogin: string; friends: Friend[] }>()

const formattedPhone = computed(() => {
  if (!props.user.phone) return 'не указано'
  return `${props.user.phone}`
})

const formattedCity = computed(() => {
  if (!props.user.city) return 'не указано'
  return `${props.user.city}`
})

const formattedBirthday = computed(() => {
  if (!props.user.birthday) return 'не указано'
  const d = new Date(props.user.birthday)
  return `${d.toLocaleDateString()}`
})

const changeUserFormRef = ref<InstanceType<typeof ChangeUserForm> | null>(null)

const openEdit = () => {
  if (changeUserFormRef.value) {
    changeUserFormRef.value.openEdit()
  }
}

const onboarding = ref(0)

const addToFriends = async () => {
  try {
    const payload = {
      friend_user_id: props.user.id,
    }
    await profileStore.addFriend(payload)
  } catch (error) {
    console.error('Ошибка добавления в друзья:', error)
  }
}
</script>
<template>
  <div class="flex justify-center">
    <div
      class="grid grid-cols-[auto_1fr] gap-[22px] content-center bg-[#7500DB] rounded-[25px] w-[1118px] h-[447px] p-[28px]"
    >
      <div class="bg-white h-[395px] w-[395px] rounded-[20px] grid place-content-center">
        <img
          :src="`data:image/jpg;base64,${props.user.image}`"
          alt="avatar"
          class="rounded-[15px] size-[350px]"
        />
      </div>
      <div class="grid gap-[16px]">
        <div class="relative bg-white rounded-[20px] h-[207px] py-[24px] px-[28px]">
          <h1 class="text-[36px]">{{ props.user.name }}</h1>
          <button
            v-if="realLogin === props.user.login"
            class="absolute top-[24px] right-[28px] hover:scale-110 duration-200"
            @click="openEdit"
          >
            <v-avatar color="#65558f" size="48">
              <v-icon icon="settings"></v-icon>
            </v-avatar>
          </button>
          <button
            @click="addToFriends"
            type="button"
            v-if="user.friend_status === 'дружбы нет'"
            class="absolute bottom-[24px] right-[28px] hover:scale-110 duration-200 h-[50px] w-[200px] bg-[#65558f] text-white rounded-[15px]"
          >
            Добавить в друзья
          </button>
          <div
            v-if="user.friend_status === 'заявка отправлена'"
            class="absolute bottom-[24px] right-[28px] h-[50px] w-[200px] bg-gray-500 text-white rounded-[15px] flex place-content-center items-center"
          >
            Заявка отправлена
          </div>
          <div
            v-if="user.friend_status === 'дружба существует'"
            class="absolute bottom-[24px] right-[28px] h-[50px] w-[200px] bg-green-700 text-white rounded-[15px] flex place-content-center items-center"
          >
            В друзьях
          </div>
          <p class="text-[24px]">Телефон: {{ formattedPhone }}</p>
          <p class="text-[24px]">Город: {{ formattedCity }}</p>
          <p class="text-[24px]">День рождения: {{ formattedBirthday }}</p>
        </div>
        <div class="bg-white rounded-[20px] h-[172px] p-[15px] text-[20px]">
          <p>Друзья</p>
          <v-window v-if="friends && realLogin === props.user.login" v-model="onboarding" show-arrows="hover">
            <v-window-item v-for="friend in friends" :key="friend.friend.id">
              <div class="grid justify-items-center">
                <v-avatar color="#D9D9D9" size="55">
                  <img
                    :src="`data:image/jpg;base64,${friend.friend.image}` || '/default-avatar.jpg'"
                    alt="avatar"
                    class="size-[55px]"
                  />
                </v-avatar>
                <p class="text-[16px]">{{ friend.friend.name }}</p>
              </div>
            </v-window-item>
          </v-window>
          <p v-if="realLogin !== props.user.login" class="grid justify-items-center">
            В этой версии приложения ещё нельзя просматривать друзей других зумеров! 😢
          </p>
          <p v-if="!friends && realLogin === props.user.login" class="grid justify-items-center">
            У этого зумера ещё нет друзей!
          </p>
        </div>
      </div>
    </div>
  </div>
  <change-user-form ref="changeUserFormRef" :login="props.user.login" />
</template>
