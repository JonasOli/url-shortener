'use client';

import { signIn } from '@/app/actions/auth';
import Container from '@/app/components/Container';
import Input from '@/app/components/Input';
import {
  Alert,
  Button,
  InputLabel,
  Snackbar,
  SnackbarCloseReason,
} from '@mui/material';
import { useActionState, useEffect, useState } from 'react';

const initialState = {
  error: '',
};

export default function SignIn() {
  const [openErrorToast, setOpenErrorToast] = useState(false);

  const [state, formAction, isPending] = useActionState(signIn, initialState);

  useEffect(() => {
    if (state?.error) {
      setOpenErrorToast(true);
    }
  }, [state]);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: SnackbarCloseReason
  ) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpenErrorToast(false);
  };

  return (
    <Container maxWidth="sm">
      <h1>Sign in</h1>

      <Snackbar
        open={openErrorToast}
        autoHideDuration={5000}
        onClose={handleClose}
      >
        <Alert severity="error" variant="filled" sx={{ width: '100%' }}>
          {state?.error}
        </Alert>
      </Snackbar>

      <form action={formAction}>
        <div>
          <InputLabel>Email</InputLabel>
          <Input
            name="email"
            placeholder="Enter email..."
            type="email"
            required
          />

          <InputLabel>Password</InputLabel>
          <Input
            name="password"
            placeholder="Enter password..."
            type="password"
            required
          />
        </div>

        <Button
          variant="contained"
          type="submit"
          loading={isPending}
          disabled={isPending}
        >
          Sign in
        </Button>
      </form>
    </Container>
  );
}
