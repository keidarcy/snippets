import { createSlice, nanoid, createAction } from '@reduxjs/toolkit';

const createTask = (title) => ({
  id: nanoid(),
  title,
  completed: false,
  assignedTo: ''
});

const initialState = [
  createTask('Order more cat food'),
  createTask('Sell more catnip')
];

export const tasksSlice = createSlice({
  name: 'tasks',
  initialState,
  reducers: {
    add: (state, action) => {
      state.push(createTask(action.payload));
    },
    // toggle: (state, action) => {
    //   const task = state.find((task) => task.id === action.payload.taskId);
    //   task.completed = action.payload.completed;
    // },
    assignToUser: (state, action) => {
      const task = state.find((task) => task.id === action.payload.taskId);
      task.assignedTo = action.payload.humanId;
    }
  }
});

export const toggleTask = createAction('tasks/toggle', (taskId, completed) => ({
  payload: {
    taskId,
    completed
  }
}));
