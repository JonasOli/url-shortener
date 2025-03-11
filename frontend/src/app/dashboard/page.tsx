'use client';

import { signOut } from '@/app/actions/auth';
import { Button } from '@mui/material';
import React from 'react';

export default function Dashboard() {
  return (
    <Button
      onClick={async () => {
        await fetch('http://localhost:8000/user/signout', {
          method: 'POST',
          credentials: 'include',
        });
        signOut();
      }}
    >
      Logout
    </Button>
  );
}
