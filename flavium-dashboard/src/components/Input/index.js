import React, { useState }  from 'react';
import { StyledInput, InputArea, StyledButton } from "./styles"

const Input = () => {

  const [text, setText] = useState("");

  const getTorrent = () => {
    //TODO: Send request to start downloading the magnetlink in the inputfield
  }

  return(
    <InputArea>
        Magnet link: &nbsp;
        <StyledInput value={text} onChange={e => setText(e.target.value)}/> 
        <StyledButton type="button" onClick={getTorrent}>Add</StyledButton>
    </InputArea>
  );
};
export default Input