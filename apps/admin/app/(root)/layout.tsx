"use client"
import CreateStoreModal from "@/components/modals/CreateStoreModal";
import LoadingPage from "@/components/pages/LoadingPage";
import { User } from "@/lib/schema";
import { useStoreModal } from "@/lib/store/useStoreModal";
import { useUser } from "@/lib/store/useUser";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function MainLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const router = useRouter()
  const { user, fetchUser } = useUser()
  const { isOpen, onOpen, onClose } = useStoreModal()
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<Error | null>(null)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const fetchedUser = await fetchUser('/api/v1/user')
        if (!fetchedUser) {
          console.log("redirecting to login since user is not logged in")
          router.push("/login")
        } else if (fetchedUser.stores?.length === 0) {
          console.log("opening modal since user has no stores")
          router.push("/")
          onOpen()
        }
      } catch (err) {
        if (err instanceof Error) {
          setError(err)
          if (err.message === 'Unauthorized') {
            console.log('Unauthorized, redirecting to login...')
            router.push("/login")
          }
        }
      } finally {
        setIsLoading(false)
      }
    }

    fetchData()
  }, [])



  if (isLoading) {
    return <LoadingPage />
  }
  
  return <>
    <CreateStoreModal isOpen={isOpen} onClose={onClose} />
      {children}
  </>
}