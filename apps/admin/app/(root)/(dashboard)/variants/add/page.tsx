"use client"
import { useState } from 'react';
import { useForm, useFieldArray } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { useRouter } from "next/navigation";
import { useStoreData } from '@/lib/store/useStoreData';
import { MoveLeftIcon, Trash2, Trash2Icon } from 'lucide-react';
import Link from 'next/link';

const variantSchema = z.object({
  name: z.string().min(3, 'Variant name must be at least 3 characters'),
  description: z.string().optional(),
  label: z.string(),
  options: z.array(
    z.object({
      value: z.string().min(1, 'Variant option must be at least 3 characters'),
      data: z.string().optional(),
      imageId: z.string().optional(),
    })
  ).min(1, 'At least one option is required'),
});

type VariantFormValues = z.infer<typeof variantSchema>;

export default function AddVariant() {
  const { selectedStore } = useStoreData()
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();
  const form = useForm<VariantFormValues>({
    resolver: zodResolver(variantSchema),
    defaultValues: {
      name: '',
      description: '',
      label: '',
      options: [{ value: '', data: '', imageId: '' }],
    },
  });

  const { fields, append, remove } = useFieldArray({
    control: form.control,
    name: 'options',
  });

  function onSubmit(values: VariantFormValues) {
    setIsLoading(true);
    fetch(`/api/v1/stores/${selectedStore?.id}/variants`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'credentials': 'include'
      },
      body: JSON.stringify(values),
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('Variant creation failed');
        }
        return response.json();
      })
      .then(res => {
        console.log('Variant created successfully:', res);
        router.push('/variants');
      })
      .catch(error => {
        console.error('Error:', error);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }

  return (
    <div className='flex flex-col items-center justify-center mt-12 max-w-2xl mx-auto'>
      <div className='flex items-baseline justify-between mb-6 border-b border-gray-600 py-4 w-full'>
        <h1 className="text-3xl font-semibold">Add Variant</h1>
        <Button asChild>
          <Link href="/variants" className='flex items-center gap-2'>
            <MoveLeftIcon width={16} height={16} />
            Back to Variants
          </Link>
        </Button>
      </div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 w-full">
          <div className="flex space-x-4">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem className="flex-1">
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Variant Name" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="label"
              render={({ field }) => (
                <FormItem className="flex-1">
                  <FormLabel>Label</FormLabel>
                  <FormControl>
                    <Input placeholder="Variant Label" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Description</FormLabel>
                <FormControl>
                  <Textarea placeholder="Variant Description" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="space-y-4">
            <div className="flex justify-between items-center">
              <FormLabel>Options</FormLabel>
              <Button 
                type="button" 
                onClick={() => append({ value: '', data: '', imageId: '' })} 
                size="sm"
              >
                Add Option
              </Button>
            </div>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Value</TableHead>
                  <TableHead>Data</TableHead>
                  <TableHead>Image ID</TableHead>
                  <TableHead>Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {fields.map((field, index) => (
                  <TableRow key={field.id}>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`options.${index}.value`}
                        render={({ field }) => (
                          <FormItem>
                            <FormControl>
                              <Input placeholder="Option Value" {...field} />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </TableCell>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`options.${index}.data`}
                        render={({ field }) => (
                          <FormItem>
                            <FormControl>
                              <Input placeholder="Option Data" {...field} />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </TableCell>
                    <TableCell>
                      <FormField
                        control={form.control}
                        name={`options.${index}.imageId`}
                        render={({ field }) => (
                          <FormItem>
                            <FormControl>
                              <Input placeholder="Image ID" {...field} />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </TableCell>
                    <TableCell>
                      <Button 
                        type="button" 
                        variant="destructive" 
                        onClick={() => fields.length > 1 && remove(index)}
                        size="sm"
                        disabled={fields.length === 1}
                      >
                        <Trash2Icon width={16} height={16} />
                      </Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
          <Button type="submit" disabled={isLoading} className='mr-0'>
            {isLoading ? 'Creating...' : 'Create Variant'}
          </Button>
        </form>
      </Form>
    </div>
  );
}