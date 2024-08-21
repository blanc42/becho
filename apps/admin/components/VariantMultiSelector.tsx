"use client"

import * as React from "react"
import { Check, ChevronsUpDown } from "lucide-react"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import { Variant } from "@/lib/types"

const VariantMultiSelector = ({
  variants,
  selectedVariants,
  onVariantSelect,
  categoryVariants,
}: {
  variants: Variant[],
  selectedVariants: Variant[],
  onVariantSelect: (variantId: string) => void,
  categoryVariants: Variant[]
}) => {
  const [open, setOpen] = React.useState(false)

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-[200px] justify-between"
        >
          Select variants...
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[200px] p-0">
        <Command>
          <CommandInput placeholder="Search variants..." />
          <CommandList>
            <CommandEmpty>No variants found.</CommandEmpty>
            <CommandGroup>
              {variants.map((variant) => (
                <CommandItem
                  key={variant.id}
                  value={variant.id}
                  onSelect={() => {
                    onVariantSelect(variant.id);
                    setOpen(false);
                  }}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      selectedVariants.some((v) => v.id === variant.id)
                        ? "opacity-100"
                        : "opacity-0"
                    )}
                  />
                  {variant.name}
                  {categoryVariants.some(cv => cv.id === variant.id) && (
                    <span className="ml-2 text-sm text-gray-500">(Category Variant)</span>
                  )}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}

export default VariantMultiSelector