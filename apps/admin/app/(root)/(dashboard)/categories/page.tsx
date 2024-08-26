"use client"
import { useEffect, useState } from "react";
import { useStoreData } from "@/lib/store/useStoreData";
import { Skeleton } from "@/components/ui/skeleton";
import CategoryTree from "./CategoryTree";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { PlusSquareIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { Category } from "@/lib/types";


export default function CategoriesPage() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { selectedStore } = useStoreData();
  const router = useRouter();


  useEffect(() => {
    setIsLoading(true);
    const fetchCategories = async () => {
      if (selectedStore) {
        try {
          if (selectedStore) {
            const response = await fetch(`/api/v1/stores/${selectedStore.id}/categories`);
            if (!response.ok) {
              throw new Error('Failed to fetch categories');
            }
            const data = await response.json();
            setCategories(data);
          } else {
            router.push('/');
          }
        } catch (error) {
          console.error('Error fetching categories:', error);
        } finally {
          setIsLoading(false);
        }
      }
    };

    fetchCategories();
  }, [selectedStore]);

  return (
    <div className="container max-w-screen-xl">
      <div className="flex items-baseline justify-between my-12">
        <h1 className="text-3xl font-semibold">Categories</h1>
        <Button asChild>
          <Link href="/categories/add" className='flex items-center gap-2'>
            <PlusSquareIcon width={16} height={16} />
            Add Category
          </Link>
        </Button>
      </div>
      {isLoading ? (
        <div>
          <div className="flex bg-gray-100 dark:bg-gray-800 font-semibold h-12">
            <div className="flex-1 flex items-center px-4">Category Name</div>
            <div className="w-24 flex items-center justify-center">Image</div>
            <div className="w-24 flex items-center justify-center">Actions</div>
          </div>
          {[...Array(5)].map((_, index) => (
            <div key={index} className="flex items-center h-16 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors">
              <Skeleton className="h-8 w-8 mr-2 ml-4" />
              <Skeleton className="h-4 w-40 flex-1" />
              <div className="w-24 flex justify-center">
                <Skeleton className="h-10 w-10 rounded" />
              </div>
              <div className="w-24 flex justify-end">
                <Skeleton className="h-8 w-8 mr-2" />
                <Skeleton className="h-8 w-8 mr-4" />
              </div>
            </div>
          ))}
        </div>
      ) : (
        <CategoryTree categories={categories} />
      )}
    </div>
  )
}