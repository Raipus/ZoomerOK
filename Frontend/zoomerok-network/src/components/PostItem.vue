<template>
  <div class="flex justify-center">
    <div
      class="grid gap-[15px] min-h-[171px] h-max w-[610px] bg-[#7500DB] mt-[20px] rounded-[40px] p-[25px]"
      :class="post.body.image ? 'grid-rows-[60px_auto_305px_35px]' : 'grid-rows-[60px_auto_35px]'"
    >
      <div class="h-[60px] grid-cols-[60px_auto] grid gap-2 items-center">
        <img
          :src="
            hasProfile
              ? `data:image/jpg;base64,${userStore.user.image}`
              : `data:image/jpg;base64,${post.user.image}`
          "
          alt="avatar"
          class="h-[60px] w-[60px] rounded-full col-1"
        />
        <div class="col-2 text-white">
          <p class="text-[24px] font-bold">
            {{ hasProfile ? userStore.user.name : post.user.name }}
          </p>
          <p class="text-[18px]">{{ formattedTime }}</p>
        </div>
      </div>
      <div
        class="min-h-[35px] h-max min-w-[0] flex bg-white rounded-[40px] text-[20px] items-center"
      >
        <p class="break-words max-w-[552px] px-[12px] py-2">{{ post.body.text }}</p>
      </div>
      <div
        v-if="post.body.image"
        class="h-[305px] rounded-[14px] bg-[#838383] grid place-content-center overflow-hidden relative"
      >
        <img
          :src="`data:image/jpg;base64,${post.body.image}`"
          alt="post photo"
          class="h-[305px] object-contain z-20 shadow-[4px_0_15px_rgba(0,0,0,0.3),-4px_0_15px_rgba(0,0,0,0.3)]"
        />
        <img
          :src="`data:image/jpg;base64,${post.body.image}`"
          alt="post photo"
          class="h-full w-full object-cover object-center absolute z-10 blur-md"
        />
      </div>
      <div class="h-[35px] bg-white rounded-[40px] px-[12px] grid grid-cols-2 place-items-center">
        <div class="grid grid-cols-2 items-center w-max">
          <p>{{ post.body.number_of_likes }}</p>
          <button @click="() => toggleLike()" class="hover:scale-110 duration-200">
            <v-icon :color="post.body.isLiked ? 'blue' : 'black'"> thumb_up_alt </v-icon>
          </button>
        </div>
        <div class="grid grid-cols-2 items-center w-max">
          <p>{{ post.body.number_of_comments }}</p>
          <button @click="toggleComments" class="hover:scale-110 duration-200">
            <v-icon :color="showComments ? 'blue' : 'black'"> comment </v-icon>
          </button>
        </div>
      </div>
      <div
        class="overflow-hidden"
        :class="{
          'max-h-0': !showComments,
          'max-h-[500px] transition-all duration-300 ease-out': showComments,
        }"
      >
        <div v-if="showComments" class="mt-4 space-y-4 bg-white rounded-[40px] p-[12px]">
          <div class="h-[300px] overflow-y-auto space-y-3 pr-2">
            <div
              v-for="comment in postStore.comments"
              :key="comment.body.id"
              class="bg-[#7500DB] rounded-[20px] p-3 relative pr-8"
            >
              <div class="flex items-center gap-2 mb-2">
                <img
                  :src="
                    comment.user.image
                      ? `data:image/jpg;base64,${comment.user.image}`
                      : defaultAvatar
                  "
                  class="w-8 h-8 rounded-full"
                />
                <span class="font-bold text-white">{{ comment.user.name }}</span>
                <span class="text-sm text-gray-400">{{
                  formatCommentTime(comment.body.time)
                }}</span>
              </div>
              <p class="text-gray-200">{{ comment.body.text }}</p>
              <button
                v-if="canDeleteComment(comment)"
                @click="deleteComment(comment)"
                class="absolute top-2 right-2 rounded-full size-7 bg-white hover:scale-110 duration-200"
              >
                <v-icon small color="red">delete</v-icon>
              </button>
            </div>
          </div>

          <div class="flex gap-2">
            <input
              v-model="newComment"
              placeholder="Введите комментарий..."
              class="flex-1 rounded-[20px] px-4 py-2 focus:outline-none"
              @keyup.enter="submitComment"
            />
            <button
              @click="submitComment"
              class="bg-white rounded-[20px] px-4 py-2 hover:bg-gray-100 transition-colors"
            >
              <v-icon>send</v-icon>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, defineProps, ref, watch } from 'vue'
import { type Post, type Comment } from '@/services/api/posts'
import { usePostsStore } from '@/stores/postStore'
import { useProfileStore } from '@/stores/profileStore'
import { useAuthStore } from '@/stores/authStore'
import { useRoute } from 'vue-router'

const authStore = useAuthStore()
const props = defineProps<{ post: Post }>()
const postStore = usePostsStore()
const userStore = useProfileStore()
const defaultAvatar = '/default-avatar.jpg'
const showComments = ref(false)
const newComment = ref('')

const route = useRoute()
const hasProfile = computed(() => {
  return route.path.toLowerCase().includes('profile')
})

const formattedTime = computed(() => {
  const d = new Date(props.post.body.time)
  return d.toLocaleString()
})

async function toggleLike() {
  await postStore.toggleLike(props.post.body.id)
}

watch(showComments, async (val) => {
  if (val) {
    await postStore.fetchComments(props.post.body.id, 1)
  }
})

const toggleComments = () => {
  showComments.value = !showComments.value
  if (showComments.value && postStore.comments.length === 0) {
    postStore.fetchComments(props.post.body.id, 1)
  }
}

const formatCommentTime = (timeString: string) => {
  const options: Intl.DateTimeFormatOptions = {
    hour: '2-digit',
    minute: '2-digit',
    day: '2-digit',
    month: '2-digit',
  }
  return new Date(timeString).toLocaleString('ru-RU', options)
}

const submitComment = async () => {
  if (!newComment.value.trim()) return

  try {
    await postStore.createComment(props.post.body.id, newComment.value.trim())
    newComment.value = ''
    props.post.body.number_of_comments += 1
  } catch (error) {
    console.error('Ошибка отправки комментария:', error)
  }
}

const canDeleteComment = (comment: Comment) => {
  return authStore.userData?.id === comment.user.id
}

const deleteComment = async (comment: Comment) => {
  try {
    await postStore.deleteComment(comment.body.post_id, comment.body.id)
    props.post.body.number_of_comments -= 1
  } catch (error) {
    console.error('Ошибка удаления комментария:', error)
  }
}
</script>
