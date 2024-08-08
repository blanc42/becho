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
import { useStore } from '@/lib/store/useStore'
import CreateStoreModal from './CreateStoreModal'

interface Store {
  id: string
  name: string
}

const StoreSelector = () => {
  const [open, setOpen] = React.useState(false)
  const [isCreateModalOpen, setIsCreateModalOpen] = React.useState(false)
  const { stores, selectedStore, setSelectedStore, fetchStores } = useStore()

  React.useEffect(() => {
    const initializeStores = async () => {
      await fetchStores()
    }
    initializeStores()
  }, [])

  React.useEffect(() => {
    console.log('Selected Store Updated:', selectedStore)
  }, [selectedStore])

  const handleStoreSelect = (storeId: string) => {
    console.log('Selecting store with ID:', storeId)
    const store = stores.find(s => s.id === storeId)
    if (store) {
      setSelectedStore(store)
      localStorage.setItem('selectedStoreId', store.id)
    }
    setOpen(false)
  }

  const handleCreateStore = () => {
    setOpen(false)
    setIsCreateModalOpen(true)
  }

  console.log('Rendering with Selected Store:', selectedStore)

  return (
    <>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            className="w-[200px] justify-between"
          >
            {selectedStore ? selectedStore.name : "Select store..."}
            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[200px] p-0">
          <Command value={selectedStore?.id}>
            <CommandInput placeholder="Search store..." />
            <CommandList>
              <CommandEmpty>No store found.</CommandEmpty>
              <CommandGroup>
                {stores.map((store) => (
                  <CommandItem

                    key={store.id}
                    value={store.id}
                    onSelect={() => handleStoreSelect(store.id)}
                  >
                    <Check
                      className={cn(
                        "mr-2 h-4 w-4",
                        selectedStore?.id === store.id ? "opacity-100" : "opacity-0"
                      )}
                    />
                    {store.name}
                  </CommandItem>
                ))}
                <CommandItem onSelect={handleCreateStore}>
                  <PlusCircle className="mr-2 h-4 w-4" />
                  Create new store
                </CommandItem>
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
      <CreateStoreModal 
        isOpen={isCreateModalOpen} 
        onClose={() => setIsCreateModalOpen(false)} 
      />
    </>
  )
}

export default StoreSelector