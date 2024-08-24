// components/CategorySingleSelector.tsx
import React, { useState } from 'react';
import { Check, ChevronsUpDown } from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Category } from '@/lib/types';

interface CategorySingleSelectorProps {
  categories: Category[];
  value: string | null;
  onChange: (value: string) => void;
}

export default function CategorySingleSelector({ categories, value, onChange }: CategorySingleSelectorProps) {
  const [open, setOpen] = useState(false);

  const selectedCategory = categories.find(category => category.id === value);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild className='max-w-sm'>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {selectedCategory ? selectedCategory.name : "Select parent category..."}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0 max-w-sm">
        <Command className='w-sm'>
          <CommandInput placeholder="Search category..." />
          <CommandList>

            <CommandEmpty>No category found.</CommandEmpty>
            <CommandGroup>
              {categories.map((category) => (
                <CommandItem
                  key={category.id}
                  onSelect={() => {
                    onChange(category.id);
                    setOpen(false);
                  }}
                  style={{marginLeft: `${category.level * 16}px`}}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      value === category.id ? "opacity-100" : "opacity-0"
                    )}
                  />
                  <span className={`${category.level === 0 ? 'font-semibold' : ''}`}>

                  {category.name}
                  </span>
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}