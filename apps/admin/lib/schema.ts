import { z } from "zod";

export const createCategorySchema = z.object({
  name: z.string().min(2, 'Category Name is required'),
  description: z.string().optional(),
  parent_id: z.string().nullable(),
  variants: z.array(z.string()),
  unique_identifier: z.string().min(2, 'Unique identifier is required'),
});

export const variantSchema = z.object({
  id: z.string(),
  created_at: z.string(),
  updated_at: z.string(),
  name: z.string(),
  label: z.string(),
  description: z.string(),
  store_id: z.string(),
  options: z.array(z.object({
      id: z.string(),
      created_at: z.string(),
      updated_at: z.string(),
      variant_id: z.string(),
      value: z.string(),
      display_order: z.number(),
      data: z.string(),
      image_id: z.string().nullable(),
  })),
});

export type CreateCategoryType = z.infer<typeof createCategorySchema>;
export type VariantType = z.infer<typeof variantSchema>;


  const variantOptionSchema = z.object({
    id: z.string(),
    value: z.string(),
    image_url: z.string().optional(),
    data: z.string().optional(),
  });
  
  const itemSchema = z.object({
    sku: z.string(),
    quantity: z.number().int().positive(),
    price: z.number().int().positive(),
    cost_price: z.number().int().positive(),
    discounted_price: z.number().int().positive(),
    variant_options: z.record(z.string(), z.string()),
    images: z.array(z.string()).optional(),
  });
  
  const createProductSchema = z.object({
    name: z.string().min(1),
    description: z.string().optional(),
    is_featured: z.boolean().default(false),
    is_archived: z.boolean().default(false),
    category_id: z.string(),
    variants: z.array(z.string()).optional(),
    items: z.array(itemSchema).min(1),
  });

  type createCategoryType = z.infer<typeof createCategorySchema>
  type createProductType = z.infer<typeof createProductSchema>
  type variantType = z.infer<typeof variantSchema>
  type variantOptionType = z.infer<typeof variantOptionSchema>
  type itemType = z.infer<typeof itemSchema>

  export { createProductSchema, variantOptionSchema, itemSchema }
  export type { createProductType, variantOptionType, itemType }