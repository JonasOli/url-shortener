'use server';

import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export async function signIn(prevState: unknown, data: FormData) {
  const res = await fetch('http://localhost:8000/user/login', {
    method: 'POST',
    body: JSON.stringify({
      email: data.get('email') as string,
      password: data.get('password') as string,
    }),
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });

  if (res.status === 400) {
    return { error: 'Invalid email or password' };
  }

  if (res.status === 204) {
    const cookieStore = await cookies();
    const sessionIdCookie = res.headers
      .getSetCookie()[0]
      .split('session-id=')[1];
    cookieStore.set('session-id', sessionIdCookie);

    redirect('/dashboard');
  }
}

export async function signUp(prevState: unknown, data: FormData) {
  const res = await fetch('http://localhost:8000/user/signup', {
    method: 'POST',
    body: JSON.stringify({
      name: data.get('name') as string,
      email: data.get('email') as string,
      password: data.get('password') as string,
    }),
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (res.status === 400) {
    return { error: 'Invalid email or password' };
  }

  if (res.status === 201) {
    redirect('/');
  }
}

export async function signOut() {
  const cookieStore = await cookies();
  const sessionId = cookieStore.get('session-id');

  const res = await fetch('http://localhost:8000/user/signout', {
    method: 'POST',
    credentials: 'include',
    headers: {
      ['Cookie']: `${sessionId?.name}=${sessionId?.value}`,
    },
  });

  if (res.status === 200) {
    const cok = await cookies();

    cok.delete('session-id');

    redirect('/');
  }
}
