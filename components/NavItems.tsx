import React from 'react'
import Link from 'next/link'

const NavItems = () => {
  return (
    <div className="flex flex-col gap-2 md:flex-row md:space-x-0 w-full">
      <Link href="/dashboard" className="hover:text-primary text-lg text-center bg-neutral-100 rounded-md p-2">Dashboard</Link>
      <Link href="/products" className="hover:text-primary text-lg text-center bg-neutral-100 rounded-md p-2">Products</Link>
      <Link href="/categories" className="hover:text-primary text-lg text-center bg-neutral-100 rounded-md p-2">Categories</Link>
      <Link href="/variants" className="hover:text-primary text-lg text-center bg-neutral-100 rounded-md p-2">Variants</Link>
    </div>
  )
}

export default NavItems