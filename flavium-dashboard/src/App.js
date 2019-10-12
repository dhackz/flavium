import React from 'react';
import './App.css';
import Header from "./components/Header"
import Input from "./components/Input"
import DownloadList from "./components/DownloadList"
import { createGlobalStyle } from "styled-components";
const GlobalStyles = createGlobalStyle`
  body {
    @import url('https://fonts.googleapis.com/css?family=Raleway');
    font-family: 'Raleway', sans-serif;
  }
  `;
function App() {
  return (
    <div className="App">
      <GlobalStyles />
      <Header />
      <Input />
      <DownloadList />
    </div>
  );
}

export default App;
