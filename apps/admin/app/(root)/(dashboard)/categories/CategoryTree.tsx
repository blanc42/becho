"use client"

import React, { useState, useEffect } from 'react'
import { ChevronRight, ChevronDown, Edit, Trash2, Plus } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { useStore } from '@/lib/store/useStore'
import { Skeleton } from "@/components/ui/skeleton"
import { useRouter } from 'next/navigation'

interface Category {
  id: string
  name: string
  description: string
  parent_id: string | null
  variants: any | null
  level: number
  path: string
  subCategories?: Category[]
}

interface CategoryItemProps {
  category: Category
  onEdit: (category: Category) => void
  onDelete: (category: Category) => void
}

const CategoryItem: React.FC<CategoryItemProps> = ({ category, onEdit, onDelete }) => {
  const [isExpanded, setIsExpanded] = useState(false)

  const toggleExpand = () => {
    if (category.subCategories && category.subCategories.length > 0) {
      setIsExpanded(!isExpanded)
    }
  }

  return (
    <div className="mb-1">
      <div 
        className="flex items-center justify-between hover:dark:bg-gray-800 p-1 rounded"
        style={{ paddingLeft: `${category.level * 20}px` }}
      >
        <div className="flex items-center cursor-pointer" onClick={toggleExpand}>
          {category.subCategories && category.subCategories.length > 0 && (
            isExpanded ? <ChevronDown className="w-4 h-4 mr-1" /> : <ChevronRight className="w-4 h-4 mr-1" />
          )}
          <span>{category.name}</span>
        </div>
        <div className="flex space-x-2">
          <Button variant="ghost" size="sm" onClick={() => onEdit(category)}>
            <Edit className="w-4 h-4" />
          </Button>
          <Button variant="ghost" size="sm" onClick={() => onDelete(category)}>
            <Trash2 className="w-4 h-4" />
          </Button>
        </div>
      </div>
      {isExpanded && category.subCategories && category.subCategories.length > 0 && (
        <div className="ml-4">
          {category.subCategories.map((subCategory) => (
            <CategoryItem 
              key={subCategory.id} 
              category={subCategory} 
              onEdit={onEdit}
              onDelete={onDelete}
            />
          ))}
        </div>
      )}
    </div>
  )
}

const CategoryTree: React.FC = () => {
  const [categories, setCategories] = useState<Category[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const { selectedStore } = useStore()
  const router = useRouter()

  useEffect(() => {
    const fetchCategories = async () => {
      if (selectedStore) {
        setIsLoading(true)
        try {
          const response = await fetch(`/api/v1/stores/${selectedStore.id}/categories`)
          if (!response.ok) {
            throw new Error('Failed to fetch categories')
          }
          const data: Category[] = await response.json()
          
          // Build the category tree
          const categoryMap = new Map<string, Category>()
          const rootCategories: Category[] = []

          data.forEach(category => {
            category.subCategories = []
            categoryMap.set(category.id, category)
          })

          data.forEach(category => {
            if (category.parent_id) {
              const parentCategory = categoryMap.get(category.parent_id)
              if (parentCategory) {
                parentCategory.subCategories?.push(category)
              }
            } else {
              rootCategories.push(category)
            }
          })

          setCategories(rootCategories)
        } catch (error) {
          console.error('Error fetching categories:', error)
        } finally {
          setIsLoading(false)
        }
      }
    }

    fetchCategories()
  }, [selectedStore])

  const handleEdit = (category: Category) => {
    // Implement edit logic here
    console.log('Edit category:', category)
  }

  const handleDelete = (category: Category) => {
    // Implement delete logic here
    console.log('Delete category:', category)
  }

  const handleCreateCategory = () => {
    // Implement create category logic here
    console.log('Create new category')
  }

  if (isLoading) {
    return (
      <div className="p-4">
        <h2 className="text-xl font-bold mb-4">Categories</h2>
        {[...Array(5)].map((_, index) => (
          <div key={index} className="mb-2">
            <Skeleton className="h-8 w-full" />
          </div>
        ))}
      </div>
    )
  }

  if (categories.length === 0) {
    return (
      <div className="p-4">
        <h2 className="text-xl font-bold mb-4">Categories</h2>
        <p className="mb-4">No categories found.</p>
        <Button onClick={()=> router.push('/categories/add')}>
          <Plus className="w-4 h-4 mr-2" />
          Create Category
        </Button>
      </div>
    )
  }

  return (
    <div className="p-4">
      <h2 className="text-xl font-bold mb-4">Categories</h2>
      {categories.map((category) => (
        <CategoryItem 
          key={category.id} 
          category={category} 
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      ))}
    </div>
  )
}

export default CategoryTree