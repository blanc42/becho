import { create } from 'zustand'
import { useStoreData } from './useStoreData'

interface Variant {
  id: string
  name: string
  options: string[]
}

interface Product {
  id: string
  name: string
  description: string
  price: number
  categoryId: string
  storeId: string
  variants: Variant[]
}

interface ProductsState {
  products: Product[]
  isLoading: boolean
  error: string | null
  fetchProducts: () => Promise<void>
  addProduct: (product: Product) => void
  updateProduct: (product: Product) => void
  deleteProduct: (productId: string) => void
}

export const useProducts = create<ProductsState>((set, get) => ({
  products: [],
  isLoading: false,
  error: null,

  fetchProducts: async () => {
    const selectedStore = useStoreData.getState().selectedStore

    if (!selectedStore) {
      set({ error: "No store selected" })
      return
    }

    set({ isLoading: true, error: null })

    try {
      const response = await fetch(`/api/v1/stores/${selectedStore.id}/products`)
      if (!response.ok) {
        throw new Error('Failed to fetch products')
      }
      const products: Product[] = await response.json()
      set({ products, isLoading: false })
    } catch (error) {
      set({ error: (error as Error).message, isLoading: false })
    }
  },

  addProduct: (product: Product) => {
    set((state) => ({ products: [...state.products, product] }))
  },

  updateProduct: (updatedProduct: Product) => {
    set((state) => ({
      products: state.products.map((product) =>
        product.id === updatedProduct.id ? updatedProduct : product
      ),
    }))
  },

  deleteProduct: (productId: string) => {
    set((state) => ({
      products: state.products.filter((product) => product.id !== productId),
    }))
  },
}))

// This effect will run once when the module is imported
// It sets up a listener for changes in the selected store
// useStore.subscribe(
//   (state) => state.selectedStore,
//   (selectedStore) => {
//     if (selectedStore) {
//       useProducts.getState().fetchProducts()
//     } else {
//       useProducts.setState({ products: [], error: null })
//     }
//   }
// )