import { create } from 'zustand'
import { useStoreData } from './useStoreData'
import { User } from '../schema'


interface UserState {
  user: User | null
  setUser: (user: User) => void
  fetchUser: (url: string) => Promise<User>
}

export const useUser = create<UserState>((set, get) => ({
  user: null,
  setUser: (user) => {
    set({ user })
    const { setStores } = useStoreData.getState()
    setStores(user.stores)
  },
  fetchUser: async (url: string) => {
      const response = await fetch(url, {
        credentials: 'include',
      });
      if (response.status === 401) {
        throw new Error('Unauthorized');
      }
      if (!response.ok) {
        throw new Error('An error occurred while fetching the data.');
      }
      const data: { data: User } = await response.json();
      console.log("data", data)
      set({ user: data.data });
      const { setStores } = useStoreData.getState();
      setStores(data.data.stores);
      return data.data;
  },
}))