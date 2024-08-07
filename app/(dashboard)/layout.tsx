import Navbar from '@/components/Navbar';
import React from 'react';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex w-full min-h-screen  flex-col">
        <Navbar/>
        <div className='flex w-full flex-col p-4 px-20'>
        {children}
        </div>
    </div>
  );
}
