"use client"

import * as React from "react"
import { Check, ChevronsUpDown, PlusCircle } from "lucide-react"
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
import { useStoreData } from '@/lib/store/useStoreData'


const SelectVariants = ({
  selectedVariants,
  setSelectedVariants,
}: {
  selectedVariants: Variant[]
  setSelectedVariants: React.Dispatch<React.SetStateAction<Variant[]>>
}) => {
  const [open, setOpen] = React.useState(false)
  const [variantsData, setVariantsData] = React.useState<Variant[]>([]);
  const { selectedStore } = useStoreData();

  React.useEffect(() => {
    const fetchVariants = async () => {
      if (selectedStore) {
        try {
          const response = await fetch(`/api/v1/stores/${selectedStore.id}/variants`);
          if (!response.ok) {
            throw new Error('Failed to fetch variants');
          }
          const data = await response.json();
          setVariantsData(data);
        } catch (error) {
          console.error('Error fetching variants:', error);
        }
      }
    };

    fetchVariants();
  }, [selectedStore]);

  const handleVariantSelect = (variantId: string) => {
    const variant = variantsData.find((v: Variant) => v.id === variantId)
    if (variant) {
      if (selectedVariants.some((v) => v.id === variantId)) {
        setSelectedVariants(
          selectedVariants.filter((v) => v.id !== variantId)
        )
      } else {
        setSelectedVariants([...selectedVariants, variant])
      }
    }
    setOpen(false)
  }

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
        <Command value={selectedVariants.slice(0, 2).map((v) => v.id).join(",")}>
          <CommandInput placeholder="Search variants..." />
          <CommandList>
            <CommandEmpty>No variants found.</CommandEmpty>
            <CommandGroup>
              {variantsData.map((variant: Variant) => (
                <CommandItem
                  key={variant.id}
                  value={variant.id}
                  onSelect={() => handleVariantSelect(variant.id)}
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
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}

export default SelectVariants