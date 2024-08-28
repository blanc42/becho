"use client"

import { useEffect, useState } from 'react';
import { useParams } from 'next/navigation';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import { Button } from "@/components/ui/button";
import { z } from 'zod';
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious, type CarouselApi } from "@/components/ui/carousel";
import { useStoreData } from '@/lib/store/useStoreData';
import useSWR from 'swr';
import { Skeleton } from "@/components/ui/skeleton";

const VariantOptionSchema = z.object({
  data: z.string(),
  id: z.string(),
  value: z.string(),
});

const VariantSchema = z.object({
  name: z.string(),
  options: z.array(VariantOptionSchema),
  variant_id: z.string(),
});

const ProductVariantSchema = z.object({
  cost_price: z.number(),
  created_at: z.string(),
  discounted_price: z.number(),
  images: z.array(z.string()),
  options: z.record(z.string(), z.string()),
  price: z.number(),
  quantity: z.number(),
  sku: z.string(),
  updated_at: z.string(),
  variant_id: z.string(),
});

const ProductSchema = z.object({
  product_id: z.string(),
  product_name: z.string(),
  product_description: z.string().nullable(),
  rating: z.number().nullable(),
  is_featured: z.boolean(),
  is_archived: z.boolean(),
  category_id: z.string(),
  store_id: z.string(),
  variants_order: z.array(z.string()),
  product_variants: z.array(ProductVariantSchema),
  variants: z.array(VariantSchema),
});

type Product = z.infer<typeof ProductSchema>;
type ProductVariant = z.infer<typeof ProductVariantSchema>;
type Variant = z.infer<typeof VariantSchema>;

const fetcher = (url: string) => fetch(url).then(res => res.json());

