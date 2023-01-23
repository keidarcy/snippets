const Character = ({ character }) => {
  const { id, name } = character;
  return (
    <article className="Character">
      <a href={`https://star-wars-characters.glitch.me/api/characters/${id}`}>
        {name}
      </a>
    </article>
  );
};

export default Character;
