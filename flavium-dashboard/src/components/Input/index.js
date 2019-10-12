import React, { useState }  from 'react';
import { StyledInput, InputArea, StyledButton } from "./styles"

const Input = () => {

  const parseTorrent = require('parse-torrent')

  const [text, setText] = useState("");

  const getTorrent = () => {
    const torrent = parseTorrent("magnet:?xt=urn:btih:"+text)
    console.log("torrent " + Object.getOwnPropertyNames(torrent))
  }

  return(
    <InputArea>
        InfoHash: &nbsp;
        <StyledInput value={text} onChange={e => setText(e.target.value)}/> 
        <StyledButton type="button" onClick={getTorrent}>Add</StyledButton>
    </InputArea>
  );
};
export default Input