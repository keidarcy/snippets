import { useSelector } from 'react-redux';
import Character from './Character';
import { Loading } from './Loading';

export const Characters = () => {
  const characters = useSelector((state) => state.characters.data);
  const loading = useSelector((state) => state.characters.loading);

  return (
    <section className="Characters">
      {loading && <Loading />}
      {characters.map((character) => (
        <Character key={character.id} character={character} />
      ))}
    </section>
  );
};

export default Characters;
