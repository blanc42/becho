"use client"
import { useState, useEffect } from 'react';
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
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { GripVertical } from 'lucide-react';

const variantSchema = z.object({
  name: z.string().min(3, 'Variant name must be at least 3 characters'),
  description: z.string().optional(),
  options: z.array(
    z.object({
      value: z.string().min(3, 'Variant option must be at least 3 characters'),
      data: z.string().optional(),
      imageId: z.string().optional().nullable(),
      displayOrder: z.number(),
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

  const { fields, append, remove, move } = useFieldArray({
    control: form.control,
    name: 'options',
  });

  const handleOnDragEnd = (result: any) => {
    if (!result.destination) return;

    const { source, destination } = result;
    
    if (destination.index !== source.index) {
      move(source.index, destination.index);
      
      form.setValue('options', 
        form.getValues('options').map((option, index) => ({
          ...option,
          displayOrder: index
        }))
      );
    }
  };

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
          <DragDropContext onDragEnd={handleOnDragEnd}>
            <Droppable droppableId="options">
              {(provided) => (
                <div {...provided.droppableProps} ref={provided.innerRef} className="space-y-2">
                  {fields.map((field, index) => (
                    <Draggable key={field.id} draggableId={field.id} index={index}>
                      {(provided) => (
                        <div
                          ref={provided.innerRef}
                          {...provided.draggableProps}
                          className="flex items-center space-x-2 p-2 border rounded"
                        >
                          <div {...provided.dragHandleProps} className="cursor-move">
                            <GripVertical size={20} />
                          </div>
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
                      )}
                    </Draggable>
                  ))}
                  {provided.placeholder}
                </div>
              )}
            </Droppable>
          </DragDropContext>
          <Button 
            type="button" 
            onClick={() => append({ value: '', data: '', imageId: null, displayOrder: fields.length })} 
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