import { createSlice } from '@reduxjs/toolkit';

interface CounterState {
    value: number;
}

const initialState: CounterState = {
    value: 0,
};

export const counterSlice = createSlice({
    name: 'counter',
    initialState,
    reducers: {
        increment: (state) => {
            state.value += 1;
        },
    },
});

export const { increment } = counterSlice.actions;

export const selectCount = (state: { counter: CounterState }) => state.counter.value;

export default counterSlice.reducer;