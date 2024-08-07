// middleware.ts
import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'
import jwt from 'jsonwebtoken';

export function middleware(request: NextRequest) {
  const token = request.cookies.get('auth_token')?.value

  if (!token) {
    if (request.nextUrl.pathname.startsWith('/login') || request.nextUrl.pathname.startsWith('/signup')) {
      return NextResponse.next()
    }
    return NextResponse.redirect(new URL('/login', request.url))
  }

  // try {
  //   const isValid = validateToken(token)
  //   if (!isValid) {

  //     const response = NextResponse.redirect(new URL('/login', request.url))
  //     response.cookies.delete('auth_token')
  //     return response
  //   }
  // } catch (error) {

  //   const response = NextResponse.redirect(new URL('/login', request.url))
  //   response.cookies.delete('auth_token')
  //   return response
  // }

  // if (request.nextUrl.pathname.startsWith('/login') || request.nextUrl.pathname.startsWith('/signup')) {
  //   return NextResponse.redirect(new URL('/', request.url))
  // }

  return NextResponse.next()
}

export const config = {
  matcher: '/((?!api|_next/static|_next/image|favicon.ico).*)',
}

function validateToken(token: string): boolean {
  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET as jwt.Secret);
    
    const currentTimestamp = Math.floor(Date.now() / 1000);
    if (typeof decoded === 'object' && decoded.exp && decoded.exp < currentTimestamp) {
      return false;
    }

    // Additional checks can be added here, such as:
    // - Checking if the user still exists in the database
    // - Verifying if the token has been blacklisted

    return true;
  } catch (error) {
    console.error('Error validating token:', error);
    return false;
  }

  return true // Placeholder return
}