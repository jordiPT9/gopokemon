import { useEffect, useState } from 'react';
import PokemonCard from './PokemonCard';

function App() {
  const pokemons = [
    {
      name: 'Emboar',
      hp: 110,
      atk: 123,
      def: 65,
      spAtk: 100,
      spDef: 65,
      speed: 65,
      total: 528,
    },
    {
      name: 'Stoutland',
      hp: 85,
      atk: 110,
      def: 90,
      spAtk: 45,
      spDef: 90,
      speed: 80,
      total: 500,
    },
    {
      name: 'Zebstrika',
      hp: 75,
      atk: 100,
      def: 63,
      spAtk: 80,
      spDef: 63,
      speed: 116,
      total: 497,
    },
    {
      name: 'Seismitoad',
      hp: 105,
      atk: 95,
      def: 75,
      spAtk: 85,
      spDef: 75,
      speed: 74,
      total: 509,
    },
    {
      name: 'Gothitelle',
      hp: 70,
      atk: 55,
      def: 95,
      spAtk: 95,
      spDef: 110,
      speed: 65,
      total: 490,
    },
    {
      name: 'Chandelure',
      hp: 60,
      atk: 55,
      def: 90,
      spAtk: 145,
      spDef: 90,
      speed: 80,
      total: 520,
    },
  ];

  const [inputContent, setInputContent] = useState('');
  const [pokemonList, setPokemonList] = useState([]);

  const fetchPokemons = async () => {
    const headers = { 'Content-Type': 'application/json' };
    const reqBody = ['Emboar', 'Stoutland', 'Zebstrika', 'Seismitoad', 'Gothitelle', 'Chandelure'];
    const options = {
      headers,
      method: 'GET',
      body: JSON.stringify({ reqBody }),
    };
    return await fetch('http://localhost:5000/best-team', options)
      .then(async (response) => await response.json())
      .catch((err) => console.log(err));
  };

  useEffect(() => {
    try {
      const response = fetchPokemons();
      console.log(response);
      // setPokemonList(data);
    } catch (err) {
      console.log(err);
    }
  }, []);

  return (
    <>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <input
          style={{ textAlign: 'center' }}
          onChange={(e) => {
            console.log(e.target.value);
            setInputContent(e.target.value);
          }}
        />
        <button
          onClick={() => {
            const response = fetchPokemons();
            console.log(response);
          }}
        >
          Find best team
        </button>
      </div>
      <div>
        <div style={{ display: 'block', gap: '0px', float: 'left' }}>
          <PokemonCard {...pokemons[0]} />
          <PokemonCard {...pokemons[1]} />
          <PokemonCard {...pokemons[2]} />
        </div>
        <div style={{ display: 'block', gap: '0px', float: 'right' }}>
          <PokemonCard {...pokemons[3]} />
          <PokemonCard {...pokemons[4]} />
          <PokemonCard {...pokemons[5]} />
        </div>
      </div>
    </>
  );
}

export default App;
