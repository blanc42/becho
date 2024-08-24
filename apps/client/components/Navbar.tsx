"use client"
import Link from 'next/link';
import { UserAvatar } from './UserAvatar';
import { Category, Store } from '@/lib/types';
import { Search } from 'lucide-react';
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "@/components/ui/navigation-menu"

type NavbarProps = {
    store: Store;
    categories: Category[];
}

export default function Navbar({ store, categories }: NavbarProps) {
  return (
    <nav className="bg-white shadow-md">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center">
            <Link href="/" className="flex-shrink-0 flex items-center mr-10">
              <img className="h-8 w-auto" src={store.logo} alt="Store Logo" />
              <span className="ml-2 text-xl font-bold">{store.name}</span>
            </Link>
            <NavigationMenu>
              <NavigationMenuList>
                {categories.map((category) => (
                  <NavigationMenuItem key={category.id}>
                    <NavigationMenuTrigger>{category.name}</NavigationMenuTrigger>
                    <NavigationMenuContent>
                      <div className="p-6 w-auto">
                        {category.children?.map((subCategory) => (
                          <div key={subCategory.id} className="mb-4">
                            <h3 className="font-semibold text-pink-500 mb-2">
                              <Link href={`/${subCategory.unique_identifier}`} className="hover:underline">
                                {subCategory.name}
                              </Link>
                            </h3>
                            <ul className="space-y-1">
                              {subCategory.children?.map((subSubCategory) => (
                                <li key={subSubCategory.id}>
                                  <NavigationMenuLink asChild>
                                    <Link
                                      href={`/${subSubCategory.unique_identifier}`}
                                      className="text-sm text-gray-600 hover:text-pink-500 transition-colors"
                                    >
                                      {subSubCategory.name}
                                    </Link>
                                  </NavigationMenuLink>
                                </li>
                              ))}
                            </ul>
                          </div>
                        ))}
                      </div>
                    </NavigationMenuContent>
                  </NavigationMenuItem>
                ))}
              </NavigationMenuList>
            </NavigationMenu>
          </div>

          <div className="flex items-center">
            <div className="relative mr-6">
              <input
                type="text"
                placeholder="Search..."
                className="pl-8 pr-4 py-2 border rounded-full focus:outline-none focus:ring-2 focus:ring-pink-300"
              />
              <Search className="absolute left-2 top-1/2 transform -translate-y-1/2 text-gray-400" size={20} />
            </div>
            <Link href="/" className="text-gray-700 hover:text-gray-900 mr-4">Home</Link>
            <UserAvatar />
          </div>
        </div>
      </div>
    </nav>
  );
}