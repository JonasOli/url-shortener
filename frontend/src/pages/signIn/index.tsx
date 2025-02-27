import Input from "@/components/Input";
import { Button, Container, InputLabel, styled } from "@mui/material";
import { grey } from "@mui/material/colors";
import { useForm } from "react-hook-form";

export default function SignIn() {
  const { register, handleSubmit } = useForm();
  const onSubmit = (data: unknown) => console.log(data);

  return (
    <CustomContainer maxWidth="sm">
      <h1>Sign in</h1>

      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <InputLabel>Email</InputLabel>
          <Input
            placeholder="Enter email..."
            type="email"
            {...register("email")}
          />

          <InputLabel>Password</InputLabel>
          <Input
            placeholder="Enter password..."
            type="password"
            {...register("password")}
          />
        </div>

        <Button type="submit">Sign in</Button>
      </form>
    </CustomContainer>
  );
}

const CustomContainer = styled(Container)`
  border: 1px solid ${grey[300]};
  border-radius: 5px;
  padding: 2rem;

  & {
    label {
      margin-top: 1rem;
    }

    button {
      margin-top: 1rem;
    }
  }
`;
