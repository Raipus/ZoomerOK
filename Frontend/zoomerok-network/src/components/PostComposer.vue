<template>
  <div class="flex justify-center">
    <div
      class="w-[610px] min-h-[83px] mt-[24px] bg-[#7500DB] rounded-[40px] grid items-center place-content-center p-[14px]"
    >
      <div
        class="place-content-center items-center grid grid-cols-[auto_426px_auto] gap-x-[10px] grid-rows-[55px_auto] px-[16px] min-h-[55px] bg-white rounded-[40px] top-[14px] left-[14px]"
      >
        <input ref="fileInput" type="file" accept="image/*" class="d-none" @change="onFileChange" />
        <v-btn icon @click="() => fileInput.click()">
          <v-icon>attach_file</v-icon>
        </v-btn>
        <v-text-field
          v-model="text"
          placeholder="Что у Вас нового?"
          variant="underlined"
          clearable
          class="h-[55px] mb-3"
        />
        <v-btn icon @click="submit">
          <v-icon>send</v-icon>
        </v-btn>
      </div>
      <v-img v-if="preview" :src="preview" class="mt-4 max-h-96 rounded-[40px]" contain />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { usePostsStore } from '@/stores/postStore'
import { ref, type Ref } from 'vue'

const store = usePostsStore()
const text = ref('')
const fileInput = ref<HTMLInputElement>() as Ref<HTMLInputElement>
const preview = ref<string | null>(null)
let image: string | undefined = undefined

function onFileChange(e: Event) {
  const f = (e.target as HTMLInputElement).files?.[0]
  if (!f) return
  if (f.size > 512 * 1024) {
    alert('Максимальный размер изображения — 512 KB')
    return
  }
  if (f) {
    const reader = new FileReader()
    reader.onload = (e) => {
      if (e.target?.result) {
        const base64Data = e.target.result.toString().split(',')[1]
        image = base64Data
      }
    }
    reader.readAsDataURL(f)
  }
  preview.value = URL.createObjectURL(f)
}

async function submit() {
  if (!text.value.trim() && !image) return
  await store.createPost(text.value, image)
  text.value = ''
  preview.value = null
}
</script>
