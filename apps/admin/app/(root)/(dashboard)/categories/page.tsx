"use client"
import { Suspense, useState } from "react";
import { useStoreData } from "@/lib/store/useStoreData";
import { Skeleton } from "@/components/ui/skeleton";
import CategoryTree from "./CategoryTree";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { PlusSquareIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { Category } from "@/lib/types";
import useSWR from "swr";

export default function CategoriesPage() {
  const { selectedStore } = useStoreData();

  const fetcher = async (url: string) => {
    const res = await fetch(url, {
      headers: {
        "Content-Type": "application/json",
      },
      credentials: 'include'
    });
    if (!res.ok) {
      throw new Error('Failed to fetch categories');
    }
    return res.json();
  };

  const { data: categories, error, isLoading } = useSWR<Category[]>(
    selectedStore ? `/api/v1/stores/${selectedStore.id}/categories` : null,
    fetcher,
    // { suspense: true }
  );

  return (
    <div className="container max-w-screen-xl px-2">
      <div className="flex items-baseline justify-between my-12">
        <h1 className="text-3xl font-semibold">Categories</h1>
        <Button asChild>
          <Link href="/categories/add" className='flex items-center gap-2'>
            <PlusSquareIcon width={16} height={16} />
            Add Category
          </Link>
        </Button>
      </div>
      { isLoading ? (
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
      ) :  error ? (
        <div className="text-center py-8">
          <p className="text-lg text-red-500">Error loading categories. Please try again later.</p>
        </div>
      ) : categories && categories.length > 0 ? (
        <CategoryTree categories={categories} />
      ) : (
        <div className="text-center py-8">
          <p className="text-lg text-gray-500">No categories found. Click the "Add Category" button to create your first category.</p>
        </div>
      )}
    </div>
  )
}