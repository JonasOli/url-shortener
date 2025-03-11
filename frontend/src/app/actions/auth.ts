'use server';

import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export async function signOut() {
  const cok = await cookies();

  cok.delete('session-id');

  redirect('/');
}
