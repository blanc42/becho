import Navbar from "@/components/Navbar";
import { Category, Store } from "@/lib/types";
import {BASE_URL, STORE_ID} from "@/lib/config"

export default async function Layout({ children }: { children: React.ReactNode }) {
    let categories: Category[] = [];

    try {
        const response = await fetch(`${BASE_URL}/${STORE_ID}/categories`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const text = await response.text(); // Get the raw text first
        console.log("Raw response:", text); // Log the raw response

        try {
            categories = JSON.parse(text); // Try to parse the JSON
        } catch (parseError) {
            console.error("JSON parse error:", parseError);
            console.log("First 100 characters of response:", text.slice(0, 100));
        }
    } catch (error) {
        console.error("Fetch error:", error);
    }

    const store: Store = {
        name: "Store 1",
        logo: "https://github.com/shadcn.png",
    };

    console.log("hopefully this is gonna be printed on the server ===>> from layout inside the with-nav layout");

    return (
        <div>
            <Navbar store={store} categories={categories} />
            {children}
        </div>
    );
}