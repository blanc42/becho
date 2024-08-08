"use client"

import { useStore } from "@/lib/store/useStore";
import { useUser } from "@/lib/store/useUser";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function MainLayout({
    children,
  }: {
    children: React.ReactNode
  }) {
    const router = useRouter();

    const {user, fetchUser} = useUser();
    // const {stores, setSelectedStore, fetchStores} = useStore();

    useEffect(() => {
        fetchUser();
        // fetchStores();
    }, []);

    if (!user) {
        router.push("/login");
    }

    // if (stores) {
    //     setSelectedStore(stores[0]);
    //     router.push("/dashboard");
    // }

    router.push("/dashboard");

    return <>{children}</>
}