import { api } from '../axios-config'

export interface UserProfileDto {
  id: number
  login: string
  name: string
  email: string
  birthday?: string
  phone?: string
  city?: string
  image?: string
  friend_status: string
}

export interface ChangeUserProfileDto {
  name: string
  birthday: string | undefined
  phone: string | undefined
  city: string | undefined
  image: string | undefined
}

export interface FriendId {
  friend_user_id: number
}

export default {
  async getUser(login: string) {
    return api.get<UserProfileDto>(`/account/user/${login}`)
  },

  async deleteUser(login: string) {
    return api.delete(`/account/user/${login}`)
  },

  async сhangeUser(login: string, userData: ChangeUserProfileDto) {
    return api.put(`/account/user/${login}`, userData)
  },

  async deleteFriend(friendId: FriendId) {
    return api.delete(`/account/delete_friend`, { data: friendId })
  },

  async acceptFriend(friendId: FriendId) {
    return api.put(`/account/accept_friend`, friendId)
  },

  async addFriend(friendId: FriendId) {
    return api.post(`/account/add_friend`, friendId)
  },

  async getFriends() {
    return api.get(`/account/get_friends`)
  },

  async getUnacceptedFriends() {
    return api.get(`/account/get_unaccepted_friends`)
  },
}
