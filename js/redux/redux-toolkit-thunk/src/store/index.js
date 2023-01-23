import { configureStore } from '@reduxjs/toolkit';
import { charactersSlice } from './charactersSlice';

export const store = configureStore({
  reducer: { characters: charactersSlice.reducer },
});
