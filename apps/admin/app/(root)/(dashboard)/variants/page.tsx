"use client"
import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import Link from "next/link";
import { useStore } from "@/lib/store/useStore";
import { Skeleton } from "@/components/ui/skeleton";
import { Eye, Edit, Trash2 } from 'lucide-react';

interface Option {
  id: string;
  value: string;
  display_order: number;
  image_id: string | null;
}

interface Variant {
  id: string;
  name: string;
  options: Option[];
}

export default function VariantsPage() {
  const [variants, setVariants] = useState<Variant[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { selectedStore } = useStore();

  useEffect(() => {
    setIsLoading(true);
    const fetchVariants = async () => {
      if (selectedStore) {
        try {
          const response = await fetch(`/api/v1/stores/${selectedStore.id}/variants`);
          if (!response.ok) {
            throw new Error('Failed to fetch variants');
          }
          const data = await response.json();
          setVariants(data);
        } catch (error) {
          console.error('Error fetching variants:', error);
        } finally {
          setIsLoading(false);
        }
      }
    };

    fetchVariants();
  }, [selectedStore]);

  return (
    <>
      <div className="flex items-baseline justify-between mb-6">
        <h1 className="text-3xl font-semibold">Variants</h1>
        <Button asChild>
          <Link href="/variants/add">
            Add Variant
          </Link>
        </Button>
      </div>
      <div>
        {isLoading ? (
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Options</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {Array.from({ length: 5 }).map((_, index) => (
                <TableRow key={index}>
                  <TableCell><Skeleton className="h-4 w-[200px]" /></TableCell>
                  <TableCell><Skeleton className="h-4 w-[300px]" /></TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        ) : variants.length > 0 ? (
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>ID</TableHead>
                <TableHead>Name</TableHead>
                <TableHead>Options</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {variants.map((variant) => (
                <TableRow key={variant.id}>
                  <TableCell>{variant.id}</TableCell>
                  <TableCell>{variant.name}</TableCell>
                  <TableCell>
                    {variant.options.map(option => option.value).join(", ")}
                  </TableCell>
                  <TableCell>
              <div className="flex space-x-2">
                <Button variant="ghost" size="sm" onClick={() => console.log("Preview")}>
                  <Eye className="w-4 h-4" />
                </Button>
                <Button variant="ghost" size="sm" onClick={() => console.log("Edit")}>
                  <Edit className="w-4 h-4" />
                </Button>
                <Button variant="ghost" size="sm" onClick={() => console.log("Delete")}>
                  <Trash2 className="w-4 h-4" />
                </Button>
              </div>
            </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        ) : (
          <div className="text-center py-8">
            <p className="text-lg text-gray-500">No variants found. Click the "Add Variant" button to create your first variant.</p>
          </div>
        )}
      </div>
    </>
  )
}