"use client"

import React, { useState } from 'react'
import { ChevronRight, ChevronDown, Edit, Trash2 } from 'lucide-react'
import { Button } from "@/components/ui/button"

interface Category {
  id: string
  name: string
  subCategories: Category[]
}

interface CategoryItemProps {
  category: Category
  level: number
  onEdit: (category: Category) => void
  onDelete: (category: Category) => void
}

const CategoryItem: React.FC<CategoryItemProps> = ({ category, level, onEdit, onDelete }) => {
  const [isExpanded, setIsExpanded] = useState(false)

  const toggleExpand = () => {
    if (category.subCategories.length > 0) {
      setIsExpanded(!isExpanded)
    }
  }

  return (
    <div className="mb-1">
      <div 
        className="flex items-center justify-between hover:dark:bg-gray-800 p-1 rounded"
        style={{ paddingLeft: `${level * 20}px` }}
      >
        <div className="flex items-center cursor-pointer" onClick={toggleExpand}>
          {category.subCategories.length > 0 && (
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
      {isExpanded && category.subCategories.length > 0 && (
        <div className="ml-4">
          {category.subCategories.map((subCategory) => (
            <CategoryItem 
              key={subCategory.id} 
              category={subCategory} 
              level={level + 1}
              onEdit={onEdit}
              onDelete={onDelete}
            />
          ))}
        </div>
      )}
    </div>
  )
}

const mockCategories: Category[] = [
  {
    id: "1",
    name: "Electronics",
    subCategories: [
      {
        id: "2",
        name: "Computers",
        subCategories: [
          {
            id: "3",
            name: "Laptops",
            subCategories: []
          },
          {
            id: "4",
            name: "Desktops",
            subCategories: []
          }
        ]
      },
      {
        id: "5",
        name: "Smartphones",
        subCategories: []
      }
    ]
  },
  {
    id: "6",
    name: "Clothing",
    subCategories: [
      {
        id: "7",
        name: "Men's",
        subCategories: []
      },
      {
        id: "8",
        name: "Women's",
        subCategories: []
      }
    ]
  }
]

const CategoryTree: React.FC = () => {
  const [categories, setCategories] = useState<Category[]>(mockCategories)

  // Commented out fetch logic
  /*
  useEffect(() => {
    const fetchCategories = async () => {
      if (selectedStore) {
        try {
          const response = await fetch(`/api/v1/stores/${selectedStore.id}/categories`)
          if (!response.ok) {
            throw new Error('Failed to fetch categories')
          }
          const data = await response.json()
          setCategories(data)
        } catch (error) {
          console.error('Error fetching categories:', error)
        }
      }
    }

    fetchCategories()
  }, [selectedStore])
  */

  const handleEdit = (category: Category) => {
    // Implement edit logic here
    console.log('Edit category:', category)
  }

  const handleDelete = (category: Category) => {
    // Implement delete logic here
    console.log('Delete category:', category)
  }

  return (
    <div className="p-4">
      <h2 className="text-xl font-bold mb-4">Categories</h2>
      {categories.map((category) => (
        <CategoryItem 
          key={category.id} 
          category={category} 
          level={0}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      ))}
    </div>
  )
}

export default CategoryTree