import { Flex } from '@twilio-paste/flex';
import { AllHumans } from './AllHumans';
import { AllTasks } from './AllTasks';

const Application = () => {
  return (
    <Flex margin="auto" width="600">
      <AllHumans />
      <AllTasks />
    </Flex>
  );
};

export default Application;
