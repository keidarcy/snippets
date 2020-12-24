- Simplest implemention of redux

```js
function createStore(reducer) {
  let state = { state: 1 };

  let listeners = [];

  function dispatch(action) {
    state = reducer(state, action);

    for (let i = 0; i < listeners.length; i++) {
      listeners[i]();
    }
  }

  function subscribe(listener) {
    listeners.push(listener);
  }

  function getState() {
    return state;
  }

  return {
    dispatch,
    getState,
    subscribe
  };
}
const reducer = (state, action) => {
  switch (action.type) {
    case 'action':
      return { ...state, state: action.payload };
    default:
      break;
  }
};

const store = createStore(reducer);

store.subscribe(() => {
  console.log('Store changed!');
});

store.dispatch({ type: 'action', payload: 2 });

const result = store.getState();

console.log(result);
```

## Get query paramater with react router

```tsx
const useQuery = () => {
  const location = useLocation();
  return new URLSearchParams(location.search);
};

const query = useQuery();

query.get('id');
```

## React helmet and default props value

```tsx
import React from 'react';
import { Helmet } from 'react-helmet';

interface MetaProps {
  title?: string;
  description?: string;
  keywords?: string;
}

export const Meta: React.FC<MetaProps> = ({ title, description, keywords }) => {
  return (
    <Helmet>
      <title>{title}</title>
      <meta name="description" content={description} />
      <meta name="keywords" content={keywords} />
    </Helmet>
  );
};

Meta.defaultProps = {
  title: 'Welcome to | Home',
  description: 'We sell the best products for cheap',
  keywords: 'electorincs cheap cool'
};
```

## Switch case usecase in JSX

```tsx
const Notification = ({ text, status }) => {
  const types = ['info', 'warning', 'error'];
  return (
    <div>
      {types.map((type, index) => (
        <div key={type}>
          {(() => {
            switch (status) {
              case 'info':
                return <Info text={text} />;
              case 'warning':
                return <Warning text={text} />;
              case 'error':
                return <Error text={text} />;
              default:
                return null;
            }
          })()}
        </div>
      ))}
    </div>
  );
};
```
