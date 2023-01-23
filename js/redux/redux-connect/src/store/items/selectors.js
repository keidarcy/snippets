import { createSelector } from 'reselect';
// import { createSelector } from '@reduxjs/toolkit';

export const selectItem = (state, props) => {
  return state.items.find((item) => item.uuid === props.uuid);
};

export const selectItemTotal = createSelector([selectItem], (item) => {
  return item.price * item.quantity;
});

export const selectItems = (state) => state.items;
export const selectTipPercentage = (state) => state.tipPercentage;
export const selectSubtotal = createSelector([selectItems], (items) => {
  const subtotal = items.reduce((acc, item) => {
    return acc + item.price * item.quantity;
  }, 0);
  return subtotal;
});

export const selectTipAmount = createSelector(
  [selectSubtotal, selectTipPercentage],
  (subtotal, tipPercentage) => {
    const tipAmount = subtotal * (tipPercentage / 100);
    return tipAmount;
  }
);

export const selectTotal = createSelector(
  [selectSubtotal, selectTipAmount],
  (subtotal, tipAmount) => {
    const total = subtotal + tipAmount;
    return total;
  }
);
