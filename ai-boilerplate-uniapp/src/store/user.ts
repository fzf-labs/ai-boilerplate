import type { GetUserInfoReply } from '@/api/v1/user/types'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import {
  getUserInfo as apiGetUserInfo,
} from '@/api/v1/user/user'

interface IUserInfoRes {
  userId: number
  username: string
  nickname: string
  avatar: string
}

function mapUserInfo(reply: GetUserInfoReply): IUserInfoRes {
  const info = reply?.info
  const userId = Number(info?.id ?? -1)
  return {
    userId: Number.isFinite(userId) ? userId : -1,
    username: info?.phone ?? '',
    nickname: info?.nickname ?? '',
    avatar: info?.avatar ?? '',
  }
}

// 初始化状态
const userInfoState: IUserInfoRes = {
  userId: -1,
  username: '',
  nickname: '',
  avatar: '/static/images/default-avatar.png',
}

export const useUserStore = defineStore(
  'user',
  () => {
    // 定义用户信息
    const userInfo = ref<IUserInfoRes>({ ...userInfoState })
    // 设置用户信息
    const setUserInfo = (val: IUserInfoRes) => {
      console.log('设置用户信息', val)
      // 若头像为空 则使用默认头像
      if (!val.avatar) {
        val.avatar = userInfoState.avatar
      }
      userInfo.value = val
    }
    const setUserAvatar = (avatar: string) => {
      userInfo.value.avatar = avatar
      console.log('设置用户头像', avatar)
      console.log('userInfo', userInfo.value)
    }
    // 删除用户信息
    const clearUserInfo = () => {
      userInfo.value = { ...userInfoState }
      uni.removeStorageSync('user')
    }

    /**
     * 获取用户信息
     */
    const fetchUserInfo = async () => {
      const res = await apiGetUserInfo({})
      const user = mapUserInfo(res)
      setUserInfo(user)
      return user
    }

    return {
      userInfo,
      clearUserInfo,
      fetchUserInfo,
      setUserInfo,
      setUserAvatar,
    }
  },
  {
    persist: true,
  },
)
