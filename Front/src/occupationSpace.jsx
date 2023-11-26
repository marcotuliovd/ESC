import React, { useState } from 'react';

const EntryVehicle = () => {
  const [dados, setDados] = useState({
    id:'',
    vehicle_id:'',
    type: '',
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
        type: dados.type,
    }
    try {
      const response = await fetch('http://localhost:3000/api/v1/occupationSpace', {
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
      const reuqetss = {
        target: {
            name: 'id',
            value: data.id
        }
      }
      await handleChange(reuqetss)
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

      <label>
        Tamanho do Veículo:
        <select
          name="type"
          value={dados.type}
          onChange={handleChange}
        >
          <option value="">Selecione...</option>
          <option value="pequeno">Pequeno</option>
          <option value="moto">Moto</option>
          <option value="medio">Médio</option>
          <option value="grande">Grande</option>
        </select>
      </label>
      <br />
      <div>
        <h3>Id: {dados.id}</h3>
      </div>

      <button type="submit">Enviar</button>
    </form>
  );
};

export default EntryVehicle;
