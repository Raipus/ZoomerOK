import { api } from '../axios-config'

export default {
  async deleteFriend(friendsData: any) {
    return api.delete(`/account/delete_friend`, friendsData)
  },

  async acceptFriend(friendsData: any) {
    return api.put(`/account/accept_friend`, friendsData)
  },

  async addFriend(friendsData: any) {
    return api.post(`/account/add_friend`, friendsData)
  },
}
