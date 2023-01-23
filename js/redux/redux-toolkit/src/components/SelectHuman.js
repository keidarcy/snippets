import { Box } from '@twilio-paste/box';
import { Select, Option } from '@twilio-paste/select';
import { useDispatch, useSelector } from 'react-redux';
import { tasksSlice } from '../store/taskSlice';

export const SelectHuman = ({ task }) => {
  const dispatch = useDispatch();
  const humans = useSelector((state) => state.humans);

  return (
    <Box marginRight="space40" width="100%">
      <Select
        className="SelectHuman"
        value={task.assignedTo}
        onChange={(event) => {
          dispatch(
            tasksSlice.actions.assignToUser({
              humanId: event.target.value,
              taskId: task.id
            })
          );
        }}
      >
        <Option>(Unassigned)</Option>
        {humans.map((human) => (
          <Option value={human.id} key={human.id}>
            {human.name}
          </Option>
        ))}
      </Select>
    </Box>
  );
};
