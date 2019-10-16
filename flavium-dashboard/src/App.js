import React,{ useState } from 'react';
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
    const [torrentPosted, setTorrentPosted] = useState(false);

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
    const onTorrentPost = (result) => {
        if(result.ok){
            setTorrentPosted(!torrentPosted);
        }
    }

    authenticate();
    
    if (signedIn) {
      return (
        <div className="App">
            <GlobalStyles/>
            <Header/>
            <Input onPost={onTorrentPost}/>
            <DownloadList postListener={torrentPosted}/>
        </div>
      );
    }else {
        return (
            <a href="http://localhost:8080/login">Sign In</a>
        );
    }
}

export default App;
