"use client"

import React from 'react';
import { useStoreData } from '@/lib/store/useStoreData';
import ProductTable from './ProductTable';
import { Button } from "@/components/ui/button";
import { PlusSquareIcon } from 'lucide-react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { Product } from '@/lib/types';
import useSWR from 'swr';

const fetcher = (url: string) => fetch(url).then(res => res.json());

export default function ProductsPage() {
  const { selectedStore } = useStoreData();
  const router = useRouter();

  const { data: products, error, isLoading } = useSWR<Product[]>(
    selectedStore ? `/api/v1/stores/${selectedStore.id}/products` : null,
    fetcher
  );

  const handlePreview = (product: Product) => {
    router.push(`/products/${product.id}`);
  };

  const handleEdit = (product: Product) => {
    router.push(`/products/${product.id}/edit`);
  };

  const handleDelete = (product: Product) => {
    console.log('Delete product:', product);
    // Implement delete logic
  };

  return (
    <div className="container max-w-screen-xl px-0">
      <div className="flex justify-between items-center my-4">
        <h1 className="text-2xl font-bold">Products</h1>
        <Button asChild>
          <Link href="/products/add">
            <PlusSquareIcon className="w-4 h-4 mr-2" />
            Add Product
          </Link>
        </Button>
      </div>
      <ProductTable
        products={products || []}
        isLoading={isLoading}
        onPreview={handlePreview}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
      {error && <div>Failed to load products</div>}
    </div>
  );
}