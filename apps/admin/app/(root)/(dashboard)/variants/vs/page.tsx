"use client"
import { useState } from "react";
import { Tag, TagLabel, TagCloseButton } from "@/components/ui/tag";
import VariantSelect from "./SelectVariants";
import { Variant } from "@/lib/types";

export default function SelectVariant() {
    const [selectedVariants, setSelectedVariants] = useState<Variant[]>([]);

    const handleDeleteVariant = (variantId: string) => {
      setSelectedVariants(selectedVariants.filter((v) => v.id !== variantId));``
    };
  
    return (
      <div>
        <h2>Variants</h2>
        <div className="flex flex-wrap gap-2 mb-4 border border-gray-200 rounded-md p-2 max-w-[500px] min-h-24">
          {selectedVariants.map((variant) => (
            <Tag key={variant.id} variant="default" size="md" className="cursor-pointer">
              <TagLabel>{variant.name}</TagLabel>
              <TagCloseButton onClick={() => handleDeleteVariant(variant.id)} />
            </Tag>
          ))}
        </div>
        <VariantSelect
          selectedVariants={selectedVariants}
          setSelectedVariants={setSelectedVariants}
        />
      </div>
    );
}