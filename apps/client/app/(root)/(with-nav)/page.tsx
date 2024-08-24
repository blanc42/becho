import ImageCarousel from "@/components/ImageCarasol";
import ProductCard from "@/components/ProductCard";
import { ProductType } from "@/lib/types";

export default async function Home() {
  const products : ProductType[] = await fetch("https://fakestoreapi.com/products").then(
    (res) => res.json()
  );

  interface CarouselImage {
    src: string;
    alt: string;
    link: string;
  }

  const images: CarouselImage[] = [
    {
      src: "https://picsum.photos/id/1018/1000/600",
      alt: "Mountain landscape",
      link: "/category/outdoors"
    },
    {
      src: "https://picsum.photos/id/1015/1000/600",
      alt: "River in forest",
      link: "/category/nature"
    },
    {
      src: "https://picsum.photos/id/1019/1000/600",
      alt: "Cityscape at night",
      link: "/category/urban"
    }
  ];

  return (
    <>
    <main>
      <ImageCarousel images={images}/>
      <div className="grid grid-cols-3 gap-3 my-6">
        {products.map((product: ProductType) => (
          <ProductCard product={product}/>
        ))}
        </div>
    </main>
        </>
  );
}
