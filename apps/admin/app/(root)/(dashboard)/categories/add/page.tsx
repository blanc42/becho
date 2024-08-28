"use client"

import { useState, useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
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
import CategoryTree from '@/components/CategoryTree';
import VariantMultiSelector from '@/components/VariantMultiSelector';
import { createCategorySchema, CreateCategoryType } from '@/lib/schema';
import { Category, Variant } from '@/lib/types';
import { Badge } from "@/components/ui/badge";
import { ArrowLeft, MoveLeftIcon, X } from "lucide-react";
import { useRouter } from 'next/navigation';
import { toast } from '@/components/ui/use-toast';
import Link from 'next/link';

export default function CategoryAddPage() {
    const [selectedVariants, setSelectedVariants] = useState<Variant[]>([]);
    const [categoryVariants, setCategoryVariants] = useState<Variant[]>([]);
    const [categories, setCategories] = useState<Category[]>([]);
    const [variants, setVariants] = useState<Variant[]>([]);
    const { selectedStore } = useStoreData();
    const [isLoading, setIsLoading] = useState(false);
    const router = useRouter();

    const form = useForm<CreateCategoryType>({
        resolver: zodResolver(createCategorySchema),
        defaultValues: {
            name: '',
            description: '',
            parent_id: null,
            variants: [],
            unique_identifier: '',
        },
    });

    useEffect(() => {
        if (selectedStore) {
            fetchCategories();
            fetchVariants();
        }
    }, [selectedStore]);

    const fetchCategories = async () => {
        try {
            const response = await fetch(`/api/v1/stores/${selectedStore?.id}/categories`);
            if (response.ok) {
                const data = await response.json();
                setCategories(data);
            }
        } catch (error) {
            console.error('Error fetching categories:', error);
        }
    };

    const fetchVariants = async () => {
        try {
            const response = await fetch(`/api/v1/stores/${selectedStore?.id}/variants`);
            if (response.ok) {
                const data = await response.json();
                setVariants(data);
            }
        } catch (error) {
            console.error('Error fetching variants:', error);
        }
    };

    const onSubmit = async (values: CreateCategoryType) => {
        if (!selectedStore) return;
        setIsLoading(true);

        try {
            if (selectedStore) {
                const response = await fetch(`/api/v1/stores/${selectedStore.id}/categories`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        ...values,
                        variants: [...categoryVariants, ...selectedVariants].map(v => v.id),
                    }),
                    credentials: 'include'
                });

                if (!response.ok) {
                    throw new Error('Failed to create category');
                }
                toast({
                    title: `Category created successfully`,
                    description: "Category created successfully",
                });
                router.push(`/categories`);
            }
        } catch (error) {
            console.error('Error creating category:', error);
        } finally {
            setIsLoading(false);
        }
    };

    const handleVariantSelect = (variantId: string) => {
        const variant = variants.find(v => v.id === variantId);
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

    const handleCategorySelect = (categoryId: string | null) => {
        if (categoryId) {
            const category = categories.find(c => c.id === categoryId);
            if (category) {
                const categoryVariants = category.variants.map(v => v.id)
                const categoryVariantsFiltered = variants.filter(v => categoryVariants.includes(v.id));
                setCategoryVariants(categoryVariantsFiltered);
            }
        } else {
            setCategoryVariants([]);
        }
        form.setValue('parent_id', categoryId);
    };

    const removeSelectedVariant = (variantId: string) => {
        setSelectedVariants(prev => prev.filter(v => v.id !== variantId));
    };

    return (
        <div className="space-y-6 max-w-screen-lg w-full mx-auto">
            <div className='flex items-baseline justify-between mb-8 border-b border-gray-600 py-4 w-full'>
                <h1 className="text-3xl font-semibold">Add Category</h1>
                <Button asChild size="sm">
                    <Link href="/categories" className='flex items-center gap-2'>
                        <MoveLeftIcon width={16} height={16} />
                        Back
                    </Link>
                </Button>
            </div>
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
                        name="unique_identifier"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Unique Identifier</FormLabel>
                                <FormControl>
                                    <Input placeholder="Unique identifier" {...field} />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="parent_id"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Parent Category</FormLabel>
                                <FormControl className='w-full'>
                                    <CategoryTree
                                        categories={categories}
                                        value={field.value}
                                        onChange={(value) => handleCategorySelect(value)}
                                    />
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
                                        variants={variants}
                                        selectedVariants={[...categoryVariants, ...selectedVariants]}
                                        onVariantSelect={handleVariantSelect}
                                        categoryVariants={categoryVariants}
                                    />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <div>
                        <h3 className="text-sm font-medium mb-2">Selected Variants:</h3>
                        <div className="flex flex-wrap gap-2">
                            {categoryVariants.map(variant => (
                                <Badge key={variant.id} variant="secondary">
                                    {variant.name}
                                </Badge>
                            ))}
                            {selectedVariants.map(variant => (
                                <Badge key={variant.id} variant="secondary" className="flex items-center gap-1">
                                    {variant.name}
                                    <X
                                        className="h-3 w-3 cursor-pointer"
                                        onClick={() => removeSelectedVariant(variant.id)}
                                    />
                                </Badge>
                            ))}
                        </div>
                    </div>

                    <Button type="submit" disabled={isLoading}>
                        {isLoading ? 'Creating...' : 'Create Category'}
                    </Button>
                </form>
            </Form>
        </div>
    );
}