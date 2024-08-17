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
import { useRouter } from "next/navigation";
import { useStore } from '@/lib/store/useStore';

const variantSchema = z.object({
  name: z.string().min(3, 'Variant name must be at least 3 characters'),
  description: z.string().optional(),
  options: z.array(
    z.object({
      value: z.string().min(3, 'Variant option must be at least 3 characters'),
      data: z.string().optional(),
      imageId: z.string().optional().nullable(),
    })
  ),
});

type VariantFormValues = z.infer<typeof variantSchema>;

export default function AddVariant() {
  const { selectedStore } = useStore()
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();
  const form = useForm<VariantFormValues>({
    resolver: zodResolver(variantSchema),
    defaultValues: {
      name: '',
      description: '',
      options: [],
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
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="Size" {...field} />
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
                <Input placeholder="variant description" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="space-y-4">
          <FormLabel>Options</FormLabel>
          <div className="space-y-2">
            {fields.map((field, index) => (
              <div key={field.id} className="flex items-center space-x-2 p-2 border rounded">
                <FormField
                  control={form.control}
                  name={`options.${index}.value`}
                  render={({ field }) => (
                    <FormItem className="flex-grow">
                      <FormControl>
                        <Input placeholder="Option Value" {...field} />
                      </FormControl>
                    </FormItem>
                  )}
                />
                <Button type="button" variant="destructive" onClick={() => remove(index)}>
                  Remove
                </Button>
              </div>
            ))}
          </div>
          <Button 
            type="button" 
            onClick={() => append({ value: '', data: '', imageId: null })} 
            className="w-full"
          >
            Add Option
          </Button>
        </div>
        <Button type="submit" disabled={isLoading}>
          {isLoading ? 'Creating...' : 'Create Variant'}
        </Button>
      </form>
    </Form>
  );
}