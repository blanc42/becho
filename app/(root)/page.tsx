"use client"

import { useStoreModal } from "@/lib/store/useStoreModal";
import { useUser } from "@/lib/store/useUser";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

const Home = () => {
  const onOpen = useStoreModal((state) => state.onOpen);
  const {user} = useUser()
  const router = useRouter()

  useEffect(() => {
    if (user && user?.stores?.length === 0) {
      onOpen();
      console.log(user)
    } else if (user && user?.stores?.length > 0) {
      router.push("/dashboard")
    }
  }, [user]);

  return null;
}

export default Home
