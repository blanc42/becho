import { useState, useEffect } from 'react'
import Link from 'next/link'
import { Button } from "@/components/ui/button"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { Menu, Home, Settings, HelpCircle, LogOut, ChevronRight, Mountain, X, ChevronLeft, ShoppingBag, Plus, User, Package, Layers, Grid } from 'lucide-react'
import StoreSelector from './StoreSelector'

export default function Component({ isOpen, setIsOpen }: { isOpen: boolean; setIsOpen: (isOpen: boolean) => void }) {
  const [showMobileNav, setShowMobileNav] = useState(true)
  const [lastScrollY, setLastScrollY] = useState(0)
  const [expandedItems, setExpandedItems] = useState<string[]>([])

  useEffect(() => {
    const handleScroll = () => {
      const currentScrollY = window.scrollY
      if (currentScrollY > lastScrollY) {
        setShowMobileNav(false)
      } else {
        setShowMobileNav(true)
      }
      setLastScrollY(currentScrollY)
    }

    window.addEventListener('scroll', handleScroll, { passive: true })
    return () => window.removeEventListener('scroll', handleScroll)
  }, [lastScrollY])

  const toggleExpand = (label: string) => {
    setExpandedItems(prev =>
      prev.includes(label) ? prev.filter(item => item !== label) : [...prev, label]
    )
  }

  const MenuItem = ({ icon: Icon, label, href, subItems }: { icon: React.ElementType; label: string; href?: string; subItems?: Array<{ label: string; href: string; icon: React.ElementType }> }) => {
    const iconElement = <Icon className="h-4 w-4" />
    const isExpanded = expandedItems.includes(label)

    if (!isOpen) {
      if (!subItems) {
        return (
          <Link href={href || '#'}>
            <Button variant="ghost" className="w-full justify-center p-2">
              {iconElement}
            </Button>
          </Link>
        )
      }

      return (
        <Popover>
          <PopoverTrigger asChild>
            <Button variant="ghost" className="w-full justify-center p-2">
              {iconElement}
            </Button>
          </PopoverTrigger>
          <PopoverContent side="right" align="start" className="w-48 p-0">
            {subItems.map((item, index) => (
              <Link key={index} href={item.href} className="flex items-center p-2 hover:bg-accent">
                <item.icon className="h-4 w-4 mr-2" />
                {item.label}
              </Link>
            ))}
          </PopoverContent>
        </Popover>
      )
    }

    const content = (
      <Button variant="ghost" className="w-full justify-start" onClick={() => subItems && toggleExpand(label)}>
        {iconElement}
        <span className="ml-2">{label}</span>
        {subItems && (
          <ChevronRight className={`ml-auto h-4 w-4 transition-transform ${isExpanded ? 'rotate-90' : ''}`} />
        )}
      </Button>
    )

    if (subItems) {
      return (
        <div>
          {content}
          {isExpanded && (
            <div className="ml-4 space-y-1">
              {subItems.map((item, index) => (
                <Link key={index} href={item.href}>
                  <Button variant="ghost" className="w-full justify-start">
                    <item.icon className="h-4 w-4 mr-2" />
                    {item.label}
                  </Button>
                </Link>
              ))}
            </div>
          )}
        </div>
      )
    }

    return <Link href={href || '#'}>{content}</Link>
  }

  return (
    <>
      <aside
        className={`fixed left-0 top-0 z-50 h-screen bg-background transition-all duration-300 ease-in-out 
                    ${isOpen ? 'w-64' : 'w-16'} 
                    bg-gradient-to-r from-background/10 to-background/10 backdrop-blur-sm
                    md:block ${isOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0'}`}
      >
        <div className="flex h-full flex-col">
          <div className="flex items-center justify-between border-b p-4">
            <Link href="/" className="flex items-center space-x-2">
              <Mountain className="h-6 w-6" />
              {isOpen && <span className="font-bold">Brand</span>}
            </Link>
          </div>
          
          {/* Add StoreSelector here */}
          {isOpen && (
            <div className="px-4 py-2">
              <StoreSelector />
            </div>
          )}

          <nav className="flex-1 space-y-2 p-2">
            <MenuItem icon={Home} label="Dashboard" href="/dashboard" />
            <MenuItem
              icon={Package}
              label="Inventory"
              subItems={[
                { label: 'Products', href: '/products', icon: Package },
                { label: 'Categories', href: '/categories', icon: Grid },
                { label: 'Variants', href: '/variants', icon: Layers },
              ]}
            />
            <MenuItem
              icon={Plus}
              label="Create"
              subItems={[
                { label: 'Add Product', href: '/products/add', icon: Package },
                { label: 'Add Category', href: '/categories/add', icon: Grid },
                { label: 'Add Variant', href: '/variants/add', icon: Layers },
              ]}
            />
            <MenuItem icon={ShoppingBag} label="Orders" href="/orders" />
            <MenuItem icon={Settings} label="Settings" href="/settings" />
            <MenuItem icon={HelpCircle} label="Help" href="/help" />
          </nav>
          <div className='flex justify-center'>
                <Button
                variant="ghost"
                size="icon"
                onClick={() => setIsOpen(!isOpen)}
                className={`w-12 h-12 mr-2 ${isOpen ? 'mb-2 ml-auto' : 'mb-2 ml-auto'}`}
                >
                {isOpen ? 
                <div className='flex'>
                    <ChevronLeft className="h-4 w-4" />
                    <ChevronLeft className="h-4 -ml-2 w-4" />
                    <ChevronLeft className="h-4 -ml-2 w-4" />
                </div>
                
                : 
                <div className='flex'>
                    <ChevronRight className="h-4 -mr-2 w-4" />
                    <ChevronRight className="h-4 -mr-2 w-4" />
                    <ChevronRight className="h-4 w-4" />
                </div>
                }
                </Button>
                </div>
          <div className="w-full border-t mx-auto py-4">
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" className={`w-full h-10 flex items-center justify-start ${isOpen ? 'px-4' : 'justify-center'}`}>
                  <Avatar className="h-8 w-8 rounded-lg">
                    <AvatarImage src="/placeholder.jpg" alt="User" />
                    <AvatarFallback>U</AvatarFallback>
                  </Avatar>
                  {isOpen && <span className="ml-2 font-medium">User</span>}
                  {isOpen && <ChevronRight className="ml-auto h-4 w-4" />}
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent side={isOpen ? "right" : "right"} align="start" className="w-56">
                <DropdownMenuItem asChild>
                  <Link href="/profile" className="flex items-center">
                    <User className="mr-2 h-4 w-4" />
                    <span>Profile</span>
                  </Link>
                </DropdownMenuItem>
                <DropdownMenuItem asChild>
                  <Link href="/settings" className="flex items-center">
                    <Settings className="mr-2 h-4 w-4" />
                    <span>Settings</span>
                  </Link>
                </DropdownMenuItem>
                <DropdownMenuItem asChild>
                  <Link href="/logout" className="flex items-center">
                    <LogOut className="mr-2 h-4 w-4" />
                    <span>Logout</span>
                  </Link>
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </div>
      </aside>

      <nav
        className={`fixed bottom-4 left-4 right-4 z-40 rounded-full bg-gradient-to-r from-background/10 to-background/10 backdrop-blur-sm transition-transform duration-300 ease-in-out md:hidden
                    ${showMobileNav ? 'translate-y-0' : 'translate-y-full'}`}
      >
        <div className="flex items-center justify-around p-4">
          <Button variant="ghost" size="icon" onClick={() => setIsOpen(!isOpen)}>
            <Menu className="h-6 w-6" />
          </Button>
          <Link href="/orders">
            <Button variant="link" size="icon">
              <ShoppingBag className="h-6 w-6" />
            </Button>
          </Link>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon">
                <Plus className="h-6 w-6" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent side="top" align="center">
              <DropdownMenuItem asChild>
                <Link href="/products/add" className="w-full">Add Product</Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link href="/categories/add" className="w-full">Add Category</Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link href="/variants/add" className="w-full">Add Variant</Link>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
          <Link href="/settings">
            <Button variant="ghost" size="icon">
              <Settings className="h-6 w-6" />
            </Button>
          </Link>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon">
                <Avatar className="h-8 w-8">
                  <AvatarImage src="/placeholder.jpg" alt="User" />
                  <AvatarFallback>U</AvatarFallback>
                </Avatar>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent side="top" align="end" className="w-48">
              <DropdownMenuItem asChild>
                <Link href="/profile" className="w-full flex items-center">
                  <User className="mr-2 h-4 w-4" />
                  Profile
                </Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link href="/settings" className="w-full flex items-center">
                  <Settings className="mr-2 h-4 w-4" />
                  Settings
                </Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link href="/logout" className="w-full flex items-center">
                  <LogOut className="mr-2 h-4 w-4" />
                  Logout
                </Link>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </nav>
    </>
  )
}