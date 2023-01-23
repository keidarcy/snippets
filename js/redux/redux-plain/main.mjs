import {
  createStore,
  compose,
  applyMiddleware,
  combineReducers,
  bindActionCreators,
} from 'redux';
import { performance } from 'perf_hooks';

/** createStore, compose, bindActionCreators
// compose
// const repeatThreeTimes = (str) => str.repeat(3);
// const uppercase = (str) => str.toUpperCase();
// const makeBold = (str) => str.bold();

// console.log(makeBold(repeatThreeTimes(makeBold(uppercase('hello redux')))));
// const haha = compose(repeatThreeTimes, makeBold, uppercase);
// console.log(haha('rerer'));

// reducer
const initialState = { value: 0 };
const INCREMENT = 'INCREMENT';
const ADD = 'ADD';
const increment = () => ({ type: INCREMENT });
const add = (amount) => ({ type: ADD, payload: amount });
const reducer = (state = initialState, action) => {
  if (action.type === INCREMENT) {
    return { value: state.value + 1 };
  }
  if (action.type === ADD) {
    return { value: state.value + action.payload };
  }
  return state;
};
const store = createStore(reducer);

// store.dispatch
store.dispatch(increment());
store.dispatch(increment());

// store.subscribe
const subscriber = () => {
  console.log('Subscribe', store.getState());
};

// store.subscribe(subscriber);
// store.dispatch(increment());
// store.dispatch(add(10));
// store.dispatch(increment());

// bindActionCreators
const actions = bindActionCreators({ increment, add }, store.dispatch);

const { increment: dispatchIncrement, add: dispatchAdd } = bindActionCreators(
  { increment, add },
  store.dispatch,
);

actions.add(1000);
actions.increment();

console.log(store.getState());
*/

/** combineReducers
const initialState = {
  users: [
    { id: 1, name: 'John' },
    { id: 2, name: 'Jane' },
  ],
  tasks: [{ title: 'File the TPS reports' }, { title: 'Call the accountant' }],
};

const ADD_TASK = 'ADD_TASK';
const ADD_USER = 'ADD_USER';

const addTask = (title) => ({ type: ADD_TASK, payload: title });
const addUser = (name) => ({ type: ADD_USER, payload: name });

const userReducer = (users = initialState.users, action) => {
  switch (action.type) {
    case ADD_USER:
      return [...users, { id: Math.random(), name: action.payload }];
    default:
      return users;
  }
};

const taskReducer = (tasks = initialState.tasks, action) => {
  switch (action.type) {
    case ADD_TASK:
      return [...tasks, { title: action.payload }];
    default:
      return tasks;
  }
};

// const reducer = (state = initialState, action) => {
//   if (action.type === ADD_TASK) {
//     return {
//       ...state,
//       tasks: [...state.tasks, action.payload],
//     };
//   }
//   if (action.type === ADD_USER) {
//     return {
//       ...state,
//       users: [...state.users, action.payload],
//     };
//   }
//   return state;
// };
const reducer = combineReducers({ users: userReducer, tasks: taskReducer });
const store = createStore(reducer);

store.dispatch(addTask('File the SFTP reports'));
store.dispatch(addUser('Jane'));

console.log('store.getState():', store.getState());
*/

// /** enhancers applyMiddleware

const reducer = (state = { value: 2 }) => state;
const monitorEnhancer = (createStore) => (reducer, initialState, enhancers) => {
  const monitoredReducer = (state, action) => {
    const start = performance.now();
    const newState = reducer(state, action);
    const end = performance.now();
    const diff = end - start;
    console.log('diff:', diff);
    return newState;
  };
  return createStore(monitoredReducer, initialState, enhancers);
};
const logEnhancer = (createStore) => (reducer, initialState, enhancers) => {
  const logReducer = (state, action) => {
    console.log('old state:', state, action);
    const newState = reducer(state, action);
    console.log('new state:', state, action);
    return newState;
  };
  return createStore(logReducer, initialState, enhancers);
};
// const store = createStore(reducer, {}, logEnhancer);
// const store = createStore(reducer, compose(logEnhancer, monitorEnhancer));
// */

const logMiddleware = (store) => (next) => (action) => {
  console.log('old state:', store.getState(), action);
  const result = next(action);
  console.log('new state:', store.getState(), action);
  return result;
};

const monitorMiddleware = (store) => (next) => (action) => {
  const start = performance.now();
  const result = next(action);
  const end = performance.now();
  const diff = end - start;
  console.log('diff:', diff);
  return result;
};

const store = createStore(reducer, applyMiddleware(logMiddleware, monitorMiddleware));

store.dispatch({ type: 'hello' });
store.dispatch({ type: 'hello' });
