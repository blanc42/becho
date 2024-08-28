"use client"
import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useRouter } from "next/navigation";
import { useUser } from '@/lib/store/useUser';
import { toast } from '@/components/ui/use-toast';

const loginSchema = z.object({
  email: z.string().email('Invalid email address'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
});

type LoginFormValues = z.infer<typeof loginSchema>;

export function LoginForm() {
  const {user, setUser} = useUser()
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();
  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: 'test@test.com',
      password: 'test1234',
    },
  });

  useEffect(() => {
    if (user) {
      console.log("User is logged in, redirecting to dashboard");
      router.push('/dashboard');
    }
  }, [user, router]);

  function onSubmit(values: LoginFormValues) {
    setIsLoading(true);
    fetch('/api/v1/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(values),
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('Login failed');
        }
        return response.json();
      })
      .then(res => {
        setUser(res.data)
        console.log('Login successful:', res);
        router.push('/dashboard');
        toast({
          title: 'pushed to dashboard',
          description: res.data.email,
        });
      })
      .catch(error => {
        console.error('Error:', error);
        alert(error.message || 'An error occurred while logging in.');
      })
      .finally(() => {
        setIsLoading(false);
      });
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input type="email" placeholder="john@example.com" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input type="password" placeholder="********" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className='flex flex-col space-y-4'>
          <Button type="submit" disabled={isLoading}>
            {isLoading ? 'Logging in...' : 'Log in'}
          </Button>
          <div className='flex items-center justify-center space-x-2'>
            <span className='text-sm text-gray-500'>Don't have an account?</span>
            <Button type="button" onClick={() => router.push('/signup')} variant='link' className='p-0'>
              Sign up
            </Button>
          </div>
        </div>
      </form>
    </Form>
  );
}