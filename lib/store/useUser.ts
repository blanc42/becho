import { create } from 'zustand'

interface User {
  id: string
  username: string
  email: string
  role: string
}

interface UserState {
  user: User | null
  setUser: (user: User) => void
  fetchUser: () => Promise<void>
}

export const useUser = create<UserState>((set) => ({
  user: null,
  setUser: (user) => set({ user }),
  fetchUser: async () => {
    try {
      const response = await fetch('/api/v1/user')
      if (!response.ok) {
        console.log(response)
        throw new Error('Failed to fetch user')
      }
      const user = await response.json()
      set({ user })
    } catch (error) {
      console.error('Error fetching user:', error)
      // Handle error (e.g., show error message to user)
    }
  },
}))