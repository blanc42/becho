import {z} from "zod"

interface Variant {
    id: string
    name: string
    label : string
    options: VariantOption[]
}

interface VariantOption {
    id: string
    value: string
    image_url: string
    data: string
  }

interface Category {
  id: string
  name: string
  description: string | null
  parent_id: string | null
  variants: string[]
  level: number
  subCategories?: Category[]
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
  description: z.string(),
  // category_id: z.string().min(1, "Category is required"),
  has_variants: z.boolean(),
  is_archived: z.boolean(),
  // variants: z.array(z.string()),
  is_featured: z.boolean(),
  // items: z.array(ItemSchema),
});

type ProductFormValues = z.infer<typeof ProductSchema>


export type { Variant, Category, ProductFormValues }
export { ProductSchema }