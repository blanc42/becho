"use client"

import React, { useState } from 'react'
import { Button } from "@/components/ui/button"
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"
import { Menu } from 'lucide-react'
import StoreSelector from './StoreSelector'
import NavItems from './NavItems'

const Navbar = () => {
  return (
    <nav className="flex w-full items-center justify-between p-4 bg-background border-b border-neutral-200">
      <div className="text-xl font-bold mr-2">Logo</div>

      <div className="hidden md:block mr-auto">
        <StoreSelector />
      </div>

      <div className="hidden md:block">
        <NavItems />
      </div>

      <Sheet>
        <SheetTrigger asChild className="md:hidden">
          <Button variant="outline" size="icon">
            <Menu className="h-6 w-6" />
          </Button>
        </SheetTrigger>
        <SheetContent side="right" className='w-[300px]'>
          <div className="flex flex-col space-y-4 mt-4 items-center w-full">
            <StoreSelector />
            <NavItems />
          </div>
        </SheetContent>
      </Sheet>
    </nav>
  )
}

export default Navbar