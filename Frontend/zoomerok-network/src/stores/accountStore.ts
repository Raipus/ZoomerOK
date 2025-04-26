import { defineStore } from 'pinia'
import Cookies from 'js-cookie'
import { jwtDecode } from 'jwt-decode'
import accountApi from '@/services/api/account'

interface DecodedToken {
  login: string
  // Другие поля
}

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null as any | null,
    isLoading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchUser() {
      this.isLoading = true
      try {
        const token = Cookies.get('access-token')

        if (!token) {
          throw new Error('Токен отсутствует')
        }

        const decodedToken = jwtDecode<DecodedToken>(token)
        const userLogin = decodedToken.login

        const { data } = await accountApi.getUser(userLogin)
        this.user = data
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.isLoading = false
      }
    },
  },
})
