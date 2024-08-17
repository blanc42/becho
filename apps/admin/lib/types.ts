interface Variant {
    id: string
    name: string
}

interface Category {
    id: string
    name: string
    description: string
    parent_id: string | null
    variants: any | null
    level: number
    path: string
    subCategories?: Category[]
}



export type { Variant, Category }