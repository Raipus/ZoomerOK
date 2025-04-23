import { api } from '../axios-config'

export default {
  async getUser(login: string) {
    return api.get(`/account/user/${login}`)
  },

  async deleteUser(login: string) {
    return api.delete(`/account/user/${login}`)
  },

  async сhangeUser(login: string) {
    return api.put(`/account/user/${login}`)
  },
}
