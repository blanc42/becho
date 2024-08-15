'use client'
import LoadingPage from '@/components/LoadingPage';
import { useUser } from '@/lib/store/useUser';
import { useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter()
  const {user} = useUser()
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    console.log("I am running in auth layout ")
    if (user) {
      router.push("/dashboard")
    }else {
      setLoading(false)
    }
  }, [user])

  if (loading) {
    return <LoadingPage />
  }

  return (
    <div className="flex min-h-screen  flex-col items-center justify-center">
      <div className="w-full max-w-md border border-gray-200 h-full space-y-8 rounded-lg p-6 shadow-md">
        {children}
      </div>
    </div>
  );
}
