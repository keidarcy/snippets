import shuffle from 'lodash.shuffle';
import { data } from './data';

// These are here just in case the API goes down. 🤷

export const fetchDogFacts = (count) => {
  return Promise.resolve(data).then((facts) => shuffle(facts).slice(0, count));
};
