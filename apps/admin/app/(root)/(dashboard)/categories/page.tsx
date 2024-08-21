"use client"
import { categoriesData } from "@/data/variants";
import CategoryTree from "@/components/CategoryTree";
import { Category } from "@/lib/types";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";


export default function CategoriesPage() {
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null);
  const router = useRouter();


  return (
    <div className="w-full">
      <div className="flex justify-between items-center gap-4 mb-8 border-b pb-4">
        <h1 className="text-2xl font-bold">Categories</h1>
        <Button onClick={() => router.push('/categories/add')}>Add Category</Button>
      </div>
        <CategoryTree categories={categoriesData} onChange={(value) => setSelectedCategory(value)} value={selectedCategory} />
    </div>
  ) 
}