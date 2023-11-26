import React from 'react';

const headerStyle = {
  backgroundColor: '#ffffff', // Cor azul para o cabeçalho
  padding: '20px',
  textAlign: 'center',
  borderRadius: '8px',
  border: '1px solid transparent',

  fontSize: '1em',
  fontWeight: '200',
  fontFamily: 'inherit',
  color: '#213547',
  display: 'flex',

  border: '5px solid #213547'
};

const logoStyle = {
  with: '100px', // Ajuste conforme necessário para o tamanho do seu logo
  height: '100px',
};

function Header() {
  return (
    <div className="Header">
      <header style={headerStyle}>
        <img src="../logo.png" alt="Logo da Empresa" style={logoStyle} />
        <h1>ESC - ESTACIONAMENTO STOCK CAR</h1>
      </header>
    </div>
  );
}

export default Header;
