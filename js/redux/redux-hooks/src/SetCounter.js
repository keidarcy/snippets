import { useEffect, useState } from 'react';
import { useCounter } from './use-counter';

export const SetCounter = () => {
  const { count: countFromStore, set } = useCounter();
  const [count, setCount] = useState(0);

  useEffect(() => {
    setCount(countFromStore);
  }, [countFromStore]);

  return (
    <section className="controls">
      <form
        onSubmit={(event) => {
          event.preventDefault();
          set(count);
        }}
      >
        <label htmlFor="set-to">Set Count</label>
        <input
          id="set-to"
          type="number"
          value={count}
          onChange={(event) => setCount(parseInt(event.target.value, 10))}
        />
        <input type="submit" />
      </form>
    </section>
  );
};
