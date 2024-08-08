"use client"
import { useState } from "react";
import VariantSelect from "./SelectVariants";
import { Tag, TagLabel, TagCloseButton } from "@/components/ui/tag"

interface Variant {
  id: string
  name: string
}

export default function VariantsPage() {
  const [selectedVariants, setSelectedVariants] = useState<Variant[]>([{id: "1", name: "size"}, {id: "2", name: "color"}]);

  const handleDeleteVariant = (variantId: string) => {
    setSelectedVariants(selectedVariants.filter((v) => v.id !== variantId));
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