import React, { useState } from 'react';

const CreateVehicle = () => {
  const [dados, setDados] = useState({
    id:'',
    owner: '',
    title: '',
    type: '',
    plate:'',
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
        owner: dados.owner,
        title: dados.title,
        type: dados.type,
        plate: dados.plate,
    }
    try {
      const response = await fetch('http://localhost:3000/api/v1/createVehicle', {
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
            name: 'id',
            value: data.id
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
        Dono:
        <input
          type="text"
          name="owner"
          value={dados.owner}
          onChange={handleChange}
        />
      </label>
      <br />

      <label>
        Modelo do Veículo:
        <input
          type="text"
          name="title"
          value={dados.name}
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

      <label>
        Placa do Veículo:
        <input
          type="text"
          name="plate"
          value={dados.plate}
          onChange={handleChange}
        />
      </label>
      <br />
      <div>
        <h3>Id: {dados.id}</h3>
      </div>

      <button type="submit">Enviar</button>
    </form>
  );
};

export default CreateVehicle;
