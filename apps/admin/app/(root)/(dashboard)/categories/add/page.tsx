"use client"

import { useState } from 'react';
import { useForm } from 'react-hook-form';
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
import { useStoreData } from '@/lib/store/useStoreData';
import { useCategoriesStore } from '@/lib/store/useCategoriesStore';
import CategorySingleSelector from '@/components/CategoryTree';
import { Variant } from '@/lib/types';
import CategoryTree from '@/components/CategoryTree';
import { createCategorySchema, createCategoryType } from '@/lib/schema';
import { categoriesData, variantsData } from '@/data/variants';
import VariantMultiSelector from '@/components/VariantMultiSelector';



export default function CategoryAddPage() {
    const [selectedVariants, setSelectedVariants] = useState<Variant[]>([]);
    const { selectedStore } = useStoreData();
    const [isLoading, setIsLoading] = useState(false);

    const form = useForm<createCategoryType>({
        resolver: zodResolver(createCategorySchema),
        defaultValues: {
            name: '',
            description: '',
            parentCategoryId: '',
            variants: [],
        },
    });

    const onSubmit = async (values: createCategoryType) => {
        if (!selectedStore) return;
        setIsLoading(true);

        try {
            const response = await fetch(`/api/v1/stores/${selectedStore.id}/categories`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(values),
                credentials: 'include'
            });

            if (!response.ok) {
                throw new Error('Failed to create category');
            }


            console.log({...values, variants: selectedVariants.map(v => v.id)})
            console.log('Category created successfully');
        } catch (error) {
            console.error('Error creating category:', error);
        } finally {
            setIsLoading(false);
        }
    };

    const handleVariantSelect = (variantId: string) => {
        const variant = variantsData.find(v => v.id === variantId);
        if (variant) {
            setSelectedVariants(prev => {
                const isSelected = prev.some(v => v.id === variantId);
                if (isSelected) {
                    return prev.filter(v => v.id !== variantId);
                } else {
                    return [...prev, variant];
                }
            });
        }
    };

    return (
        <div className="space-y-6">
            <h1 className="text-2xl font-bold">Create Category</h1>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Name</FormLabel>
                                <FormControl>
                                    <Input placeholder="Category name" {...field} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="description"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Description</FormLabel>
                                <FormControl>
                                    <Textarea placeholder="Category description" {...field} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="parentCategoryId"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Parent Category</FormLabel>
                                <FormControl className='w-full'>
                                    <CategoryTree categories={categoriesData} value={field.value} onChange={field.onChange} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="variants"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Variants</FormLabel>
                                <FormControl>
                                    <VariantMultiSelector
                                        variants={variantsData}
                                        selectedVariants={selectedVariants}
                                        onVariantSelect={handleVariantSelect}
                                        categoryVariants={[]}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <Button type="submit" disabled={isLoading}>
                        {isLoading ? 'Creating...' : 'Create Category'}
                    </Button>
                </form>
            </Form>
        </div>
    );
}