import { defineStore } from 'pinia'
import { ref } from 'vue'
import accountApi, { type ChangeUserProfileDto, type FriendId } from '../services/api/account'

export interface UserProfile {
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

export interface Friend {
  friend: {
    id: number
    name: string
    login: string
    image?: string
  }
}

export const useProfileStore = defineStore('profile', () => {
  const user = ref<UserProfile>({} as UserProfile)
  const friends = ref<Friend[]>([])
  const unacceptedFriends = ref<Friend[]>([])

  async function fetchUser(login: string) {
    try {
      const { data } = await accountApi.getUser(login)
      user.value = data
    } catch (error) {
      console.log('Произошла ошибка')
      throw(error)
    }
  }

  async function changeUser(login: string, userData: ChangeUserProfileDto) {
    const stashedUser = { ...user.value }
    user.value = { ...user.value, ...userData }
    try {
      await accountApi.сhangeUser(login, userData)
    } catch (error) {
      user.value = stashedUser
      console.log('Произошла ошибка')
      throw(error)
    }
  }

  async function fetchFriends() {
    try {
      const { data } = await accountApi.getFriends()
      friends.value = data?.users || []
    } catch (error) {
      console.log('Произошла ошибка')
      throw error
    }
  }

  async function fetchUnacceptedFriends() {
    try {
      const { data } = await accountApi.getUnacceptedFriends()
      unacceptedFriends.value = data?.unaccepted_friends || []
    } catch (error) {
      console.log('Произошла ошибка')
      throw error
    }
  }

  async function addFriend(friendId: FriendId) {
    try {
      await accountApi.addFriend(friendId)
      user.value.friend_status = 'заявка отправлена'
    } catch (error) {
      user.value.friend_status = 'дружбы нет'
      console.log('Произошла ошибка')
      throw error
    }
  }

  async function acceptFriend(friendId: FriendId) {
    try {
      await accountApi.acceptFriend(friendId)
      const foundFriend = unacceptedFriends.value.find(
        (f) => f.friend.id === friendId.friend_user_id,
      )

      if (foundFriend) {
        friends.value = [...(friends.value || []), foundFriend]

        unacceptedFriends.value = unacceptedFriends.value.filter(
          (f) => f.friend.id !== friendId.friend_user_id,
        )
      }
    } catch (error) {
      console.log('Произошла ошибка')
      throw error
    }
  }

  async function deleteFriend(friendId: FriendId) {
    try {
      await accountApi.deleteFriend(friendId)
      friends.value = friends.value.filter((f) => f.friend.id !== friendId.friend_user_id)
    } catch (error) {
      console.log('Произошла ошибка')
      throw error
    }
  }

  return {
    user,
    friends,
    unacceptedFriends,
    fetchUser,
    fetchFriends,
    changeUser,
    fetchUnacceptedFriends,
    addFriend,
    acceptFriend,
    deleteFriend,
  }
})
