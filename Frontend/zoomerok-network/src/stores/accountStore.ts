import { defineStore } from 'pinia'
import Cookies from 'js-cookie'
import { jwtDecode } from 'jwt-decode'
import accountApi from '@/services/api/account'

interface DecodedToken {
  login: string
  // Другие поля
}

export const useAccountStore = defineStore('account', {
  state: () => ({
    account: null as any | null,
    isAuthenticated: false,
    isLoading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchUser() {
      this.isLoading = true
      try {
        const token = Cookies.get('access-token')

        if (!token) {
          this.isAuthenticated = false
          throw new Error('Токен отсутствует')
        }

        this.isAuthenticated = true
        const decodedToken = jwtDecode<DecodedToken>(token)
        const userLogin = decodedToken.login

        const { data } = await accountApi.getUser(userLogin)
        this.account = data
      } catch (error: any) {
        this.error = error.message
      } finally {
        this.isLoading = false
      }
    },
  },
})
