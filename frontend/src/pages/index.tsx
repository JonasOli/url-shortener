import { Container } from "@mui/material";
import Link from "next/link";

export default function Page() {
  return (
    <Container maxWidth="sm">
      <h1 style={{ textAlign: "center" }}>URL shortener</h1>
      <Link href={"signIn"}>Sign-in</Link>
      <Link href={"signUp"}>Sign-up</Link>
    </Container>
  );
}
