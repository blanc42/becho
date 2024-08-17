import { Loader2 } from "lucide-react"

export default function LoadingPage() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <div className="p-8 rounded-lg shadow-lg">
        <Loader2 className="w-16 h-16 text-blue-600 animate-spin" />
        <h2 className="mt-4 text-2xl font-bold text-gray-800">Loading...</h2>
        <p className="mt-2 text-gray-600">Please wait while we prepare your content.</p>
      </div>
    </div>
  )
}
