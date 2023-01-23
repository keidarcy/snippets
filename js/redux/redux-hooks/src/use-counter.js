import { useMemo } from 'react';
import { useSelector } from 'react-redux';
import { decrement, increment, set } from './actions';
import { useActions } from './use-actions';

export const useCounter = () => {
  const count = useSelector((state) => state.count);
  const actions = useActions({ increment, decrement, set });
  return useMemo(() => {
    return { count, ...actions };
  }, [count, actions]);
};
