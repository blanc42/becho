"use client"
import { useState } from "react";
import UploadImage from "./UploadImage";

export default function ImagesPage() {
 const [images, setImages] = useState<string[]>([]);
 const [isOpen, setIsOpen] = useState(false);
    return (
    <div>
      <UploadImage images={images} setImages={setImages} isOpen={isOpen} setIsOpen={setIsOpen} maxImages={10}/>
    </div>
  );
}