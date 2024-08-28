"use client"
import { Button } from "@/components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import Link from "next/link";
import { useStoreData } from "@/lib/store/useStoreData";
import { Skeleton } from "@/components/ui/skeleton";
import { Edit, PlusSquareIcon, Trash2 } from 'lucide-react';
import useSWR from 'swr';

interface Option {
  id: string;
  value: string;
  display_order: number;
  image_id: string | null;
}

interface Variant {
  id: string;
  name: string;
  label: string;
  options: Option[];
}

const fetcher = async (url: string) => {
  const response = await fetch(url, {
    credentials: 'include'
  });
  if (!response.ok) {
    throw new Error('Failed to fetch variants');
  }
  return response.json();
};

export default function VariantsPage() {
  const { selectedStore } = useStoreData();

  const { data: variants, error, isLoading } = useSWR<Variant[]>(
    selectedStore ? `/api/v1/stores/${selectedStore.id}/variants` : null,
    fetcher
  );

  return (
    <div className="container max-w-screen-xl px-2">
      <div className="flex items-baseline justify-between mb-8">
        <h1 className="text-3xl font-semibold">Variants</h1>
        <Button size='sm' asChild>
          <Link href="/variants/add" className='flex items-center gap-2'>
            <PlusSquareIcon width={16} height={16} />
            Add Variant 
          </Link>
        </Button>
      </div>
      <div className="w-full">
        {isLoading ? (
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Label</TableHead>
                <TableHead>Options</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {Array.from({ length: 5 }).map((_, index) => (
                <TableRow key={index}>
                  <TableCell><Skeleton className="h-4 w-[200px]" /></TableCell>
                  <TableCell><Skeleton className="h-4 w-[100px]" /></TableCell>
                  <TableCell><Skeleton className="h-4 w-[300px]" /></TableCell>
                  <TableCell><Skeleton className="h-4 w-[100px] ml-auto" /></TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        ) : error ? (
          <div className="text-center py-8">
            <p className="text-lg text-red-500">Error loading variants. Please try again later.</p>
          </div>
        ) : variants && variants.length > 0 ? (
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Label</TableHead>
                <TableHead>Options</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {variants.map((variant) => (
                <TableRow key={variant.id}>
                  <TableCell className="font-medium">{variant.name}</TableCell>
                  <TableCell>{variant.label}</TableCell>
                  <TableCell>
                    {variant.options.slice(0, 3).map(option => option.value).join(", ")}
                    {variant.options.length > 3 && "..."}
                  </TableCell>
                  <TableCell className="text-right">
                    <div className="flex justify-end space-x-2">
                      <Button variant="ghost" size="sm" onClick={() => console.log("Edit")}>
                        <Edit className="w-4 h-4" />
                      </Button>
                      <Button variant="ghost" size="sm" onClick={() => console.log("Delete")}>
                        <Trash2 className="w-4 h-4 text-red-500" />
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
    </div>
  )
}