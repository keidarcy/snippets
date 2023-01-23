// import { createStore } from 'redux';

// export const store = createStore(
//   (state = { humans: [], tasks: [] }, action) => state
// );

import { configureStore } from '@reduxjs/toolkit';
import { humanSlice } from './peopleSlice';
import { tasksSlice } from './taskSlice';

export const store = configureStore({
  reducer: {
    tasks: tasksSlice.reducer,
    humans: humanSlice.reducer
  }
});
