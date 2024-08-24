"use client"

import React, { useState } from 'react'
import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useStoreData } from '@/lib/store/useStoreData'

interface CreateStoreModalProps {
  isOpen: boolean;
  onClose: () => void;
}

const CreateStoreModal: React.FC<CreateStoreModalProps> = ({ isOpen, onClose }) => {
  const [storeName, setStoreName] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const { addStore } = useStoreData()

  const handleCreateStore = async () => {
    if (storeName) {
      setIsLoading(true)
      try {
        const response = await fetch('/api/v1/stores', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
          body: JSON.stringify({ name: storeName }),
        })

        if (!response.ok) {
          throw new Error('Failed to create store')
        }

        const res = await response.json()
        addStore(res.data)
        setStoreName('')
        onClose()
      } catch (error) {
        console.error('Error creating store:', error)
        // Handle error (e.g., show error message to user)
      } finally {
        setIsLoading(false)
      }
    }
  }

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Create New Store</DialogTitle>
          <DialogDescription>
            Enter the name for your new store.
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Name
            </Label>
            <Input
              id="name"
              value={storeName}
              onChange={(e) => setStoreName(e.target.value)}
              className="col-span-3"
            />
          </div>
        </div>
        <DialogFooter>
          <Button type="submit" onClick={handleCreateStore} disabled={isLoading}>
            {isLoading ? 'Creating...' : 'Create Store'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

export default CreateStoreModal