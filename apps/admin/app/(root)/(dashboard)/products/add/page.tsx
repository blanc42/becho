"use client"

import { useState, useEffect, useCallback } from 'react';
import { useForm, useFieldArray } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { createProductSchema, createProductType, itemType, variantOptionType } from '@/lib/schema';
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Checkbox } from "@/components/ui/checkbox";
import { Textarea } from "@/components/ui/textarea";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import CategorySingleSelector from '@/components/CategoryTree';
import VariantMultiSelector from '@/components/VariantMultiSelector';
import { Tag, TagLabel, TagCloseButton } from "@/components/ui/tag";
import { Variant, Category } from "@/lib/types";
import { X, Plus, Trash2, PlusSquareIcon, ArrowLeft } from 'lucide-react';
import { useStoreData } from '@/lib/store/useStoreData';
import { toast } from '@/components/ui/use-toast';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import LoadingPage from '@/components/pages/LoadingPage';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";

export default function AddProductPage() {
  const [selectedVariants, setSelectedVariants] = useState<Variant[]>([]);
  const [categoryVariants, setCategoryVariants] = useState<Variant[]>([]);
  const [removedOptions, setRemovedOptions] = useState<{ [key: string]: { id: string, value: string }[] }>({});
  const [removedCombinations, setRemovedCombinations] = useState<Record<string, string>[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [variants, setVariants] = useState<Variant[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const { selectedStore } = useStoreData();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const router = useRouter();

  const form = useForm<createProductType>({
    resolver: zodResolver(createProductSchema),
    defaultValues: {
      name: '',
      description: '',
      is_featured: false,
      is_archived: false,
      category_id: '',
      variants: [],
      items: [{
        sku: '',
        quantity: 0,
        price: 0,
        cost_price: 0,
        discounted_price: 0,
        variant_options: {},
      }],
    },
  });

  const { fields, replace, remove } = useFieldArray({
    control: form.control,
    name: "items",
  });

  useEffect(() => {
    const fetchData = async () => {
      if (selectedStore) {
        setIsLoading(true);
        try {
          const [categoriesResponse, variantsResponse] = await Promise.all([
            fetch(`/api/v1/stores/${selectedStore.id}/categories`),
            fetch(`/api/v1/stores/${selectedStore.id}/variants`)
          ]);

          if (!categoriesResponse.ok || !variantsResponse.ok) {
            throw new Error('Failed to fetch data');
          }

          const [categoriesData, variantsData] = await Promise.all([
            categoriesResponse.json(),
            variantsResponse.json()
          ]);

          setCategories(categoriesData);
          setVariants(variantsData);
        } catch (error) {
          console.error('Error fetching data:', error);
        } finally {
          setIsLoading(false);
        }
      }
    };

    fetchData();
  }, [selectedStore]);

  const onSubmit = async (data: createProductType) => {
    console.log(data);
    // Handle form submission
    setIsSubmitting(true);
    try {
      if (selectedStore) {
        const response = await fetch(`/api/v1/stores/${selectedStore.id}/products`, {
          method: 'POST',
          body: JSON.stringify(data),
          credentials: 'include',
        });
        if (!response.ok) {
          throw new Error('Failed to submit form');
        }
        const json = await response.json();
        console.log(json);
        toast({
          title: 'Product created',
          description: `Product ${json.id} created successfully`,
        });
      }
      router.push(`/products`);
    } catch (error) {
      console.error('Error submitting form:', error);
      toast({
        title: 'Error',
        description: 'Failed to create product',
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleCategoryChange = (categoryId: string) => {
    const category = categories.find(c => c.id === categoryId);
    if (category) {
      const newCategoryVariants = variants.filter(v => category.variants.includes(v.id));
      setCategoryVariants(newCategoryVariants);
      setSelectedVariants(prevVariants => {
        const nonCategoryVariants = prevVariants.filter(v => !category.variants.includes(v.id));
        const updatedCategoryVariants = newCategoryVariants.map(variant => ({
          ...variant,
          options: [variant.options[0]] // Select only the first option
        }));
        return [...nonCategoryVariants, ...updatedCategoryVariants];
      });

      // Update removedOptions for category variants
      newCategoryVariants.forEach(variant => {
        setRemovedOptions(prev => ({
          ...prev,
          [variant.id]: variant.options.slice(1)
        }));
      });
    } else {
      // If category is deselected, move category variants to selected variants
      setSelectedVariants(prevVariants => [
        ...prevVariants,
        ...categoryVariants.filter(cv => !prevVariants.some(pv => pv.id === cv.id))
      ]);
      setCategoryVariants([]);
    }
    form.setValue('category_id', categoryId);
  };

  const handleDeleteVariantOption = (variantId: string, optionId: string) => {
    setSelectedVariants(prevVariants =>
      prevVariants.map(variant => {
        if (variant.id === variantId) {
          if (categoryVariants.some(cv => cv.id === variantId) && variant.options.length <= 1) {
            return variant; // Don't remove the last option for category variants
          }
          const removedOption = variant.options.find(o => o.id === optionId);
          if (removedOption) {
            setRemovedOptions(prev => {
              const existingOptions = prev[variantId] || [];
              // Check if the option is already added to prevent duplicates
              if (!existingOptions.some(option => option.id === removedOption.id)) {
                return {
                  ...prev,
                  [variantId]: [...existingOptions, removedOption]
                };
              }
              return prev;
            });
          }
          return { ...variant, options: variant.options.filter(option => option.id !== optionId) };
        }
        return variant;
      }).filter(variant => categoryVariants.some(cv => cv.id === variant.id) || variant.options.length > 0)
    );
  };


  const handleAddVariantOption = (variantId: string, option: any) => {
    setSelectedVariants(prevVariants =>
      prevVariants.map(variant =>
        variant.id === variantId
          ? { ...variant, options: [...variant.options, option] }
          : variant
      )
    );
    setRemovedOptions(prev => ({
      ...prev,
      [variantId]: (prev[variantId] || []).filter(o => o.id !== option.id)
    }));
  };

  const handleVariantSelect = (variantId: string) => {
    const variant = variants.find((v) => v.id === variantId);
    if (variant) {
      if (selectedVariants.some((v) => v.id === variantId)) {
        // If it's not a category variant, remove it
        if (!categoryVariants.some((cv) => cv.id === variantId)) {
          setSelectedVariants(selectedVariants.filter((v) => v.id !== variantId));
        }
      } else {
        const newVariant = {
          ...variant,
          options: [variant.options[0]] // Select only the first option
        };
        setRemovedOptions(prev => ({
          ...prev,
          // TODO: make it better
          [variantId]: variant.options.slice(1)
          // (prev[variantId] || []).filter(o => o.id !== option.id) 
        }));
        setSelectedVariants([...selectedVariants, newVariant]);
      }
    }
  };

  const generateCombinations = useCallback((variants: Variant[]): Record<string, string>[] => {
    const combinations: Record<string, string>[] = [{}];

    variants.forEach(variant => {
      const newCombinations: Record<string, string>[] = [];
      combinations.forEach(combo => {
        variant.options.forEach(option => {
          newCombinations.push({ ...combo, [variant.id]: option.id });
        });
      });
      combinations.splice(0, combinations.length, ...newCombinations);
    });

    return combinations;
  }, []);

  const handleRemoveCombination = useCallback((index: number) => {
    if (fields.length > 1) {
      const removedCombination = fields[index].variant_options;
      setRemovedCombinations(prev => [...prev, removedCombination]);
      remove(index);
    }
  }, [fields, remove]);

  const handleAddCombination = useCallback((combination: Record<string, string>) => {
    const newItem = {
      sku: Object.values(combination).join('-'),
      quantity: 0,
      price: 0,
      cost_price: 0,
      discounted_price: 0,
      variant_options: combination,
    };
    form.setValue('items', [...fields, newItem]);
    setRemovedCombinations(prev => prev.filter(c => JSON.stringify(c) !== JSON.stringify(combination)));
  }, [fields, form]);

  const handleSetAllPrices = useCallback((field: keyof (itemType)) => {
    const value = form.getValues(`items.0.${field}`);
    fields.forEach((_, index) => {
      form.setValue(`items.${index}.${field}`, value);
    });
  }, [fields, form]);

  const generateSKU = () => {
    return `sku-${Math.floor(100000000 + Math.random() * 900000000)}`;
  };

  useEffect(() => {
    const combinations = generateCombinations(selectedVariants);
    const currentItems = form.getValues('items');
    const newItems = combinations.map(combo => {
      const existingItem = currentItems.find(item =>
        Object.entries(combo).every(([key, value]) => item.variant_options[key] === value)
      );
      if (existingItem) {
        return existingItem;
      }
      return {
        sku: generateSKU(),
        quantity: 0,
        price: 0,
        cost_price: 0,
        discounted_price: 0,
        variant_options: combo,
      };
    });
    replace(newItems.length > 0 ? newItems : [{
      sku: generateSKU(),
      quantity: 0,
      price: 0,
      cost_price: 0,
      discounted_price: 0,
      variant_options: {},
    }]);
    form.setValue('variants', selectedVariants.map(v => v.id));
    setRemovedCombinations([]);
  }, [selectedVariants, replace, form, generateCombinations]);

  if (isLoading) {
    return <div className='w-full'>
      <LoadingPage />
    </div>
  }

  return (
    <div className="p-4 w-full max-w-screen-xl mx-auto">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Add Product</h1>
        <Button asChild>
          <Link href="/products">
            <ArrowLeft className="w-4 h-4 mr-2" />
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
                <FormLabel>Product Name</FormLabel>
                <FormControl>
                  <Input placeholder="Enter product name" {...field} />
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
                  <Textarea placeholder="Enter product description" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <div className="space-y-4 md:space-y-0 md:flex md:space-x-4">
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  <FormField
                    control={form.control}
                    name="is_featured"
                    render={({ field }) => (
                      <FormItem className="flex flex-row items-center space-x-3 space-y-0">
                        <FormControl>
                          <Checkbox
                            checked={field.value}
                            onCheckedChange={field.onChange}
                          />
                        </FormControl>
                        <FormLabel>Featured Product</FormLabel>
                      </FormItem>
                    )}
                  />
                </TooltipTrigger>
                <TooltipContent>
                  <p>Check this if the product should be featured</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>

            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  <FormField
                    control={form.control}
                    name="is_archived"
                    render={({ field }) => (
                      <FormItem className="flex flex-row items-center space-x-3 space-y-0">
                        <FormControl>
                          <Checkbox
                            checked={field.value}
                            onCheckedChange={field.onChange}
                          />
                        </FormControl>
                        <FormLabel>Archived</FormLabel>
                      </FormItem>
                    )}
                  />
                </TooltipTrigger>
                <TooltipContent>
                  <p>Check this if the product is archived</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </div>

          <FormField
            control={form.control}
            name="category_id"
            render={({ field }) => (
              <FormItem>
                <FormControl>
                  <CategorySingleSelector
                    categories={categories}
                    value={field.value}
                    onChange={handleCategoryChange}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <div>
            <VariantMultiSelector
              variants={variants}
              selectedVariants={selectedVariants}
              onVariantSelect={handleVariantSelect}
              categoryVariants={categoryVariants}
            />
            {selectedVariants.map((variant) => (
              <div key={variant.id} className="flex flex-col mt-4 p-4 border rounded">
                <div className="flex flex-wrap gap-2">
                  {variant.options.map((option) => (
                    <Tag key={option.id} variant="default" size="md">
                      <TagLabel className="text-base">{option.value}</TagLabel>
                      {(variant.options.length > 1 || !categoryVariants.some(cv => cv.id === variant.id)) && (
                        <TagCloseButton onClick={() => handleDeleteVariantOption(variant.id, option.id)} />
                      )}
                    </Tag>
                  ))}
                </div>
                {removedOptions[variant.id] && removedOptions[variant.id].length > 0 && (
                  <div className="mt-2">
                    <span className="text-sm font-semibold">Removed options:</span>
                    <div className="flex flex-wrap gap-2 mt-1">
                      {removedOptions[variant.id].map((option) => (
                        <Button
                          key={option.id}
                          variant="outline"
                          size="sm"
                          onClick={() => handleAddVariantOption(variant.id, option)}
                        >
                          <Plus className="h-4 w-4 mr-1" />
                          {option.value}
                        </Button>
                      ))}
                    </div>
                  </div>
                )}
              </div>
            ))}
          </div>

          {selectedVariants.length > 0 ? (
            <div>
              <FormLabel>Product Variants</FormLabel>
              <div className="space-y-4 overflow-x-auto">
                <div className="grid grid-cols-7 gap-4 font-bold min-w-[800px]">
                  <div>Variants</div>
                  <div>SKU</div>
                  <div>Quantity</div>
                  <div>Cost Price</div>
                  <div>Price</div>
                  <div>Discounted Price</div>
                  <div className="text-right">Actions</div>
                </div>
                {fields.map((field, index) => (
                  <div key={field.id} className="grid grid-cols-7 gap-4 items-center min-w-[800px]">
                    <div>
                      {Object.entries(field.variant_options).map(([variantId, optionId]) => {
                        const variant = selectedVariants.find(v => v.id === variantId);
                        const option = variant?.options.find(o => o.id === optionId);
                        return option?.value;
                      }).join(', ')}
                    </div>
                    <FormField
                      control={form.control}
                      name={`items.${index}.sku`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input {...field} />
                          </FormControl>
                        </FormItem>
                      )}
                    />
                    <FormField
                      control={form.control}
                      name={`items.${index}.quantity`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              type="number"
                              {...field}
                              onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                              onFocus={e => e.target.select()}
                            />
                          </FormControl>
                        </FormItem>
                      )}
                    />
                    <FormField
                      control={form.control}
                      name={`items.${index}.cost_price`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              type="number"
                              {...field}
                              onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                              onFocus={e => e.target.select()}
                            />
                          </FormControl>
                        </FormItem>
                      )}
                    />
                    <FormField
                      control={form.control}
                      name={`items.${index}.price`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              type="number"
                              {...field}
                              onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                              onFocus={e => e.target.select()}
                            />
                          </FormControl>
                        </FormItem>
                      )}
                    />
                    <FormField
                      control={form.control}
                      name={`items.${index}.discounted_price`}
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input
                              type="number"
                              {...field}
                              onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                              onFocus={e => e.target.select()}
                            />
                          </FormControl>
                        </FormItem>
                      )}
                    />
                    <div className="text-right">
                      <Button type="button" variant="ghost" onClick={() => handleRemoveCombination(index)} className="p-1">
                        <Trash2 className="h-3 w-5" />
                      </Button>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          ) : (
            <div className="space-y-4">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <FormField
                  control={form.control}
                  name="items.0.sku"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>SKU</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="items.0.quantity"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Quantity</FormLabel>
                      <FormControl>
                        <Input
                          type="number"
                          {...field}
                          onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                          onFocus={e => e.target.select()}
                        />
                      </FormControl>
                    </FormItem>
                  )}
                />
              </div>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <FormField
                  control={form.control}
                  name="items.0.cost_price"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Cost Price</FormLabel>
                      <FormControl>
                        <Input
                          type="number"
                          {...field}
                          onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                          onFocus={e => e.target.select()}
                        />
                      </FormControl>
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="items.0.price"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Price</FormLabel>
                      <FormControl>
                        <Input
                          type="number"
                          {...field}
                          onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                          onFocus={e => e.target.select()}
                        />
                      </FormControl>
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="items.0.discounted_price"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Discounted Price</FormLabel>
                      <FormControl>
                        <Input
                          type="number"
                          {...field}
                          onChange={e => field.onChange(e.target.value === '' ? '' : +e.target.value)}
                          onFocus={e => e.target.select()}
                        />
                      </FormControl>
                    </FormItem>
                  )}
                />
              </div>
            </div>
          )}

          {removedCombinations.length > 0 && (
            <div>
              <FormLabel>Removed Combinations</FormLabel>
              <div className="space-y-2">
                {removedCombinations.map((combination, index) => (
                  <div key={index} className="flex items-center space-x-2">
                    <Button type="button" variant="outline" onClick={() => handleAddCombination(combination)}>
                      <Plus className="h-4 w-4 mr-2" />
                      Add
                    </Button>
                    <span>
                      {Object.entries(combination).map(([variantId, optionId]) => {
                        const variant = selectedVariants.find(v => v.id === variantId);
                        const option = variant?.options.find(o => o.id === optionId);
                        return `${variant?.name}: ${option?.value}`;
                      }).join(', ')}
                    </span>
                  </div>
                ))}
              </div>
            </div>
          )}

          <Button type="submit" disabled={isSubmitting}>
            {
              isSubmitting ? 'Creating Product...' : 'Create Product'
            }
          </Button>
        </form>
      </Form>
    </div>
  );
}