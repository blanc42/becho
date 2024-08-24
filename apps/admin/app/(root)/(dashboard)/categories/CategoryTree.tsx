"use client"
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { PlusSquareIcon, MinusSquareIcon, ChevronRightIcon, ChevronDownIcon, Edit, Trash2 } from "lucide-react";
import { toast } from "@/components/ui/use-toast";

interface Category {
  id: string;
  name: string;
  description: string;
  level: number;
  parent_id: string | null;
  unique_identifier: string;
  variants: string[];
}

export default function CategoryTree({ categories, parentId = null }: { categories: Category[], parentId?: string | null }) {
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

  const filteredCategories = categories.filter(cat => cat.parent_id === parentId);

  return (
    <ul className="pl-4">
      {filteredCategories.map(category => (
        <li key={category.id} className="mb-2 hover:bg-gray-800">
          <div className="flex items-center justify-between group">
            <div className="flex items-center">
              <Button
                variant="ghost"
                size="icon"
                className="mr-1"
                disabled={!categories.some(cat => cat.parent_id === category.id)}
                onClick={() => toggleExpand(category.id)}
              >
                {expanded[category.id] ? <MinusSquareIcon size={16} /> : <PlusSquareIcon size={16} />}
              </Button>
              <span>{category.name}</span>
            </div>
            <div className="flex items-center opacity-0 group-hover:opacity-100 transition-opacity">
              <Button variant="ghost" size="icon" onClick={() => handleEdit(category.id)}>
                <Edit size={16} />
              </Button>
              <Button variant="ghost" size="icon" onClick={() => handleDelete(category.id)}>
                <Trash2 size={16} />
              </Button>
            </div>
          </div>
          {expanded[category.id] && (
            <CategoryTree categories={categories} parentId={category.id} />
          )}
        </li>
      ))}
    </ul>
  );
};