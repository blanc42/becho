"use client"

import { useEffect, useState } from 'react';
import { FileUploaderMinimal, OutputCollectionState, OutputCollectionStatus } from "@uploadcare/react-uploader";
// import "@uploadcare/file-uploader/web/uc-file-uploader-regular.min.css";


interface SignedParams {
    public_key: string;
  signature: string;
  expire: number;
}

const UploadComponent: React.FC = () => {
  const [signedParams, setSignedParams] = useState<SignedParams | null>(null);

  const handleChange = (e: OutputCollectionState<OutputCollectionStatus, "maybe-has-group">) => {
    console.log(e.errors)
    console.log(e.allEntries)
  }


  useEffect(() => {
    try{
      fetch('/api/v1/images/upload')
      .then(res => res.json())
      .then((data: SignedParams) => {
        setSignedParams(data)
        console.log(data)
      });
    } catch (error) {
      console.error('Error fetching signed params:', error);
    }
  }, []);

  if (!signedParams) return <div>Loading...</div>;

  return (
    <FileUploaderMinimal
      pubkey={signedParams.public_key}
      maxLocalFileSizeBytes={2000000}
      multipleMax={5}
      imgOnly={true}
      sourceList="local, camera"
      secureSignature={signedParams.signature}
      secureExpire={signedParams.expire.toString()}
      classNameUploader="my-config"
      onChange={handleChange}
    />
  );
};

export default UploadComponent;