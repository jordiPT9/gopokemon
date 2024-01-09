import PokemonCard from './PokemonCard';

function App() {
  return (
    <>
      <div style={{ display: 'flex', gap: '0px' }}>
        <PokemonCard />
        <PokemonCard />
      </div>
      <div style={{ display: 'flex', gap: '0px' }}>
        <PokemonCard />
        <PokemonCard />
      </div>
      <div style={{ display: 'flex', gap: '0px' }}>
        <PokemonCard />
        <PokemonCard />
      </div>
    </>
  );
}

export default App;
