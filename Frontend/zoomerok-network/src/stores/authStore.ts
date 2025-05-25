import { defineStore } from 'pinia'
import type { UserProfile } from './profileStore'
import accountApi from '@/services/api/account'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    login: null as string | null,
    userData: {} as UserProfile,
  }),
  actions: {
    setLogin(login: string) {
      this.login = login
    },
    async fetchUserData() {
      if (this.login) {
        try {
          const { data } = await accountApi.getUser(this.login)
          this.userData = data
        } catch (error) {
          console.log('Произошла ошибка: ' + error)
        }
      }
    },
  },
})
