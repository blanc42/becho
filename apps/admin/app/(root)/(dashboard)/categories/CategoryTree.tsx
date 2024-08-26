"use client"
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { ChevronRightIcon, ChevronDownIcon, Edit, Trash2 } from "lucide-react";
import { toast } from "@/components/ui/use-toast";
import Image from "next/image";
import { Category } from "@/lib/types";
import { Badge } from "@/components/ui/badge";

export default function CategoryTree({ categories }: { categories: Category[] }) {
  const [expanded, setExpanded] = useState<Record<string, boolean>>({});

  const toggleExpand = (id: string) => {
    setExpanded(prev => ({ ...prev, [id]: !prev[id] }));
  };

  const handleEdit = (id: string) => {
    toast({
      title: "Edit Category",
      description: `Editing category with ID: ${id}`,
    });
  };

  const handleDelete = (id: string) => {
    toast({
      title: "Delete Category",
      description: `Deleting category with ID: ${id}`,
      variant: "destructive",
    });
  };

  const renderCategoryRow = (category: Category, depth: number) => {
    const hasChildren = categories.some(cat => cat.parent_id === category.id);
    const isExpanded = expanded[category.id];

    return (
      <div key={category.id}>
        <div 
          className="flex items-center h-16 hover:bg-muted/50 transition-colors"
          style={{ paddingLeft: `${depth * 20}px` }}
        >
          <Button
            variant="ghost"
            size="icon"
            className="mr-2"
            onClick={() => toggleExpand(category.id)}
            disabled={!hasChildren}
          >
            {hasChildren && (isExpanded ? <ChevronDownIcon size={16} /> : <ChevronRightIcon size={16} />)}
          </Button>
          <div className="flex-1">{category.name}</div>
          <div className="flex-1 flex flex-wrap gap-1">
            {category.variants.map(variant => (
              <Badge key={variant.id} variant="secondary">{variant.name}</Badge>
            ))}
          </div>
          <div className="w-24 flex justify-center">
            {category.image ? (
              <Image src={category.image} alt={category.name} width={40} height={40} className="rounded" />
            ) : (
              <div className="w-10 h-10 bg-muted rounded"></div>
            )}
          </div>
          <div className="w-24 flex justify-end">
            <Button variant="ghost" size="icon" onClick={() => handleEdit(category.id)}>
              <Edit size={16} />
            </Button>
            <Button variant="ghost" size="icon" onClick={() => handleDelete(category.id)}>
              <Trash2 size={16} />
            </Button>
          </div>
        </div>
        {isExpanded && renderChildCategories(category.id, depth + 1)}
      </div>
    );
  };

  const renderChildCategories = (parentId: string | null, depth: number) => {
    return categories
      .filter(category => category.parent_id === parentId)
      .map(category => renderCategoryRow(category, depth));
  };

  return (
    <div>
      <div className="flex text-muted-foreground border-b hover:bg-muted/50 font-medium h-12">
        <div className="flex-1 flex items-center px-4">Category Name</div>
        <div className="flex-1 flex items-center px-4">Variants</div>
        <div className="w-24 flex items-center justify-center">Image</div>
        <div className="w-24 flex items-center justify-center">Actions</div>
      </div>
      <div>
        {renderChildCategories(null, 0)}
      </div>
    </div>
  );
}