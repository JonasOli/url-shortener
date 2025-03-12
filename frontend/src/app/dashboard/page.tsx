'use client';

import { signOut } from '@/app/actions/auth';
import { Button } from '@mui/material';

export default function Dashboard() {
  return (
    <Button
      onClick={async () => {
        await signOut();
      }}
    >
      Logout
    </Button>
  );
}
