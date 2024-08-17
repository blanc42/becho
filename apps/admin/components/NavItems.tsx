import React from 'react'
import Link from 'next/link'
import { Button } from './ui/button'

const NavItems = () => {
  return (
    <div className="flex flex-col gap-2 md:flex-row md:space-x-0 w-full">
      <Link href="/dashboard" ><Button variant="link">Dashboard</Button></Link>
      <Link href="/products" ><Button variant="link">Products</Button></Link>
      <Link href="/categories" ><Button variant="link">Categories</Button></Link>
      <Link href="/variants" ><Button variant="link">Variants</Button></Link>
    </div>
  )
}

export default NavItems