import React, {useState, useEffect} from 'react';
import { ItemStyle, Name, Percentage, Bottom, ProgressBar, Size } from "./styles"

const parseTorrent = require('parse-torrent')

const Item = ({showList, magnetLink})  => {

    useEffect(()=> {
        getTorrentInfo(magnetLink)
    }, [])

  const [name, setName] = useState("It: Chapter 2");

  let size = 700.0;
  let doneSize = 254.2;
  let status = "Downloading";
  let started = "2019-10-10";

  const getTorrentInfo = (link) => {
    const torrentData = parseTorrent(link)
    const regex = /(.*) ([12][09]\d\d)[ \n]/;
    let stringVal = torrentData.name
    stringVal = stringVal.replace(/\./g,' ')
    stringVal = stringVal.match(regex)[0];
    setName(stringVal)
  }

  return (
      <ItemStyle showList={showList}>
          <Name showList={showList}>{name}</Name>
          {/* <Size>{size.toFixed(1)}MB</Size>*/}
          <Bottom showList={showList}>
            <div>{status}</div>
            <Percentage>({((doneSize/size)*100).toFixed(2)}%)</Percentage>
            <ProgressBar percent={((doneSize/size)*100).toFixed(2)}>
                <div></div>
            </ProgressBar>
          </Bottom>
      </ItemStyle>
  );
};

export default Item