<script setup lang="ts">
import { onMounted, onUnmounted, ref, defineExpose, watchEffect } from 'vue'
import { useProfileStore } from '@/stores/profileStore'
import type { ChangeUserProfileDto } from '@/services/api/account'

const props = defineProps<{
  login: string
}>()

const store = useProfileStore()

const isFormOpen = ref(false)
const isFormVisible = ref(false)

const formData = ref<ChangeUserProfileDto>({
  name: '',
  birthday: '',
  phone: '',
  city: '',
  image: '',
})

const previewImage = ref<string | null>(null)

const handleFileUpload = (event: Event) => {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]

  if (file) {
    const reader = new FileReader()
    reader.onload = (e) => {
      if (e.target?.result) {
        const base64Data = e.target.result.toString().split(',')[1]
        formData.value.image = base64Data
        previewImage.value = e.target.result.toString()
      }
    }
    reader.readAsDataURL(file)
  }
}

watchEffect(() => {
  if (store.user) {
    formData.value = {
      name: store.user.name || '',
      birthday: store.user.birthday?.split('T')[0] || '',
      phone: store.user.phone || '',
      city: store.user.city || '',
      image: store.user.image || '',
    }
    previewImage.value = store.user.image ? `data:image/jpeg;base64,${store.user.image}` : null
  }
})

const openEdit = () => {
  isFormOpen.value = true
  setTimeout(() => {
    isFormVisible.value = true
  }, 10)
}

defineExpose({
  openEdit,
})

const closeEdit = () => {
  isFormVisible.value = false
  setTimeout(() => {
    isFormOpen.value = false
  }, 300)
  formData.value = {
    name: store.user.name || '',
    birthday: store.user.birthday?.split('T')[0] || '',
    phone: store.user.phone || '',
    city: store.user.city || '',
    image: store.user.image || '',
  }
}

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Escape') closeEdit()
}

const saveForm = async () => {
  try {
    const payload = {
      ...formData.value,
      birthday: formData.value.birthday ? `${formData.value.birthday}T00:00:00Z` : undefined,
    }

    await store.changeUser(props.login, payload)
  } catch (error) {
    console.error('Ошибка обновления профиля:', error)
  }
  closeEdit()
}

onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})
</script>
<template>
  <div
    ref="editForm"
    class="fixed top-0 left-0 bg-black/20 w-full h-full z-[100] transition-all duration-300 ease-in-out"
    :class="{
      'opacity-0 scale-95 -translate-y-2': !isFormVisible,
      'opacity-100 scale-100 translate-y-0': isFormVisible,
    }"
    v-show="isFormOpen"
  >
    <div class="grid place-items-center h-full">
      <div class="text-black text-xl rounded-3xl h-max w-[450px] bg-[#fff] shadow-2xl p-5">
        <h1 class="text-2xl place-content-center grid">Редактирование профиля</h1>
        <div class="mt-4 flex flex-col items-center">
          <div class="relative w-32 h-32 rounded-full overflow-hidden border-2 border-gray-300">
            <img
              :src="previewImage || '/default-avatar.jpg'"
              class="w-full h-full object-cover"
              alt="Аватар"
            />
            <label
              class="absolute bottom-0 w-full bg-black/50 text-white text-center cursor-pointer py-1.5 text-lg"
            >
              <input type="file" accept="image/*" class="hidden" @change="handleFileUpload" />
              Изменить
            </label>
          </div>
        </div>
        <form>
          <div class="grid gap-5 mt-5">
            <div class="grid grid-cols-2">
              <label>Имя:</label>
              <input
                v-model="formData.name"
                type="text"
                name=""
                id=""
                class="bg-[#D9D9D9] rounded-md px-2"
              />
            </div>
            <div class="grid grid-cols-2">
              <label>День рождения:</label>
              <input
                v-model="formData.birthday"
                type="date"
                name=""
                id=""
                class="bg-[#D9D9D9] rounded-md px-2"
              />
            </div>
            <div class="grid grid-cols-2">
              <label>Телефон:</label>
              <input
                v-model="formData.phone"
                type="text"
                name=""
                id=""
                class="bg-[#D9D9D9] rounded-md px-2"
              />
            </div>
            <div class="grid grid-cols-2">
              <label>Город:</label>
              <input
                v-model="formData.city"
                type="text"
                name=""
                id=""
                class="bg-[#D9D9D9] rounded-md px-2"
              />
            </div>
          </div>
          <div class="flex justify-between w-3/4 justify-self-center">
            <button
              @click="closeEdit"
              type="button"
              class="text-2xl justify-self-center mt-5 grid bg-[#D9D9D9] rounded-md px-2 hover:scale-110 hover:bg-red-500 duration-200"
            >
              Отменить
            </button>
            <button
              @click="saveForm"
              type="button"
              class="text-2xl justify-self-center mt-5 grid bg-[#D9D9D9] rounded-md px-2 hover:scale-110 hover:bg-green-500 duration-200"
            >
              Изменить
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.relative:hover .absolute {
  opacity: 1;
}
.absolute {
  opacity: 0;
  transition: opacity 0.3s ease;
}
</style>