export default function ProductPage() {
  const { product_id } = useParams();
  const [selectedVariant, setSelectedVariant] = useState<ProductVariant | null>(null);
  const [selectedOptions, setSelectedOptions] = useState<Record<string, string>>({});
  const { selectedStore } = useStoreData();
  const [availableOptions, setAvailableOptions] = useState<Record<string, string[]>>({});
  const [api, setApi] = useState<CarouselApi>();

  const { data: product, error } = useSWR<Product>(
    selectedStore ? `/api/v1/stores/${selectedStore.id}/products/${product_id}` : null,
    fetcher
  );

  useEffect(() => {
    if (product) {
      // Initialize selected options with the first option of each variant
      const initialSelectedOptions: Record<string, string> = {};
      product.variants.forEach(variant => {
        initialSelectedOptions[variant.name] = variant.options[0].value;
      });

      setSelectedOptions(initialSelectedOptions);
      updateAvailableOptions(product, initialSelectedOptions);

      // Find the selected variant based on initial options
      const initialVariant = product.product_variants.find(v => 
        Object.entries(initialSelectedOptions).every(([key, value]) => v.options[key] === value)
      );
      setSelectedVariant(initialVariant || null);
    }
  }, [product]);

  const updateAvailableOptions = (product: Product, currentSelections: Record<string, string>) => {
    const newAvailableOptions: Record<string, string[]> = {};

    product.variants.forEach((variant, index) => {
      if (index === 0) {
        // All options are available for the first variant
        newAvailableOptions[variant.name] = variant.options.map(opt => opt.value);
      } else {
        // For subsequent variants, check which options are available based on previous selections
        const previousSelections = Object.fromEntries(
          Object.entries(currentSelections).slice(0, index)
        );

        newAvailableOptions[variant.name] = product.product_variants
          .filter(v => 
            Object.entries(previousSelections)
              .every(([key, value]) => v.options[key] === value)
          )
          .map(v => v.options[variant.name])
          .filter((value, i, self) => self.indexOf(value) === i);
      }
    });

    setAvailableOptions(newAvailableOptions);
  };

  const handleOptionSelect = (variantName: string, optionValue: string) => {
    const newSelectedOptions = { ...selectedOptions, [variantName]: optionValue };
    
    // Update selections for subsequent variants if they become invalid
    const updatedSelections = { ...newSelectedOptions };
    const variantIndex = product!.variants.findIndex(v => v.name === variantName);
    for (let i = variantIndex + 1; i < product!.variants.length; i++) {
      const nextVariant = product!.variants[i];
      const availableOptions = product!.product_variants
        .filter(v => 
          Object.entries(updatedSelections)
            .slice(0, i)
            .every(([key, value]) => v.options[key] === value)
        )
        .map(v => v.options[nextVariant.name])
        .filter((value, index, self) => self.indexOf(value) === index);

      if (!availableOptions.includes(updatedSelections[nextVariant.name])) {
        updatedSelections[nextVariant.name] = availableOptions[0];
      }
    }
    
    setSelectedOptions(updatedSelections);
    updateAvailableOptions(product!, updatedSelections);

    // Find and set the new selected variant
    const newVariant = product?.product_variants.find(v => 
      Object.entries(updatedSelections).every(([key, value]) => v.options[key] === value)
    );
    setSelectedVariant(newVariant || null);
  };

  if (error) {
    return <div className="text-center py-10 text-red-500">Error: {error.message}</div>;
  }

  if (!product) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="flex flex-wrap -mx-4">
          <div className="w-full md:w-1/2 px-4 mb-8">
            <Skeleton className="w-full h-[300px] rounded-lg" />
            <div className="flex justify-center mt-4">
              {[...Array(4)].map((_, i) => (
                <Skeleton key={i} className="w-16 h-16 mx-1 rounded-md" />
              ))}
            </div>
          </div>
          <div className="w-full md:w-1/2 px-4">
            <Skeleton className="h-8 w-3/4 mb-4" />
            <Skeleton className="h-4 w-full mb-6" />
            <Skeleton className="h-6 w-1/4 mb-4" />
            <Skeleton className="h-10 w-1/2 mb-6" />
            {[...Array(3)].map((_, i) => (
              <div key={i} className="mb-4">
                <Skeleton className="h-6 w-1/4 mb-2" />
                <div className="flex flex-wrap gap-2">
                  {[...Array(4)].map((_, j) => (
                    <Skeleton key={j} className="h-10 w-20" />
                  ))}
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex flex-wrap -mx-4">
        <div className="w-full md:w-1/2 px-4 mb-8">
          <Carousel className="w-full max-w-xs mx-auto" setApi={setApi}>
            <CarouselContent>
              {selectedVariant?.images && selectedVariant.images.length > 0 ? (
                selectedVariant.images.map((image, index) => (
                  <CarouselItem key={index}>
                    <div className="w-full h-[300px] relative">
                      <img
                        src={`https://ucarecdn.com/${image}/-/scale_crop/800x600/center/`}
                        alt={`${product.product_name} - Image ${index + 1}`}
                        className="rounded-lg object-cover"
                      />
                    </div>
                  </CarouselItem>
                ))
              ) : (
                <CarouselItem>
                  <div className="w-full h-[300px] relative">
                    <img
                      src="/placeholder.jpg"
                      alt="Placeholder"
                      className="rounded-lg w-full h-full object-cover"
                    />
                  </div>
                </CarouselItem>
              )}
            </CarouselContent>
            <CarouselPrevious />
            <CarouselNext />
          </Carousel>
          <div className="flex justify-center mt-4 overflow-x-auto">
            {selectedVariant?.images && selectedVariant.images.length > 0 ? (
              selectedVariant.images.map((image, index) => (
                <div key={index} className="w-16 h-16 relative mx-1">
                  <img
                    src={`https://ucarecdn.com/${image}/-/scale_crop/100x100/center/`}
                    alt={`Thumbnail ${index + 1}`}
                    className="rounded-md cursor-pointer w-16 h-16 object-cover"
                    onClick={() => api?.scrollTo(index)}
                  />
                </div>
              ))
            ) : (
              <div className="w-16 h-16 relative mx-1">
                <img
                  src="/placeholder.jpg"
                  alt="Placeholder Thumbnail"
                  className="rounded-md cursor-pointer w-16 h-16 object-cover"
                  onClick={() => api?.scrollTo(0)}
                />
              </div>
            )}
          </div>
        </div>
        <div className="w-full md:w-1/2 px-4">
          <h1 className="text-3xl font-bold mb-4">{product.product_name}</h1>
          <p className="text-gray-600 mb-6">{product.product_description || "No description available."}</p>
          
          <div className="flex items-center mb-4">
            <Badge variant={product.is_featured ? "default" : "secondary"}>
              {product.is_featured ? "Featured" : "Not Featured"}
            </Badge>
            <Badge variant={product.is_archived ? "destructive" : "outline"} className="ml-2">
              {product.is_archived ? "Archived" : "Active"}
            </Badge>
          </div>
          
          <div className="mb-6">
            <span className="text-lg text-gray-500">Cost: ${selectedVariant?.cost_price.toFixed(2)}</span>
            <br />
            {selectedVariant?.discounted_price && selectedVariant.discounted_price < selectedVariant.price && (
              <span className="text-xl text-gray-500 line-through mr-2">${selectedVariant.discounted_price.toFixed(2)}</span>
            )}
            <span className="text-3xl font-bold">${selectedVariant?.price.toFixed(2)}</span>
          </div>

          {product.variants.map((variant, variantIndex) => (
            <div key={variant.variant_id} className="mb-4">
              <h3 className="font-medium mb-2">{variant.name}</h3>
              <div className="flex flex-wrap gap-2">
                {variant.options.map((option) => {
                  const isAvailable = availableOptions[variant.name]?.includes(option.value);
                  return (
                    <Button
                      key={option.id}
                      variant={selectedOptions[variant.name] === option.value ? "default" : "outline"}
                      onClick={() => handleOptionSelect(variant.name, option.value)}
                      disabled={!isAvailable}
                    >
                      {option.value}
                    </Button>
                  );
                })}
              </div>
            </div>
          ))}
          
          <Separator className="my-4" />
          
          <div className="mt-4">
            <p><strong>SKU:</strong> {selectedVariant?.sku}</p>
            <p><strong>Availability:</strong> {selectedVariant?.quantity && selectedVariant.quantity > 0 ? 'In Stock' : 'Out of Stock'}</p>
          </div>
        </div>
      </div>
    </div>
  );
}