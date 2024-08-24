import React, { useState, useCallback, useEffect } from 'react';
import { useDropzone } from 'react-dropzone';
import { X } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';

interface UploadImageProps {
  images: string[];
  setImages: React.Dispatch<React.SetStateAction<string[]>>;
  maxImages: number;
  isOpen: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const UploadImage: React.FC<UploadImageProps> = ({ images, setImages, maxImages, isOpen, setIsOpen }) => {
  const [isUploading, setIsUploading] = useState(false);
  const [selectedFiles, setSelectedFiles] = useState<File[]>([]);

  const onDrop = useCallback((acceptedFiles: File[]) => {
    setSelectedFiles(prev => [...prev, ...acceptedFiles].slice(0, maxImages - images.length));
  }, [images, maxImages]);

  const { getRootProps, getInputProps, isDragActive } = useDropzone({ 
    onDrop,
    disabled: images.length >= maxImages || isUploading
  });

  const handleRemoveImage = (index: number) => {
    setImages(prev => prev.filter((_, i) => i !== index));
  };

  const handleRemoveSelectedFile = (index: number) => {
    setSelectedFiles(prev => prev.filter((_, i) => i !== index));
  };

  const uploadFiles = async () => {
    setIsUploading(true);
    for (const file of selectedFiles) {
      try {
        // Get signed upload URL from your backend
        const response = await fetch('/api/get-upload-url', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ filename: file.name, contentType: file.type }),
        });
        const { url, fields } = await response.json();

        // Prepare form data for upload
        const formData = new FormData();
        Object.entries(fields).forEach(([key, value]) => {
          formData.append(key, value as string);
        });
        formData.append('file', file);

        // Upload to Uploadcare
        const uploadResponse = await fetch(url, {
          method: 'POST',
          body: formData,
        });

        if (uploadResponse.ok) {
          const uploadedFileUrl = uploadResponse.url;
          const fileId = uploadedFileUrl.split('/').pop();
          setImages(prev => [...prev, fileId as string]);
        } else {
          console.error('Upload failed');
        }
      } catch (error) {
        console.error('Error during upload:', error);
      }
    }
    setSelectedFiles([]);
    setIsUploading(false);
  };

  useEffect(() => {
    if (selectedFiles.length > 0 && !isUploading) {
      uploadFiles();
    }
  }, [selectedFiles]);

  return (
    <>
      <Button onClick={() => setIsOpen(true)}>Upload Images</Button>
      <Dialog open={isOpen} onOpenChange={setIsOpen}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Upload Images</DialogTitle>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div {...getRootProps()} className={`border-2 border-dashed rounded-md p-4 ${isDragActive ? 'border-primary' : 'border-gray-300'} ${images.length >= maxImages || isUploading ? 'opacity-50 cursor-not-allowed' : ''}`}>
              <input {...getInputProps()} />
              {isDragActive ? (
                <p>Drop the files here ...</p>
              ) : (
                <p>Drag 'n' drop some files here, or click to select files</p>
              )}
            </div>
            <div className="grid grid-cols-3 gap-2">
              {images.map((image, index) => (
                <div key={index} className="relative">
                  <img src={`https://ucarecdn.com/${image}/`} alt={`Uploaded ${index}`} className="w-full h-24 object-cover rounded" />
                  <button
                    onClick={() => handleRemoveImage(index)}
                    className="absolute top-0 right-0 bg-red-500 text-white rounded-full p-1"
                  >
                    <X size={16} />
                  </button>
                </div>
              ))}
              {selectedFiles.map((file, index) => (
                <div key={index} className="relative">
                  <img src={URL.createObjectURL(file)} alt={`Selected ${index}`} className="w-full h-24 object-cover rounded" />
                  <button
                    onClick={() => handleRemoveSelectedFile(index)}
                    className="absolute top-0 right-0 bg-red-500 text-white rounded-full p-1"
                  >
                    <X size={16} />
                  </button>
                </div>
              ))}
            </div>
          </div>
          <Button onClick={() => setIsOpen(false)} disabled={isUploading}>
            Done
          </Button>
        </DialogContent>
      </Dialog>
    </>
  );
};

export default UploadImage;

