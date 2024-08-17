"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm, useFieldArray } from "react-hook-form"
import { z } from "zod"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Checkbox } from "@/components/ui/checkbox"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Trash2 } from "lucide-react"
import { CategorySelector } from "@/components/CategorySelector"

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
  category_id: z.string().min(1, "Category is required"),
  has_variants: z.boolean(),
  is_archived: z.boolean(),
  variants: z.array(z.string()),
  is_featured: z.boolean(),
  items: z.array(ItemSchema),
});

type ProductFormValues = z.infer<typeof ProductSchema>

export default function AddProductPage() {
  const form = useForm<ProductFormValues>({
    resolver: zodResolver(ProductSchema),
    defaultValues: {
      name: "",
      description: "",
      category_id: "",
      has_variants: false,
      is_archived: false,
      variants: [],
      is_featured: false,
      items: [],
    },
  })

  const { fields, append, remove } = useFieldArray({
    control: form.control,
    name: "items",
  })

  function onSubmit(data: ProductFormValues) {
    console.log(data)
  }

  function handleHasVariantsChange(checked: boolean) {
    if (checked) {
      // Logic for when has_variants is checked
      console.log("Has variants checked")
    } else {
      // Logic for when has_variants is unchecked
      console.log("Has variants unchecked")
    }
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <div className="flex space-x-4">
          <div className="flex-1 space-y-8">
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
          </div>
          <div className="w-64 h-64 bg-gray-200 flex items-center justify-center">
            Image Upload
          </div>
        </div>

        <FormField
          control={form.control}
          name="category_id"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Category</FormLabel>
              <Select onValueChange={field.onChange} defaultValue={field.value}>
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a category" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectItem value="category1">Category 1</SelectItem>
                  <SelectItem value="category2">Category 2</SelectItem>
                  {/* Add more categories as needed */}
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          )}
        />

        <CategorySelector onSelect={(category) => console.log(category)} />

        <div className="flex space-x-4">
          <FormField
            control={form.control}
            name="has_variants"
            render={({ field }) => (
              <FormItem className="flex-1 flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                <FormControl>
                  <Checkbox
                    checked={field.value}
                    onCheckedChange={(checked) => {
                      field.onChange(checked)
                      handleHasVariantsChange(checked as boolean)
                    }}
                  />
                </FormControl>
                <div className="space-y-1 leading-none">
                  <FormLabel>Has Variants</FormLabel>
                  <FormDescription>
                    Check this if the product has variants
                  </FormDescription>
                </div>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="is_archived"
            render={({ field }) => (
              <FormItem className="flex-1 flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                <FormControl>
                  <Checkbox
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                </FormControl>
                <div className="space-y-1 leading-none">
                  <FormLabel>Is Archived</FormLabel>
                  <FormDescription>
                    Check this if the product should be archived
                  </FormDescription>
                </div>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="is_featured"
            render={({ field }) => (
              <FormItem className="flex-1 flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4">
                <FormControl>
                  <Checkbox
                    checked={field.value}
                    onCheckedChange={field.onChange}
                  />
                </FormControl>
                <div className="space-y-1 leading-none">
                  <FormLabel>Is Featured</FormLabel>
                  <FormDescription>
                    Check this if the product should be featured
                  </FormDescription>
                </div>
              </FormItem>
            )}
          />
        </div>

        <div>
          <h2 className="text-lg font-semibold mb-4">Items</h2>
          {fields.map((field, index) => (
            <div key={field.id} className="space-y-4 mb-4 p-4 border rounded relative">
              <Button
                type="button"
                onClick={() => remove(index)}
                className="absolute top-2 right-2 p-2"
                variant="ghost"
              >
                <Trash2 className="h-4 w-4" />
              </Button>
              <FormField
                control={form.control}
                name={`items.${index}.sku`}
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>SKU</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Add other item fields (quantity, cost_price, etc.) here */}
            </div>
          ))}
          <Button
            type="button"
            onClick={() => append({ sku: "", quantity: 0, cost_price: 0, discounted_price: 0, price: 0, variant_options: {} })}
          >
            Add Item
          </Button>
        </div>

        <Button type="submit">Submit</Button>
      </form>
    </Form>
  )
}