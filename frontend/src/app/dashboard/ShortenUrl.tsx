'use client';

import { Button } from '@mui/material';
import { useActionState } from 'react';
import { shortenLink } from '../actions/urls';
import Input from '../components/Input';

export default function ShortenUrl() {
  const initialState = {
    error: '',
  };
  const [state, formAction, isPending] = useActionState(
    shortenLink,
    initialState
  );

  if (state?.error) {
    console.error(state.error);
  }

  return (
    <form action={formAction}>
      <Input name="url" placeholder="Place your long URL here..." />
      <Button
        variant="contained"
        type="submit"
        loading={isPending}
        disabled={isPending}
      >
        Shorten link
      </Button>
    </form>
  );
}
