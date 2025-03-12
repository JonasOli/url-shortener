'use client';

import { Container } from '@mui/material';
import { grey } from '@mui/material/colors';
import styled from 'styled-components';

export default styled(Container)`
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
