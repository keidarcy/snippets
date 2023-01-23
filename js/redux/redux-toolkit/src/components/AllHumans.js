import { Box } from '@twilio-paste/box';
import { Heading } from '@twilio-paste/heading';
import { Stack } from '@twilio-paste/stack';
import { useSelector } from 'react-redux';
import { CreateHuman } from './CreateHuman';
import { Human } from './Human';

export const AllHumans = () => {
  const humans = useSelector((state) => state.humans);

  return (
    <Box width="100%" margin="space100" padding="space50">
      <Heading>Humans</Heading>
      <CreateHuman />
      <Stack orientation="vertical" spacing="space60">
        {humans.map((human) => (
          <Human key={human.id} humanId={human.id} />
        ))}
      </Stack>
    </Box>
  );
};
