'use server';

export async function signIn(email: string, password: string) {
  fetch('http://localhost:8000/user/login', {
    method: 'POST',
    body: JSON.stringify({ email, password }),
    headers: {
      'Content-Type': 'application/json',
    },
  });
}
