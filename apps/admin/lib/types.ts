import {z} from "zod"

export interface Variant {
    id: string;
    created_at: string;
    updated_at: string;
    name: string;
    label: string;
    description: string;
    store_id: string;
    options: VariantOption[];
}

export interface VariantOption {
    id: string;
    created_at: string;
    updated_at: string;
    variant_id: string;
    value: string;
    display_order: number;
    data: string;
    image_id: string | null;
}

export interface Category {
    id: string;
    name: string;
    description: string;
    level: number;
    parent_id: string | null;
    unique_identifier: string;
    variants: VariantInsideCategory[];
    subCategories?: Category[];
    image?: string;
}

type VariantInsideCategory = {
    id: string;
    name: string;
}

const VariantOptionSchema = z.record(z.string(), z.string());

const ItemSchema = z.object({
  sku: z.string(),
  quantity: z.number().int().positive(),
  cost_price: z.number().positive(),
  discounted_price: z.number().positive(),
  price: z.number().positive(),
  variant_options: VariantOptionSchema,
});

const ProductSchema = z.object({
  name: z.string().min(1, "Name is required"),
  description: z.string().optional(),
  // category_id: z.string().min(1, "Category is required"),
  has_variants: z.boolean(),
  is_archived: z.boolean(),
  // variants: z.array(z.string()),
  is_featured: z.boolean(),
  // items: z.array(ItemSchema),
});

type ProductFormValues = z.infer<typeof ProductSchema>


export type { ProductFormValues }
export { ProductSchema }