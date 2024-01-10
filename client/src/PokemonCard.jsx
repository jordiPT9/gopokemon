import React from 'react';
import PropTypes from 'prop-types';
import './PokemonCard.css';

const PokemonCard = ({ name, hp, atk, def, spAtk, spDef, speed, total }) => {
  const getBgColor = (statValue) => {
    const max = 200;
    const value = (statValue * max) / 255;

    let r = (max - value * 2) * 2 + (max - value * 2) * 2 * 0.5;
    let g = value * 2 + value * 2 * 0.5 + 0;
    let b = g >= 255 ? value : 0;

    return `rgb(${r}, ${g}, ${b})`;
  };

  return (
    <div className="pokemon-card">
      <div className="pokemon-name">{name}</div>
      <div className="stat-wrapper">
        <span className="stat-label">Hp</span>
        <span className="stat-value">{hp}</span>
        <div
          className="stat-bar"
          style={{
            width: `calc(${hp} / 255 * 450px)`,
            backgroundColor: getBgColor(hp),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Attack</span>
        <span className="stat-value">{atk}</span>
        <div
          className="stat-bar"
          style={{
            width: `calc(${atk} / 255 * 450px)`,
            backgroundColor: getBgColor(atk),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Defense</span>
        <span className="stat-value">{def}</span>
        <div
          className="stat-bar"
          style={{
            width: `calc(${def} / 255 * 450px)`,
            backgroundColor: getBgColor(def),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Sp. Atk</span>
        <span className="stat-value">{spAtk}</span>
        <div
          className="stat-bar"
          style={{
            width: `calc(${spAtk} / 255 * 450px)`,
            backgroundColor: getBgColor(spAtk),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Sp. Def</span>
        <span className="stat-value">{spDef}</span>
        <div
          className="stat-bar"
          style={{
            width: `calc(${spDef} / 255 * 450px)`,
            backgroundColor: getBgColor(spDef),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Speed</span>
        <span className="stat-value">{speed}</span>
        <div
          className="stat-bar"
          style={{
            width: `calc(${speed} / 255 * 450px)`,
            backgroundColor: getBgColor(speed),
          }}
        ></div>
      </div>
      <div
        className="stat-wrapper"
        style={{ fontWeight: 'bold' }}
      >
        <span className="stat-label">Total</span>
        <span className="stat-value">{total}</span>
      </div>
    </div>
  );
};

PokemonCard.propTypes = {
  name: PropTypes.string,
  hp: PropTypes.number,
  atk: PropTypes.number,
  def: PropTypes.number,
  spAtk: PropTypes.number,
  spDef: PropTypes.number,
  speed: PropTypes.number,
  total: PropTypes.number,
};

PokemonCard.defaultProps = {
  name: 'Unknown',
  hp: 0,
  atk: 0,
  def: 0,
  spAtk: 0,
  spDef: 0,
  speed: 0,
  total: 0,
};

export default PokemonCard;
