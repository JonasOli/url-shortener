import { Button, Container } from '@mui/material';
import { useRouter } from 'next/navigation';

export default function Page() {
  const router = useRouter();

  return (
    <Container maxWidth="sm">
      <h1 style={{ textAlign: 'center' }}>URL shortener</h1>

      <div style={{ display: 'flex', justifyContent: 'center', gap: '1rem' }}>
        <Button variant="contained" onClick={() => router.push('/signIn')}>
          Sign-in
        </Button>

        <Button variant="contained" onClick={() => router.push('/signUp')}>
          Sign-up
        </Button>
      </div>
    </Container>
  );
}
