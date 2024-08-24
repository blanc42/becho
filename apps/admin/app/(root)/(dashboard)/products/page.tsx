"use client"

import React, { useState, useEffect } from 'react';
import { useStoreData } from '@/lib/store/useStoreData';
import ProductTable from './ProductTable';
import { Button } from "@/components/ui/button";
import { Plus } from 'lucide-react';
import Link from 'next/link';

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
    console.log('Preview product:', product);
    // Implement preview logic
  };

  const handleEdit = (product: Product) => {
    console.log('Edit product:', product);
    // Implement edit logic
  };

  const handleDelete = (product: Product) => {
    console.log('Delete product:', product);
    // Implement delete logic
  };

  return (
    <div className="p-4">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Products</h1>
        <Button asChild>
          <Link href="/products/add">
            <Plus className="w-4 h-4 mr-2" />
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