"use client"
import Navbar from '@/components/Navbar';
import { useStoreModal } from '@/lib/store/useStoreModal';
import { useUser } from '@/lib/store/useUser';
import { useRouter } from 'next/navigation';
import React, { useEffect } from 'react';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {

  const { isOpen, onClose, onOpen } = useStoreModal()
  const { user } = useUser()
  const router = useRouter()

  useEffect(() => {
    if (user && user?.stores?.length === 0) {
      onOpen();
      console.log(user)
    } else if (user && user?.stores?.length > 0) {
      console.log("user has stores")
    }
  }, [user]);

  if (isOpen) {
    return null;
  }
  return (
    <div className="flex w-full min-h-screen flex-col">
      <Navbar />
      <div className='flex w-full p-4 px-2 md:px-20'>
        {children}
      </div>
    </div>
  );
}
