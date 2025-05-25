import { api } from '@/services/axios-config'

export interface Post {
  id: number
  user: {
    id: number
    login: string
    name: string
    image?: string
  }
  body: {
    id: number
    text: string
    image?: string
    time: string
    number_of_comments: number
    number_of_likes: number
    isLiked: boolean
  }
}

export interface Comment {
  user: {
    id: number
    login: string
    name: string
    image?: string
  }
  body: {
    id: number
    text: string
    time: string
    post_id: number
  }
}

export const PostsApi = {
  async getPosts(page: number): Promise<Post[]> {
    const response = await api.get('/blog/posts', { params: { page } })
    return response.data.posts
  },

  async getUserPosts(userId: number, page: number): Promise<Post[]> {
    const response = await api.get(`/blog/user/${userId}/posts`, { params: { page } })
    return response.data.posts
  },

  async getPost(postId: number): Promise<Post> {
    const response = await api.get(`/blog/post/${postId}`)
    return response.data.post
  },

  async createPost(text: string, image?: string): Promise<number> {
    const formData = new FormData()
    formData.append('text', text)
    if (image) formData.append('image', image)

    const response = await api.post('/blog/create_post', formData)
    return response.data.id
  },

  async deletePost(postId: number): Promise<void> {
    await api.delete(`/blog/post/${postId}`)
  },

  async likePost(postId: number): Promise<void> {
    await api.post(`/blog/post/${postId}/like`)
  },

  async getComments(postId: number, page: number): Promise<Comment[]> {
    const response = await api.get(`/blog/post/${postId}/comments`, { params: { page } })
    return response.data.comments
  },

  async createComment(postId: number, text: string): Promise<void> {
    await api.post(`/blog/post/${postId}/create_comment`, { text })
  },

  async deleteComment(postId: number, commentId: number): Promise<void> {
    await api.delete(`/blog/post/${postId}/comments/${commentId}`)
  },
}
