import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { fetchDogFacts } from '../api/fetchDogFacts';

export const fetchDogFactsFromAPI = createAsyncThunk(
  'facts/fetchFacts',
  async (count) => {
    const facts = await fetchDogFacts(count);
    return facts;
  },
);
export const dogSlice = createSlice({
  name: 'facts',
  initialState: [],
  extraReducers: {
    [fetchDogFactsFromAPI.fulfilled]: (state, action) => {
      return action.payload;
    },
  },
});

// dispatch(fetchDogFactsFromAPI(count));
