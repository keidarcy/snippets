import { Card } from '@twilio-paste/card';
import { Stack } from '@twilio-paste/stack';
import { useSelector } from 'react-redux';
import { Task } from './Task';

export const Human = ({ humanId }) => {
  const human = useSelector((state) =>
    state.humans.find((human) => human.id === humanId)
  );

  return (
    <Card>
      <h2>{human.name}</h2>
      <Stack orientation="vertical" spacing="space60">
        {human.taskIds.map((taskId) => (
          <Task key={taskId} taskId={taskId} />
        ))}
      </Stack>
    </Card>
  );
};
