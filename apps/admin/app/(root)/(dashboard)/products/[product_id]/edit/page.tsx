"use client"


export default function UpdateProductPage({ params }: { params: { product_id: string } }) {
    const { product_id } = params;

    return (
        <div>
            <h1>Update page for {product_id}</h1>
            <h2>Still working on this page</h2>
        </div>
    )
}