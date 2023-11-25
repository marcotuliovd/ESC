import './App.css'
import CreateVehicle from './createVehicle';
import EntryVehicle from './occupationSpace';
import CreateUser from './page';
import ReportService from './reportService';
import VehicleOut from './vehicleOut';

function App() {
  return (
    <div>
      <h1>Formulário de Criação de Cliente</h1>
      <CreateUser />
      <br></br>
      <h1>Formulário de Cadastro de Veículo</h1>
      <CreateVehicle />
      <h1>Entrada de Veículo no Estacionamento</h1>
      <EntryVehicle />
      <h1>Saída do Veículo no Estacionamento</h1>
      <VehicleOut />
      <h1>Relatório de Faturamento</h1>
      <ReportService/>
    </div>
  );
}

export default App
