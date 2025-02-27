"use client";

import { Button, Container, Input, InputLabel, styled } from "@mui/material";
import { blue, grey } from "@mui/material/colors";
import { useForm } from "react-hook-form";

export default function SignIn() {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm();
  const onSubmit = (data: any) => console.log(data);

  return (
    <Container maxWidth="sm">
      <h1>Sign in</h1>

      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <InputLabel>Email</InputLabel>
          <InputElement
            placeholder="Enter email..."
            type="email"
            {...register("email")}
          />

          <InputLabel>Password</InputLabel>
          <InputElement
            placeholder="Enter password..."
            type="password"
            {...register("password")}
          />
        </div>

        <Button type="submit">Sign in</Button>
      </form>
    </Container>
  );
}

const InputElement = styled(Input)(
  ({ theme }) => `
  width: 320px;
  font-family: 'IBM Plex Sans', sans-serif;
  font-size: 0.875rem;
  font-weight: 400;
  line-height: 1.5;
  padding: 8px 12px;
  border-radius: 8px;
  color: ${theme.palette.mode === "dark" ? grey[300] : grey[900]};
  background: ${theme.palette.mode === "dark" ? grey[900] : "#fff"};
  border: 1px solid ${theme.palette.mode === "dark" ? grey[700] : grey[200]};
  box-shadow: 0 2px 4px ${
    theme.palette.mode === "dark" ? "rgba(0,0,0, 0.5)" : "rgba(0,0,0, 0.05)"
  };

  &:before, &:after {
    border-bottom: none !important;
  }

  &:hover {
    border-color: ${blue[400]};
  }

  &:focus {
    border-color: ${blue[400]};
    box-shadow: 0 0 0 3px ${
      theme.palette.mode === "dark" ? blue[600] : blue[200]
    };
  }

  /* firefox */
  &:focus-visible {
    outline: 0;
  }
`
);
