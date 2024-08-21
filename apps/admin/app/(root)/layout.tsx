"use client"

import CreateStoreModal from "@/components/modals/CreateStoreModal";
import LoadingPage from "@/components/pages/LoadingPage";
import { useStore } from "@/lib/store/useStore";
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
  const {user, fetchUser} = useUser()
  const [isLoading, setIsLoading] = useState(true)
  const {isOpen, onOpen} = useStoreModal()
  
  useEffect(() => {
    const userLogic = async () => {
      setIsLoading(true);
      try {
       const fetchedUser = await fetchUser()
        console.log(fetchedUser)
        if (fetchedUser && fetchedUser.stores?.length > 0) {
          console.log("redirecting to dashboard since user has stores")
        } else if (fetchedUser && fetchedUser.stores?.length === 0) {
          console.log("opening modal since user has no stores")
          onOpen()
          setIsLoading(false)
        } else if (!fetchedUser){
          console.log("redirecting to login since user is not logged in")
          setIsLoading(false)
          router.push("/login")
        }
      } catch (error) {
        console.error("Error fetching user:", error);
      } finally {
        setIsLoading(false);
      }
  }
  userLogic()
  }, [])
  
  if (isLoading) {
    return <LoadingPage />
  }

  return <>
  <CreateStoreModal isOpen={isOpen} onClose={()=>router.push("/")} />
  {children}
  </>
}