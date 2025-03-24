'use server';

import { revalidatePath } from 'next/cache';
import { cookies } from 'next/headers';

export async function shortenLink(prevState: unknown, data: FormData) {
  const cookieStore = await cookies();
  const sessionId = cookieStore.get('session-id');

  const res = await fetch('http://localhost:8000/urls/shorten', {
    method: 'POST',
    body: JSON.stringify({
      url: data.get('url') as string,
    }),
    headers: {
      'Content-Type': 'application/json',
      ['Cookie']: `${sessionId?.name}=${sessionId?.value}`,
    },
    credentials: 'include',
  });

  if (res.status !== 200) {
    revalidatePath('/dashboard');
    return { error: 'error' };
  }
}
