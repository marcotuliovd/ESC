import React, { useState } from 'react';

const CreateUser = () => {
  const [dados, setDados] = useState({
    id:'',
    username: '',
    name: '',
    phone: '',
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
        username: dados.username,
        name: dados.name,
        phone: dados.phone,
    }
    try {
      const response = await fetch('http://localhost:3000/api/v1/createUser', {
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
        Username:
        <input
          type="text"
          name="username"
          value={dados.username}
          onChange={handleChange}
        />
      </label>
      <br />

      <label>
        Nome:
        <input
          type="text"
          name="name"
          value={dados.name}
          onChange={handleChange}
        />
      </label>
      <br />

      <label>
        Telefone:
        <input
          type="text"
          name="phone"
          value={dados.phone}
          onChange={handleChange}
        />
      </label>
      <br />
      <div>
        <h3>ID: {dados.id}</h3>
      </div>

      <button type="submit">Enviar</button>
    </form>
  );
};

export default CreateUser;
