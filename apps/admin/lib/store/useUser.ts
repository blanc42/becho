import { create } from 'zustand'
import { useStoreData } from './useStoreData'

interface User {
  id: string
  username: string
  email: string
  role: string
  stores: Store[]
}

interface Store {
  id: string
  name: string
}

interface UserState {
  user: User | null
  setUser: (user: User) => void
  fetchUser: () => Promise<User>
}

export const useUser = create<UserState>((set, get) => ({
  user: null,
  setUser: (user) => {
    set({ user })
    const { setStores } = useStoreData.getState()
    setStores(user.stores)
  },
  fetchUser: async () => {
    try {
      const response = await fetch('/api/v1/user', {
        credentials: 'include',
      })
      if (!response.ok) {
        console.log(response)
        throw new Error('Failed to fetch user')
      }
      const res = await response.json()
      set({ user : res.data })
      const { setStores } = useStoreData.getState()
      setStores(res.data.stores)

      return res.data
    } catch (error) {
      console.error('Error fetching user:', error)
      // Handle error (e.g., show error message to user)
    } finally {
      console.log("finished fetching user")
    }
  },
}))