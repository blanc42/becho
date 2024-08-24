import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ProductType } from "@/lib/types";

export default function ProductCard({ product }: { product: ProductType }) {
    return (
        <Card className="w-full max-w-sm overflow-hidden transition-all duration-300 hover:shadow-lg">
            <CardHeader className="p-0">
                <div className="h-64 w-full flex justify-center">
                    <img
                        src={product.image}
                        alt={product.title}
                        className="object-cover h-full"
                    />
                </div>
            </CardHeader>
            <CardContent className="p-4">
                <div className="flex justify-between items-center">
                    <h3 className="text-lg font-semibold truncate flex-1">{product.title}</h3>
                    <div className="flex items-center space-x-2">
                        <p className="text-lg font-bold text-green-600">${product.price.toFixed(2)}</p>
                        <Button>
                            Add to Cart
                        </Button>
                    </div>
                </div>
            </CardContent>
        </Card>
    );
}