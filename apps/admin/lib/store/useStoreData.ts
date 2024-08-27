import { create } from 'zustand'

interface Store {
  id: string
  name: string
}

interface StoreState {
  stores: Store[]
  selectedStore: Store | null
  setStores: (stores: Store[]) => void
  setSelectedStore: (store: Store) => void
  createStore: (store: Store) => Promise<void>
  addStore: (store: Store) => void
  fetchStores: () => Promise<void>
}

export const useStoreData = create<StoreState>((set) => ({
  stores: [],
  setStores: (stores) => {
    set({ stores })
    if (stores.length > 0) {
      console.log("we are running inside the setStores logic")
      console.log('stores', stores)
      const storedStoreId = localStorage.getItem('selectedStoreId')
      const selectedStore = storedStoreId 
        ? stores.find(store => store.id === storedStoreId) 
        : stores[0]
      if (selectedStore) {
        console.log("I am running inside the if selected store logic")
        set({ selectedStore })
        localStorage.setItem('selectedStoreId', selectedStore.id)
      }
    }
  },
  selectedStore: null,
  setSelectedStore: (store) => set({ selectedStore: store }),
  createStore: async (store) => {
    const response = await fetch('/api/v1/stores', {
      method: 'POST',
      body: JSON.stringify(store),
    })
    if (!response.ok) {
      throw new Error('Failed to create store')
    }
    const newStore = await response.json()
    set((state) => ({ stores: [...state.stores, newStore] }))
    set({ selectedStore: newStore })
  },
  addStore: (store) => set((state) => ({ stores: [...state.stores, store] })),
  fetchStores: async () => {
    try {
      const response = await fetch('/api/v1/stores')
      if (!response.ok) {
        throw new Error('Failed to fetch stores')
      }
      const stores = await response.json()
      set((state) => ({ ...state, stores }))
      const storedStoreId = localStorage.getItem('selectedStoreId')
      
      if (storedStoreId) {
        const store = stores.find((s: Store) => s.id === storedStoreId)
        if (store) {
          set((state) => ({ ...state, selectedStore: store }))
        }
      } else if (stores.length > 0) {
        set((state) => ({ ...state, selectedStore: stores[0] }))
        localStorage.setItem('selectedStoreId', stores[0].id)
      }
    } catch (error) {
      console.error('Error fetching stores:', error)
      // Handle error (e.g., show error message to user)
    }
  },
}))