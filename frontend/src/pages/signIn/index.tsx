import { signIn } from '@/action/signin';
import Input from '@/components/Input';
import { Button, Container, InputLabel, styled } from '@mui/material';
import { grey } from '@mui/material/colors';
import { useRouter } from 'next/navigation';
import { useForm } from 'react-hook-form';

type FormData = {
  email: string;
  password: string;
};

export default function SignIn() {
  const router = useRouter();

  const { register, handleSubmit } = useForm<FormData>();

  const onSubmit = async (data: FormData) => {
    try {
      await signIn(data.email, data.password);
    } catch (e) {
      console.error(e);
    }

    router.push('/');
  };

  return (
    <CustomContainer maxWidth="sm">
      <h1>Sign in</h1>

      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <InputLabel>Email</InputLabel>
          <Input
            placeholder="Enter email..."
            type="email"
            {...register('email')}
          />

          <InputLabel>Password</InputLabel>
          <Input
            placeholder="Enter password..."
            type="password"
            {...register('password')}
          />
        </div>

        <Button variant="contained" type="submit">
          Sign in
        </Button>
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
