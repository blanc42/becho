"use client"
import { useEffect, useState } from "react";
import { useStoreData } from "@/lib/store/useStoreData";
import { Skeleton } from "@/components/ui/skeleton";
import CategoryTree from "./CategoryTree";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { PlusSquareIcon } from "lucide-react";
import { useRouter } from "next/navigation";

interface Category {
  id: string;
  name: string;
  description: string;
  level: number;
  parent_id: string | null;
  unique_identifier: string;
  variants: string[];
}

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
        <ul className="pl-4">
          {[...Array(5)].map((_, index) => (
            <li key={index} className="mb-2">
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <Skeleton className="h-6 w-6 mr-2" />
                  <Skeleton className="h-4 w-40" />
                </div>
                <div className="flex items-center">
                  <Skeleton className="h-6 w-6 mr-2" />
                  <Skeleton className="h-6 w-6" />
                </div>
              </div>
            </li>
          ))}
        </ul>
      ) : (
        <CategoryTree categories={categories} />
      )}
    </div>
  )
}