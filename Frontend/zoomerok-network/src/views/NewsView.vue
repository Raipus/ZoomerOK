<script lang="ts" setup>
import { onMounted } from 'vue'
import PostComposer from '@/components/PostComposer.vue'
import PostItem from '@/components/PostItem.vue'
import { storeToRefs } from 'pinia'
import { usePostsStore } from '@/stores/postStore'

const postStore = usePostsStore()

onMounted(async () => {
  await postStore.fetchPosts(1)
})

const { posts, loading } = storeToRefs(postStore)
</script>

<template>
  <post-composer />
  <div v-if="posts.length != 0">
    <post-item v-for="post in posts" :key="post.id" :post="post" />
    <div class="flex justify-center" v-if="loading">
      <v-progress-circular indeterminate class="my-5" />
    </div>
  </div>
  <p v-if="posts.length === 0" class="flex justify-center m-7 text-2xl font-bold">
    Посты не найдены!
  </p>
</template>
