import React from 'react';
import './App.css';
import Header from "./components/Header"
import Input from "./components/Input"
import DownloadList from "./components/DownloadList"

function App() {
  return (
    <div className="App">
      <Header />
      <Input />
      <DownloadList />
    </div>
  );
}

export default App;
