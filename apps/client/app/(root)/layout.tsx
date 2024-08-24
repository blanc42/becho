"use client"
import Navbar from "@/components/Navbar";


export default function RootLayout({ children }: { children: React.ReactNode }) {

    // user fetching logic and all happens here

    console.log("hopefully this is gonna be printed on the client ===>> from layout at the top");
    return (
        <div>
            {children}
        </div>
    );
}