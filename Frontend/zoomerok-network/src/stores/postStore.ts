import { defineStore } from 'pinia'
import { PostsApi, type Post, type Comment } from '@/services/api/posts'
import { useAuthStore } from './authStore'

interface PostsState {
  userPosts: Post[],
  posts: Post[]
  currentPost: Post | null
  comments: Comment[]
  loading: boolean
  error: string | null
}

export const usePostsStore = defineStore('posts', {
  state: (): PostsState => ({
    userPosts: [],
    posts: [],
    currentPost: null,
    comments: [],
    loading: false,
    error: null,
  }),

  actions: {
    async fetchPosts(page: number) {
      try {
        this.loading = true
        const data = await PostsApi.getPosts(page)
        this.posts = Array.isArray(data) ? data : []
        this.error = null
      } catch (error) {
        this.error = 'Ошибка загрузки постов'
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchUserPosts(userId: number, page: number) {
      try {
        this.loading = true
        const data = await PostsApi.getUserPosts(userId, page)
        this.userPosts = Array.isArray(data) ? data : []
        this.error = null
      } catch (error) {
        this.error = 'Ошибка загрузки постов пользователя'
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchPost(postId: number) {
      try {
        this.loading = true
        this.currentPost = await PostsApi.getPost(postId)
        this.error = null
      } catch (error) {
        this.error = 'Ошибка загрузки поста'
        throw error
      } finally {
        this.loading = false
      }
    },

    async createPost(text: string, image?: string) {
      const authStore = useAuthStore()
      let id: number | null = null

      try {
        id = Date.now()
        const tempPost = {
          id: id,
          user: {
            id: authStore.userData.id,
            login: authStore.userData.login,
            name: authStore.userData.name,
            image: authStore.userData.image,
          },
          body: {
            id: id,
            text,
            image: image ? image : undefined,
            time: new Date().toISOString(),
            number_of_comments: 0,
            number_of_likes: 0,
            isLiked: false,
          },
        }

        this.posts.unshift(tempPost)
        this.userPosts.unshift(tempPost)

        const realPostId = await PostsApi.createPost(text, image)
        const realPost = await PostsApi.getPost(realPostId)

        const index = this.posts.findIndex((p) => p.id === tempPost.id)
        this.posts.splice(index, 1, realPost)

        const index1 = this.userPosts.findIndex((p) => p.id === tempPost.id)
        this.userPosts.splice(index1, 1, realPost)
      } catch (error) {
        if (id !== null && Array.isArray(this.posts)) {
          this.posts = this.posts.filter((p) => p.id !== id)
          this.userPosts = this.userPosts.filter((p) => p.id !== id)
        }
        throw error
      }
    },

    async deletePost(postId: number) {
      try {
        this.loading = true
        await PostsApi.deletePost(postId)
        this.posts = this.posts.filter((post) => post.body.id !== postId)
        this.error = null
      } catch (error) {
        this.error = 'Ошибка удаления поста'
        throw error
      } finally {
        this.loading = false
      }
    },

    async toggleLike(postId: number) {
      try {
        await PostsApi.likePost(postId)
        const post = this.posts.find((p) => p.body.id === postId)
        if (post) {
          if (post.body.isLiked === true) {
            post.body.number_of_likes -= 1
            post.body.isLiked = false
          } else {
            post.body.number_of_likes += 1
            post.body.isLiked = true
          }
        }
      } catch (error) {
        this.error = 'Ошибка обработки лайка'
        throw error
      }
    },

    async fetchComments(postId: number, page: number) {
      try {
        this.comments = await PostsApi.getComments(postId, page)
        this.error = null
      } catch (error) {
        this.error = 'Ошибка загрузки комментариев'
        throw error
      }
    },

    async createComment(postId: number, text: string) {
      try {
        await PostsApi.createComment(postId, text)
        await this.fetchComments(postId, 1)
      } catch (error) {
        this.error = 'Ошибка создания комментария'
        throw error
      }
    },

    async deleteComment(postId: number, commentId: number) {
      try {
        await PostsApi.deleteComment(postId, commentId)
        this.comments = this.comments.filter((c) => c.body.id !== commentId)
      } catch (error) {
        this.error = 'Ошибка удаления комментария'
        throw error
      }
    },
  },
})
