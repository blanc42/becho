import React, { useState, useCallback, useEffect } from 'react';
import { useDropzone } from 'react-dropzone';
import { Upload, X } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from '@/components/ui/dialog';

interface UploadImageProps {
  images: string[];
  setImages: React.Dispatch<React.SetStateAction<string[]>>;
  maxImages: number;
  variant?: "small" | "default" | "large";
}

const UploadImage: React.FC<UploadImageProps> = ({ images, setImages, maxImages, variant = "default" }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [isUploading, setIsUploading] = useState(false);
  const [signedParams, setSignedParams] = useState<any>(null);

  const onDrop = useCallback(async (acceptedFiles: File[]) => {
    const remainingSlots = maxImages - images.length;
    const filesToUpload = acceptedFiles.slice(0, remainingSlots);
    
    if (filesToUpload.length > 0) {
      await uploadFiles(filesToUpload);
    }
  }, [images, maxImages]);

  const { getRootProps, getInputProps, isDragActive, open } = useDropzone({ 
    onDrop,
    disabled: images.length >= maxImages || isUploading,
    noClick: true
  });

  const handleRemoveImage = (index: number) => {
    setImages(prev => prev.filter((_, i) => i !== index));
  };

  const fetchSignedParams = async () => {
    if (signedParams) return signedParams;
    try {
      const response = await fetch('/api/v1/images/upload');
      if (!response.ok) {
        throw new Error('Failed to fetch signed params');
      }
      const params = await response.json();
      setSignedParams(params);
      return params;
    } catch (error) {
      console.error('Error fetching signed params:', error);
      return null;
    }
  };

  const uploadFiles = async (files: File[]) => {
    setIsUploading(true);
    try {
      let params = await fetchSignedParams();
      if (!params) {
        throw new Error('Failed to get signed params');
      }
      
      for (const file of files) {
        const formData = new FormData();
        formData.append('UPLOADCARE_PUB_KEY', params.public_key);
        formData.append('signature', params.signature);
        formData.append('expire', params.expire);
        formData.append('file', file);

        const response = await fetch('https://upload.uploadcare.com/base/', {
          method: 'POST',
          body: formData,
        });

        if (response.ok) {
          const result = await response.json();
          // setImages(prev => Array.isArray(prev) ? [...prev, result.file] : [result.file]);
          setImages([...images, result.file]);
        } else {
          console.error('Upload failed');
        }
      }
    } catch (error) {
      console.error('Error during upload:', error);
    } finally {
      setIsUploading(false);
    }
  };

  useEffect(() => {
    fetchSignedParams();
  }, []);

  const handleButtonClick = async () => {
    if (maxImages === 1 && variant === "large") {
      await fetchSignedParams();
      open();
    } else {
      setIsOpen(true);
    }
  };

  const renderUploadButton = () => {
    switch (variant) {
      case "small":
        return (
          <Button type='button' onClick={handleButtonClick} variant="outline" size="icon" disabled={isUploading}>
            {isUploading ? <span className="animate-spin">...</span> : <Upload size={16} />}
          </Button>
        );
      case "default":
        return (
          <Button type='button' onClick={handleButtonClick} variant="outline" disabled={isUploading}>
            <Upload size={16} className="mr-2" />
            {isUploading ? 'Uploading...' : 'Upload'}
          </Button>
        );
      case "large":
        if (maxImages === 1 && images.length === 1) {
          return (
            <div className="relative">
              <img 
                src={`https://ucarecdn.com/${images[0]}/`} 
                alt="Uploaded" 
                className="w-full aspect-square object-cover rounded"
              />
              <button
                onClick={() => handleRemoveImage(0)}
                className="absolute top-2 right-2 bg-red-500 text-white rounded-full p-1"
              >
                <X size={16} />
              </button>
            </div>
          );
        }
        return (
          <div 
            {...getRootProps()} 
            className={`border-2 border-dashed rounded-md p-4 aspect-square flex flex-col items-center justify-center cursor-pointer ${isDragActive ? 'border-primary' : 'border-gray-300'} ${images.length >= maxImages || isUploading ? 'opacity-50 cursor-not-allowed' : ''}`}
            onClick={handleButtonClick}
          >
            <input {...getInputProps()} />
            <Upload size={48} className="text-gray-400 mb-4" />
            <p className="text-sm text-gray-600 text-center">
              {isDragActive ? "Drop the files here ..." : "Drag 'n' drop some files here, or click to select files"}
            </p>
          </div>
        );
    }
  };

  return (
    <>
      {renderUploadButton()}
      <Dialog open={isOpen} onOpenChange={setIsOpen}>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Upload Images</DialogTitle>
            <DialogDescription>
              Upload your images here. You can drag and drop files or click to select them.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div 
              {...getRootProps()} 
              className={`border-2 border-dashed rounded-md p-4 ${isDragActive ? 'border-primary' : 'border-gray-300'} ${images.length >= maxImages || isUploading ? 'opacity-50 cursor-not-allowed' : ''}`}
              onClick={open}
            >
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
                  <img 
                    src={`https://ucarecdn.com/${image}/`} 
                    alt={`Uploaded ${index}`} 
                    className="w-full h-24 object-cover rounded" 
                  />
                  <button
                    onClick={() => handleRemoveImage(index)}
                    className="absolute top-0 right-0 bg-red-500 text-white rounded-full p-1"
                  >
                    <X size={16} />
                  </button>
                </div>
              ))}
            </div>
          </div>
          <Button onClick={() => setIsOpen(false)} disabled={isUploading}>
            {isUploading ? 'Uploading...' : 'Done'}
          </Button>
        </DialogContent>
      </Dialog>
    </>
  );
};

export default UploadImage;