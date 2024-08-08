"use client"

import * as React from "react"
import { cn } from "@/lib/utils"
import { X } from "lucide-react"

interface TagProps extends React.ComponentPropsWithoutRef<"span"> {
  variant?: "default" | "gray" | "red" | "green" | "blue" | "indigo" | "purple" | "yellow" | "orange" | "cyan" | "teal" | "pink"
  size?: "sm" | "md" | "lg"
  className?: string
  children: React.ReactNode
  onClose?: () => void
}

const Tag = React.forwardRef<HTMLSpanElement, TagProps>(
  ({ variant = "default", size = "md", className, children, onClose, ...props }, ref) => {
    return (
      <span
        ref={ref}
        className={cn(
          "rounded-md inline-flex items-center max-w-full",
          size === "sm" ? "px-1.5 py-0.5 text-xs" : 
          size === "md" ? "px-2 py-1 text-sm" : 
          "px-3 py-1.5 text-base",
          "font-medium",
          variant === "default"
            ? "bg-gray-100 text-gray-700"
            : `bg-${variant}-100 text-${variant}-700`,
          "max-h-8",
          className
        )}
        {...props}
      >
        {children}
        {onClose && (
          <TagCloseButton onClick={onClose} className="ml-1" size={size} />
        )}
      </span>
    )
  }
)

Tag.displayName = "Tag"

export { Tag }

interface TagCloseButtonProps extends React.ComponentPropsWithoutRef<"button"> {
  className?: string
  size?: "sm" | "md" | "lg"
}

const TagCloseButton = React.forwardRef<HTMLButtonElement, TagCloseButtonProps>(
  ({ className, size = "md", ...props }, ref) => {
    return (
      <button
        ref={ref}
        className={cn(
          "rounded-sm p-0.5 hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-gray-300",
          size === "sm" ? "h-4 w-4" : 
          size === "md" ? "h-5 w-5" : 
          "h-6 w-6",
          className
        )}
        {...props}
      >
        <X className={cn(
          size === "sm" ? "h-2 w-2" : 
          size === "md" ? "h-3 w-3" : 
          "h-4 w-4"
        )} />
      </button>
    )
  }
)

TagCloseButton.displayName = "TagCloseButton"

export { TagCloseButton }

interface TagLabelProps extends React.ComponentPropsWithoutRef<"span"> {
  className?: string
}

const TagLabel = React.forwardRef<HTMLSpanElement, TagLabelProps>(
  ({ className, ...props }, ref) => {
    return (
      <span
        ref={ref}
        className={cn(
          "flex items-center gap-1 truncate",
          className
        )}
        {...props}
      />
    )
  }
)

TagLabel.displayName = "TagLabel"

export { TagLabel }