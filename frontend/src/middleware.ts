import type { NextRequest } from 'next/server';
import { NextResponse } from 'next/server';

// Define protected routes
const protectedRoutes = ['/dashboard'];

export default async function middleware(req: NextRequest) {
  const token = req.cookies.get('session-id')?.value; // Assuming token is stored in a cookie

  // Check if the request path is protected
  if (protectedRoutes.includes(req.nextUrl.pathname)) {
    if (!token) {
      // Redirect to login if not authenticated
      return NextResponse.redirect(new URL('/', req.url));
    }
  }

  return NextResponse.next(); // Allow request to proceed
}

// Define paths where middleware should run
export const config = {
  matcher: ['/dashboard'], // Only apply middleware to these paths
};
