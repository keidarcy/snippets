import { Box } from '@twilio-paste/box';
import { Heading } from '@twilio-paste/heading';
import { Stack } from '@twilio-paste/stack';
import { useSelector } from 'react-redux';
import { CreateTask } from './CreateTask';
import { Task } from './Task';

export const AllTasks = () => {
  const tasks = useSelector((state) => state.tasks);

  return (
    <Box width="100%" margin="space100" padding="space50">
      <Heading>Tasks</Heading>
      <CreateTask />
      <Stack orientation="vertical" spacing="space60">
        {tasks.map((task) => (
          <Task key={task.id} taskId={task.id} />
        ))}
      </Stack>
    </Box>
  );
};
