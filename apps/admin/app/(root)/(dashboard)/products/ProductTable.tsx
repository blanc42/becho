"use client";
import React from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { Eye, Edit, Trash2 } from 'lucide-react';
import { Skeleton } from "@/components/ui/skeleton";
import { useRouter } from 'next/navigation';

interface ProductVariant {
  id: string;
  sku: string;
  price: number;
  quantity: number;
  images : string[];
}

interface Product {
  id: string;
  name: string;
  description: string | null;
  is_featured: boolean;
  is_archived: boolean;
  has_variants: boolean;
  category_name: string;
  product_variants: ProductVariant[];
}

interface ProductTableProps {
  products: Product[];
  isLoading: boolean;
  onPreview: (product: Product) => void;
  onEdit: (product: Product) => void;
  onDelete: (product: Product) => void;
}

const ProductTable: React.FC<ProductTableProps> = ({ products, isLoading, onPreview, onEdit, onDelete }) => {
  if (isLoading) {
    return (
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Category</TableHead>
            <TableHead>Price</TableHead>
            <TableHead>Quantity</TableHead>
            <TableHead>Images</TableHead>
            <TableHead className="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {[...Array(5)].map((_, index) => (
            <TableRow key={index}>
              <TableCell><Skeleton className="h-4 w-[200px]" /></TableCell>
              <TableCell><Skeleton className="h-4 w-[100px]" /></TableCell>
              <TableCell><Skeleton className="h-4 w-[80px]" /></TableCell>
              <TableCell><Skeleton className="h-4 w-[60px]" /></TableCell>
              <TableCell><Skeleton className="h-8 w-[80px] ml-auto" /></TableCell>
              <TableCell><Skeleton className="h-8 w-[60px] ml-auto" /></TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    );
  }

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Category</TableHead>
          <TableHead>Price</TableHead>
          <TableHead>Quantity</TableHead>
          <TableHead>Images</TableHead>
          <TableHead className="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {products.map((product) => (
          <TableRow key={product.id}>
            <TableCell>{product.name}</TableCell>
            <TableCell>{product.category_name}</TableCell>
            <TableCell>
              {product.product_variants.length > 0
                ? `$${product.product_variants[0].price.toFixed(2)}`
                : 'N/A'}
            </TableCell>
            <TableCell>
              {product.product_variants.length > 0
                ? product.product_variants[0].quantity
                : 'N/A'}
            </TableCell>
            <TableCell className='flex items-baseline space-x-2'>
            {product.product_variants[0].images.slice(0, 3).map((image, index) => (
              <img key={image} src={`https://ucarecdn.com/${image}/-/scale_crop/50x50/center/`} alt={`Product Image ${index}`} width={50} height={50} />
            ))}
            </TableCell>
            <TableCell>
              <div className="flex space-x-2 justify-end">
                <Button  variant="ghost" size="sm" onClick={() => onPreview(product)}>
                  <Eye className="w-4 h-4" />
                </Button>
                <Button variant="ghost" size="sm" onClick={() => onEdit(product)}>
                  <Edit className="w-4 h-4" />
                </Button>
                <Button variant="ghost" size="sm" onClick={() => onDelete(product)}>
                  <Trash2 className="w-4 h-4" />
                </Button>
              </div>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

export default ProductTable;