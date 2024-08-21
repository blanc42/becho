import { z } from "zod";

const createCategorySchema = z.object({
    name: z.string().min(2, 'Category Name is required'),
    description: z.string().optional(),
    parentCategoryId: z.string().nullable(),
    variants: z.array(z.string()),
  });


  const variantSchema = z.object({
    id: z.string(),
    name: z.string(),
    label: z.string(),
    options: z.array(z.string()).min(1, 'At least one option is required'),
  });

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
  });
  
  const createProductSchema = z.object({
    name: z.string().min(1),
    description: z.string().min(1),
    is_featured: z.boolean(),
    is_archived: z.boolean(),
    category_id: z.string(),
    variants: z.array(z.string()),
    items: z.array(itemSchema),
  });




  type createCategoryType = z.infer<typeof createCategorySchema>
  type createProductType = z.infer<typeof createProductSchema>
  type variantType = z.infer<typeof variantSchema>
  type variantOptionType = z.infer<typeof variantOptionSchema>
  type itemType = z.infer<typeof itemSchema>

  export { createCategorySchema, createProductSchema, variantSchema, variantOptionSchema, itemSchema }
  export type { createCategoryType, createProductType, variantType, variantOptionType, itemType }