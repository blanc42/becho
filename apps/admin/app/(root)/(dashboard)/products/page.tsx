"use client"

import React, { useState, useEffect } from 'react';
import { useStoreData } from '@/lib/store/useStoreData';
import ProductTable from './ProductTable';
import { Button } from "@/components/ui/button";
import { Plus, PlusSquareIcon } from 'lucide-react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

interface Product {
  id: string;
  name: string;
  description: string | null;
  is_featured: boolean;
  is_archived: boolean;
  has_variants: boolean;
  category_name: string;
  product_variants: any[];
}

export default function ProductsPage() {
  const [products, setProducts] = useState<Product[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { selectedStore } = useStoreData();
  const router = useRouter()

  useEffect(() => {
    const fetchProducts = async () => {
      if (selectedStore) {
        setIsLoading(true);
        try {
          const response = await fetch(`/api/v1/stores/${selectedStore.id}/products`);
          if (!response.ok) {
            throw new Error('Failed to fetch products');
          }
          const data = await response.json();
          setProducts(data);
        } catch (error) {
          console.error('Error fetching products:', error);
        } finally {
          setIsLoading(false);
        }
      }
    };

    fetchProducts();
  }, [selectedStore]);

  const handlePreview = (product: Product) => {
      router.push(`/products/${product.id}`)
  };

  const handleEdit = (product: Product) => {
    console.log('Edit product:', product);
    router.push(`/products/${product.id}/edit`)
  };

  const handleDelete = (product: Product) => {
    console.log('Delete product:', product);
    // Implement delete logic
  };

  return (
    <div className="w-full max-w-screen-xl mx-auto">
      <div className="flex justify-between items-center my-12">
        <h1 className="text-2xl font-bold">Products</h1>
        <Button asChild>
          <Link href="/products/add">
            <PlusSquareIcon className="w-4 h-4 mr-2" />
            Add Product
          </Link>
        </Button>
      </div>
      <ProductTable
        products={products}
        isLoading={isLoading}
        onPreview={handlePreview}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
    </div>
  );
}