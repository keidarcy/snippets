import { bindActionCreators } from 'redux';
import { useDispatch } from 'react-redux';
import { useMemo } from 'react';

export function useActions(actions) {
  const dispatch = useDispatch();
  return useMemo(() => {
    return bindActionCreators(actions, dispatch);
  }, [actions, dispatch]);
}
