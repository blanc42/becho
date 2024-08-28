"use client"
import { Button } from '@/components/ui/button'
import { ArrowLeftIcon } from 'lucide-react'
import { useRouter } from 'next/navigation'

export default function NotFound() {
    const router = useRouter()
  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h1 className="text-6xl font-bold mb-4">404</h1>
      <h2 className="text-2xl font-semibold mb-6">Page Not Found</h2>
      <p className=" mb-8">Oops! The page you're looking for doesn't exist.</p>
      <Button variant='link' onClick={() => router.back()} className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors">
        <ArrowLeftIcon className="w-4 h-4 mr-2" />
        Go Back
      </Button>
    </div>
  )
}