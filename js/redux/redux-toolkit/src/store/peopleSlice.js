import { createSlice, nanoid } from '@reduxjs/toolkit';
import { tasksSlice } from './taskSlice';

const createHuman = (name) => ({
  id: nanoid(),
  name,
  taskIds: []
});

const initialState = [
  createHuman('Jane'),
  createHuman('John'),
  createHuman('Jim')
];

export const humanSlice = createSlice({
  name: 'human',
  initialState,
  reducers: {
    add: (state, action) => {
      state.push(createHuman(action.payload));
    }
  },
  extraReducers: (builder) => {
    builder.addCase(tasksSlice.actions.assignToUser, (state, action) => {
      for (const human of state) {
        if (human.id === action.payload.humanId) {
          human.taskIds.push(action.payload.taskId);
        } else {
          human.taskIds = human.taskIds.filter(
            (id) => id !== action.payload.taskId
          );
        }
      }
    });
  }
  // [tasksSlice.actions.add.type]: (state, action) => {
});
