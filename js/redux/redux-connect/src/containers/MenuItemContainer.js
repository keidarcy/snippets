import { connect } from 'react-redux';
import MenuItem from '../components/MenuItem';
import {
  removeItem,
  updatePrice,
  updateQuantity
} from '../store/items/actions';
import { selectItemTotal } from '../store/items/selectors';

const mapStateToProps = (state, props) => {
  return {
    total: selectItemTotal(state, props)
  };
};
const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    remove: () => dispatch(removeItem(ownProps.uuid)),
    updatePrice: (price) => dispatch(updatePrice(ownProps.uuid, price)),
    updateQuantity: (quantity) =>
      dispatch(updateQuantity(ownProps.uuid, quantity))
  };
};
export const MenuItemContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(MenuItem);
