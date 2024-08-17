import Link from "next/link";
import { LoginForm } from "./LoginForm";

export default function LoginPage() {
  return (
    <>
      <h2 className="mt-6 text-3xl font-extrabold text-center">
        Log in to your account
      </h2>
      <Link href={'/products'}>products</Link>
      <LoginForm />
    </>
  );
}