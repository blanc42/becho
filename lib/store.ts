import { create } from 'zustand'

interface Store {
  id: string
  name: string
}

interface StoreState {
  stores: Store[]
  selectedStore: Store | null
  setSelectedStore: (store: Store) => void
  addStore: (store: Store) => void
  fetchStores: () => Promise<void>
}

export const useStore = create<StoreState>((set) => ({
  stores: [],
  selectedStore: null,
  setSelectedStore: (store) => set({ selectedStore: store }),
  addStore: (store) => set((state) => ({ stores: [...state.stores, store] })),
  fetchStores: async () => {
    try {
      const response = await fetch('/api/v1/store')
      if (!response.ok) {
        throw new Error('Failed to fetch stores')
      }
      const stores = await response.json()
      set({ stores })
    } catch (error) {
      console.error('Error fetching stores:', error)
      // Handle error (e.g., show error message to user)
    }
  },
}))