import React, { useState }  from 'react';
import { StyledInput, InputArea, StyledButton } from "./styles"

const Input = ({onPost}) => {

  const [text, setText] = useState("");

  const postTorrent = () => {
      fetch("http://localhost:8080/v1/torrent", {
          method: 'POST',
          credentials: 'include',
          headers: {
              "Access-Control-Allow-Credentials":"true"
          },
          body: JSON.stringify({"magnetLink":text})
      }).then(res => res.json()).then(
          (result) => {
              onPost(result)
              setText("");
          }
      );
  };

  return(
    <InputArea>
        Magnet link: &nbsp;
        <StyledInput value={text} onChange={e => setText(e.target.value)}/> 
        <StyledButton type="button" onClick={postTorrent}>Add</StyledButton>
    </InputArea>
  );
};
export default Input