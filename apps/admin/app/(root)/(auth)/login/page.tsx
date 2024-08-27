import Link from "next/link";
import { LoginForm } from "./LoginForm";

export default function LoginPage() {
  return (
    <>
      <h2 className="mt-6 text-3xl font-extrabold text-center">
        Log in to your account
      </h2>
      <div className="text-sm -mt-4 text-gray-500 text-center">Test credentials are already filled</div>
      <LoginForm />
    </>
  );
}