import React from 'react';
import { ItemStyle } from "./styles"

const Item = ({infoHash, size, doneSize, status, started})  => {

  return (
      <ItemStyle>
          <div>{infoHash}</div>
          <div>{size.toFixed(1)}MB</div>
          <div>{status} ({((doneSize/size)*100).toFixed(2)}%)</div>
          <div>{started}</div>
      </ItemStyle>
  );
};

export default Item