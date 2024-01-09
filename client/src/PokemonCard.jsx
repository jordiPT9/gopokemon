import React from 'react';
import './PokemonCard.css';

const PokemonCard = () => {
  const calculateColor = (value) => {
    const constant = 200
    const stat = (value * constant) / 255
    if (stat < (constant/2)) {
      return `rgb(${(constant - stat * 2)*2}, ${(stat * 2)*2}, ${0})`;
    } else if (stat >= (constant/2)) {
      return `rgb(${0}, ${((constant - stat) * 2)}, ${(((stat - (constant/2)) * 2))*3})`;
    }
  };

  return (
    <div className="pokemon-card">
      <div className="pokemon-name">Emboar</div>
      <div className="stat-wrapper">
        <span className="stat-label">Hp</span>
        <span className="stat-value">110</span>
        <div
          className="stat-bar"
          style={{
            width: 'calc(160 / 255 * 450px)',
            backgroundColor: calculateColor(160),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Attack</span>
        <span className="stat-value">123</span>
        <div
          className="stat-bar"
          style={{
            width: 'calc(123 / 255 * 450px)',
            backgroundColor: calculateColor(123),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Defense</span>
        <span className="stat-value">65</span>
        <div
          className="stat-bar"
          style={{
            width: 'calc(65 / 255 * 450px)',
            backgroundColor: calculateColor(65),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Sp. Atk</span>
        <span className="stat-value">100</span>
        <div
          className="stat-bar"
          style={{
            width: 'calc(100 / 255 * 450px)',
            backgroundColor: calculateColor(100),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Sp. Def</span>
        <span className="stat-value">65</span>
        <div
          className="stat-bar"
          style={{
            width: 'calc(65 / 255 * 450px)',
            backgroundColor: calculateColor(65),
          }}
        ></div>
      </div>
      <div className="stat-wrapper">
        <span className="stat-label">Speed</span>
        <span className="stat-value">65</span>
        <div
          className="stat-bar"
          style={{
            width: 'calc(65 / 255 * 450px)',
            backgroundColor: calculateColor(65),
          }}
        ></div>
      </div>
      <div
        className="stat-wrapper"
        style={{ fontWeight: 'bold' }}
      >
        <span className="stat-label">Total</span>
        <span className="stat-value">528</span>
      </div>
    </div>
  );
};

export default PokemonCard;
