import { create } from 'zustand';
import { Category } from '@/lib/types';
import { useStore } from './useStore';

interface CategoryWithSubCategories extends Category {
  subCategories?: CategoryWithSubCategories[];
}

interface CategoriesState {
  categories: CategoryWithSubCategories[];
  isLoading: boolean;
  error: string | null;
  fetchCategories: () => Promise<void>;
}

const buildCategoryTree = (categories: Category[]): CategoryWithSubCategories[] => {
  console.log("Building category tree from:", categories);
  
  const categoryMap: Record<string, CategoryWithSubCategories> = {};
  const rootCategories: CategoryWithSubCategories[] = [];

  categories.forEach(category => {
    categoryMap[category.id] = { ...category, subCategories: [] };
  });

  categories.forEach(category => {
    if (category.parent_id) {
      const parent = categoryMap[category.parent_id];
      if (parent) {
        parent.subCategories!.push(categoryMap[category.id]);
      } else {
        console.warn("Parent category not found for:", category);
      }
    } else {
      rootCategories.push(categoryMap[category.id]);
    }
  });

  console.log("Built category tree:", rootCategories);
  return rootCategories;
};

export const useCategoriesStore = create<CategoriesState>((set) => ({
  categories: [],
  isLoading: false,
  error: null,
  fetchCategories: async () => {
    const storeId = useStore.getState().selectedStore?.id;
    set({ isLoading: true, error: null });
    try {
      console.log("Fetching categories for store:", storeId);
      const response = await fetch(`/api/v1/stores/${storeId}/categories`);
      if (!response.ok) {
        throw new Error('Failed to fetch categories');
      }
      const data: Category[] = await response.json();
      console.log("Fetched categories:", data);
      const categoryTree = buildCategoryTree(data);
      console.log("Setting categories state:", categoryTree);
      set({ categories: categoryTree, isLoading: false });
    } catch (error) {
      console.error("Error fetching categories:", error);
      set({ error: (error as Error).message, isLoading: false, categories: [] });
    }
  },
}));