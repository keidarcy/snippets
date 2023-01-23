import { connect } from 'react-redux';
import { Summary } from '../components/Summary';
import {
  selectSubtotal,
  selectTipAmount,
  selectTotal
} from '../store/items/selectors';

const mapStateToProps = (state) => {
  // const subtotal = state.items.reduce((acc, item) => {
  //   return acc + item.price * item.quantity;
  // }, 0);
  // const tipAmount = subtotal * 0.15;
  // const total = subtotal + (state.tipPercentage * tipAmount) / 100;
  const subtotal = selectSubtotal(state);
  const tipAmount = selectTipAmount(state);
  const total = selectTotal(state);
  return {
    subtotal,
    tipAmount,
    total
  };
};

export const SummaryContainer = connect(mapStateToProps)(Summary);
