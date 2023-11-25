import React, { useState } from 'react';

const ReportService = () => {
  const [dados, setDados] = useState({
    history:'',
    init: '',
    finish:'',
    user:'',
    password:'',
    amount: '',
  });


  const handleChange = async(e) => {
    const { name, value } = e.target;
    setDados({
      ...dados,
      [name]: value,
    });
  };

  const convertHour = (hora) => {
    return hora + ":05Z";
  }

  const handleSubmit = async (e) => {
    e.preventDefault();
    const request = {
        init: convertHour(dados.init),
        finish: convertHour(dados.finish),
        user: dados.user,
        password: dados.password,
    }
    try {
      const response = await fetch('http://localhost:3000/api/v1/report', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(request),
      });
      const data = await response.json();

      if (response.ok) {
        console.log(data);
      } else {
        console.error(data);
      }
      const reuqetss = {
        target: {
            name: 'amount',
            value: data.amountTotal
        }
      }
      await handleChange(reuqetss)

      console.log("linha 51",dados.id)
    } catch (error) {
      console.error('Erro ao realizar a solicitação:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        User:
        <input
          type="text"
          name="user"
          value={dados.user}
          onChange={handleChange}
        />
      </label>
      <br />

      <label>
        Password:
        <input
          type="password"
          name="password"
          value={dados.password}
          onChange={handleChange}
        />
      </label>
      <br />

      <label>
        Data/Hora inicial:
        <input
          type="datetime-local"
          name="init"
          value={dados.init}
          onChange={handleChange}
        />
      </label>
      <br />
      <label>
        Data/Hora final:
        <input
          type="datetime-local"
          name="finish"
          value={dados.finish}
          onChange={handleChange}
        />
      </label>
      <br />
      <div>
        <h3>Valor Total Faturado: R${dados.amount}</h3>
      </div>

      <button type="submit">Enviar</button>
    </form>
  );
};

export default ReportService;
