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
} from "@/components/ui/command"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import { useCategoriesStore } from "@/lib/store/useCategoriesStore"
import { useStore } from "@/lib/store/useStore"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Category } from "@/lib/types"

interface CategoryWithSubCategories extends Category {
  subCategories?: CategoryWithSubCategories[];
}

interface CategorySelectorProps {
  onSelect: (categoryId: string) => void;
}

export function CategorySelector({ onSelect }: CategorySelectorProps) {
  const [open, setOpen] = React.useState(false)
  const [value, setValue] = React.useState("")
  const { categories, fetchCategories, isLoading } = useCategoriesStore()
  const { selectedStore } = useStore()

  React.useEffect(() => {
    if (selectedStore?.id) {
      fetchCategories()
    }
  }, [selectedStore, fetchCategories])
  const renderCategories = (categories: CategoryWithSubCategories[] | undefined, search: string, level = 0): JSX.Element[] => {
    console.log("Rendering categories:", categories, "search:", search, "level:", level);
    
    if (!categories || categories.length === 0) {
      console.log("No categories to render at level:", level);
      return [];
    }
  
    return categories.flatMap((category) => {
      console.log("Processing category:", category);
      
      if (!category) {
        console.log("Encountered undefined category");
        return [];
      }
  
      const matchesSearch = category.name.toLowerCase().includes(search.toLowerCase());
      const childMatches = category.subCategories?.some(child => 
        child && child.name.toLowerCase().includes(search.toLowerCase())
      );
  
      if (!matchesSearch && !childMatches) return [];
  
      const items = [];
  
      if (matchesSearch || (childMatches && search !== "")) {
        items.push(
          <CommandItem
            key={category.id}
            value={category.id}
            onSelect={() => {
              setValue(category.name)
              setOpen(false)
              onSelect(category.id)
            }}
            className={cn("pl-" + (level * 4 + 2))}
          >
            <Check
              className={cn(
                "mr-2 h-4 w-4",
                value === category.id ? "opacity-100" : "opacity-0"
              )}
            />
            {category.name}
          </CommandItem>
        );
      }
  
      if (category.subCategories && category.subCategories.length > 0) {
        console.log("Rendering subcategories for:", category.name);
        items.push(...renderCategories(category.subCategories, search, level + 1));
      }
  
      return items;
    });
  };

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-[300px] justify-between"
        >
          {value || "Select category..."}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[300px] p-0">
        <Command>
          <CommandInput placeholder="Search category..." />
          <CommandEmpty>No category found.</CommandEmpty>
          <CommandGroup>
            <ScrollArea className="h-72">
              {isLoading ? (
                <div className="p-2">Loading categories...</div>
              ) : categories && categories.length > 0 ? (
                renderCategories(categories, "")
              ) : (
                <div className="p-2">No categories available</div>
              )}
            </ScrollArea>
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  )
}