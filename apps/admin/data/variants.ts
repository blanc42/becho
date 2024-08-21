import { Category, Variant } from "@/lib/types";

export const variantsData: Variant[] = [
    { id: "1", name: "Size", label: "Size", options: [
        { id: "1a", value: "XS", image_url: "", data: "" },
        { id: "1b", value: "S", image_url: "", data: "" },
        { id: "1c", value: "M", image_url: "", data: "" },
        { id: "1d", value: "L", image_url: "", data: "" },
        { id: "1e", value: "XL", image_url: "", data: "" },
        { id: "1f", value: "XXL", image_url: "", data: "" }
    ]},
    { id: "2", name: "Color", label: "Color", options: [
        { id: "2a", value: "Black", image_url: "", data: "" },
        { id: "2b", value: "White", image_url: "", data: "" },
        { id: "2c", value: "Navy", image_url: "", data: "" },
        { id: "2d", value: "Red", image_url: "", data: "" },
        { id: "2e", value: "Green", image_url: "", data: "" },
        { id: "2f", value: "Yellow", image_url: "", data: "" }
    ]},
    { id: "3", name: "RAM", label: "RAM", options: [
        { id: "3a", value: "4GB", image_url: "", data: "" },
        { id: "3b", value: "8GB", image_url: "", data: "" },
        { id: "3c", value: "16GB", image_url: "", data: "" },
        { id: "3d", value: "32GB", image_url: "", data: "" }
    ]},
    { id: "4", name: "Storage", label: "Storage", options: [
        { id: "4a", value: "64GB", image_url: "", data: "" },
        { id: "4b", value: "128GB", image_url: "", data: "" },
        { id: "4c", value: "256GB", image_url: "", data: "" },
        { id: "4d", value: "512GB", image_url: "", data: "" },
        { id: "4e", value: "1TB", image_url: "", data: "" }
    ]},
    { id: "5", name: "Material", label: "Material", options: [
        { id: "5a", value: "Leather", image_url: "", data: "" },
        { id: "5b", value: "Canvas", image_url: "", data: "" },
        { id: "5c", value: "Nylon", image_url: "", data: "" },
        { id: "5d", value: "Polyester", image_url: "", data: "" }
    ]},
    { id: "6", name: "Style", label: "Style", options: [
        { id: "6a", value: "Casual", image_url: "", data: "" },
        { id: "6b", value: "Formal", image_url: "", data: "" },
        { id: "6c", value: "Sports", image_url: "", data: "" },
        { id: "6d", value: "Vintage", image_url: "", data: "" }
    ]}
]

export const categoriesData: Category[] = [
    { id: "1", name: "Clothing", description: "All types of clothing", parent_id: null, variants: ["1","2"], level: 0 },
    { id: "2", name: "Electronics", description: "Electronic devices", parent_id: null, variants: [], level: 0 },
    { id: "3", name: "Accessories", description: "Fashion accessories", parent_id: null, variants: ["5","6"], level: 0 },
    { id: "4", name: "T-Shirts", description: "All types of T-Shirts", parent_id: "1", variants: ["1","2"], level: 1 },
    { id: "5", name: "Jeans", description: "Denim jeans", parent_id: "1", variants: ["1","2"], level: 1 },
    { id: "6", name: "Mobile Phones", description: "Smartphones and feature phones", parent_id: "2", variants: ["3","4"], level: 1 },
    { id: "7", name: "Laptops", description: "Portable computers", parent_id: "2", variants: ["3","4"], level: 1 },
    { id: "8", name: "Bags", description: "All types of bags", parent_id: "3", variants: ["5","6"], level: 1 },
    { id: "9", name: "Watches", description: "Wristwatches", parent_id: "3", variants: ["5","6"], level: 1 }
]