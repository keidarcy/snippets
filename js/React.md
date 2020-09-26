 - Simplest implemention of redux
 
 ```js
 function createStore(reducer) {
  let state = { state: 1 };

  let listeners = [];

  function dispatch(action) {
    state = reducer(state, action);

    for (let i = 0; i < listeners.length; i++) {
      listeners[i]();
    }
  }

  function subscribe(listener) {
    listeners.push(listener);
  }

  function getState() {
    return state;
  }

  return {
    dispatch,
    getState,
    subscribe
  };
}
const reducer = (state, action) => {
  switch (action.type) {
    case 'action':
      return { ...state, state: action.payload };
    default:
      break;
  }
};

const store = createStore(reducer);

store.subscribe(() => {
  console.log('Store changed!');
});

store.dispatch({ type: 'action', payload: 2 });

const result = store.getState();

console.log(result);
 
 ```
