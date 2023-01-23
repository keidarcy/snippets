import Characters from './Characters';
import FetchCharacters from './FetchCharacters';

export const Application = () => {
  return (
    <div className="Application">
      <h1>Star Wars Characters</h1>
      <FetchCharacters />
      <Characters />
    </div>
  );
};
