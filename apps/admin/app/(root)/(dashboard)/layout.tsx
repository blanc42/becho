"use client"
import Navbar from '@/components/Navbar';
import SideBar from '@/components/SideBar';
import { useStoreModal } from '@/lib/store/useStoreModal';
import { useUser } from '@/lib/store/useUser';
import { useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {

  const { isOpen, onClose, onOpen } = useStoreModal()
  const { user } = useUser()
  const router = useRouter()
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

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
    <div className="flex min-h-screen">
      <SideBar isOpen={isSidebarOpen} setIsOpen={setIsSidebarOpen} />
      <div className={`flex flex-1 transition-all duration-300 ease-in-out w-[calc(100%-256px)] px-2 py-4 lg:px-10 ${isSidebarOpen ? 'md:ml-64' : 'md:ml-16'}`}>
        {children}
      </div>
    </div>
  );
}
