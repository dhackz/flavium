import React, { useState }  from 'react';
import { StyledInput, InputArea, StyledButton } from "./styles"

const Input = () => {

  const [text, setText] = useState("");

  const getTorrent = () => {
      console.log("adasd");
    //TODO: Send request to start downloading the magnetlink in the inputfield
      fetch("http://localhost:8080/v1/torrent", {
          method: 'POST',
          credentials: 'include',
          headers: {
              "Access-Control-Allow-Credentials":"true"
          },
          body: JSON.stringify({"magnetLink":text})
      }).then(res => res.json()).then(
          (result) => {
              console.log(result);
          }
      );
  };

  return(
    <InputArea>
        Magnet link: &nbsp;
        <StyledInput value={text} onChange={e => setText(e.target.value)}/> 
        <StyledButton type="button" onClick={getTorrent}>Add</StyledButton>
    </InputArea>
  );
};
export default Input