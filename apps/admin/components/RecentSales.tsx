import React from 'react';
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"

export function RecentSales() {
  return (
    <div className="space-y-8 ">
      <div className="flex items-center py-4">
        <Avatar className="h-9 w-9">
          <AvatarImage src="/avatars/01.png" alt="Avatar" />
          <AvatarFallback>OM</AvatarFallback>
        </Avatar>
        <div className="ml-4 space-y-1">
          <p className="text-sm font-medium leading-none">Olivia Martin</p>
          <p className="text-sm text-muted-foreground">
            olivia.martin@email.com
          </p>
        </div>
        <div className="ml-auto font-medium">+$1,999.00</div>
      </div>
      {/* Add more recent sales items */}
    </div>
  )
}