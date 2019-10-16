import React,{ useState, useEffect } from 'react';
import './App.css';
import Header from "./components/Header"
import Input from "./components/Input"
import DownloadList from "./components/DownloadList"
import { createGlobalStyle } from "styled-components";
const GlobalStyles = createGlobalStyle`
  body {
    @import url('https://fonts.googleapis.com/css?family=Raleway');
    font-family: 'Raleway', sans-serif;
    color: white;
  }
  `;

function App() {
    const [signedIn, setSignedIn] = useState(false);
    const [postEvent, setPostEvent] = useState(false);

    const authenticate = () => {
        fetch("http://localhost:8080/auth", {
            method: 'GET',
            credentials: 'include',
            headers: {
                "Access-Control-Allow-Credentials":"true"
            },
        }).then(
            (result) => {
                console.log(result);
                if (result.ok === true) {
                    setSignedIn(true)
                }
            }
        );
    };
    const onPostEvent = (result) => {
        if(result.ok){
            setPostEvent(!postEvent);
        }
    };

    useEffect(() => {
        authenticate();
    });
    
    if (signedIn) {
      return (
        <div className="App">
            <GlobalStyles/>
            <Header/>
            <Input onPost={onPostEvent}/>
            <DownloadList postListener={postEvent}/>
        </div>
      );
    }else {
        return (
            <a href="http://localhost:8080/login">Sign In</a>
        );
    }
}

export default App;
