import React, { useState } from 'react';

const VehicleOut = () => {
  const [dados, setDados] = useState({
    amount:'',
    entry:'',
    exit: '',
    type:'',
    vehicle_id:'',
  });


  const handleChange = async(e) => {
    const { name, value } = e.target;
    setDados({
      ...dados,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const request = {
        vehicle_id: dados.vehicle_id,
    }
    try {
      const response = await fetch('http://localhost:3000/api/v1/outSpace', {
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
        console.log("error request:")
        console.error(data);
      }
      const amountAtt = {
        target: {
            name: 'amount',
            value: data.amount
        }
      }
      await handleChange(amountAtt)
    } catch (error) {
      console.error('Erro ao realizar a solicitação:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Id do Veículo:
        <input
          type="text"
          name="vehicle_id"
          value={dados.vehicle_id}
          onChange={handleChange}
        />
      </label>
      <br />
      <div>
        <h3>Valor total a pagar: R$ {dados.amount}</h3>
      </div>

      <button type="submit">Enviar</button>
    </form>
  );
};

export default VehicleOut;
