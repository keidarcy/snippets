import { configureStore } from '@reduxjs/toolkit';
import { dogSlice } from './factsSlice';

export const store = configureStore({
  reducer: {
    facts: dogSlice.reducer,
  },
});
