"use client"
import React, { useEffect } from 'react'
import { Check, ChevronsUpDown, X } from "lucide-react"
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

interface Category {
    id: string
    label: string
    level: number
    parent_category_id: string | null
}

interface TreeComboboxProps {
    categories: Category[]
}

const TreeCombobox: React.FC<TreeComboboxProps> = ({ categories }) => {
    const [open, setOpen] = React.useState(false)
    const [value, setValue] = React.useState<string>("")

    useEffect(() => {
        console.log("====> categories ", categories)
    }, [categories])

    const handleDeselect = (e: React.MouseEvent) => {
        e.stopPropagation()
        setValue("")
    }

    return (
        <>
            <Popover open={open} onOpenChange={setOpen} >
                <PopoverTrigger asChild>
                    <Button
                        variant="outline"
                        role="combobox"
                        aria-expanded={open}
                        className="w-[200px] justify-between"
                    >
                        {value
                            ? (
                                <div className="flex items-center justify-between w-full">
                                    <span>{categories?.find((category) => category?.id === value)?.label}</span>
                                    <X
                                        className="h-4 w-4 opacity-50 hover:opacity-100"
                                        onClick={handleDeselect}
                                    />
                                </div>
                            )
                            : "Select category..."}
                        <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                    </Button>
                </PopoverTrigger>
                <PopoverContent className="w-[250px] p-0" align='start'>
                    <Command value={value}>
                        <CommandInput placeholder="Search category..." />
                        <CommandList>
                            <CommandEmpty>No category found.</CommandEmpty>
                            <CommandGroup>
                                {categories?.map((category) => (
                                    <CommandItem
                                        key={category.id}
                                        value={category.id}
                                        onSelect={() => setValue(category.id)}
                                        style={{ paddingLeft: `${category.level * 16}px` }}
                                    >
                                        <Check
                                            className={cn(
                                                "mr-2 h-4 w-4",
                                                value === category.id ? "opacity-100" : "opacity-0"
                                            )}
                                        />
                                        {category.label}
                                    </CommandItem>
                                ))}
                            </CommandGroup>
                        </CommandList>
                    </Command>
                </PopoverContent>
            </Popover>
        </>
    )
}

export default TreeCombobox