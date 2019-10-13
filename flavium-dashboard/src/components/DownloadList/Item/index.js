import React, {useState, useEffect} from 'react';
import { ItemStyle, Name, Percentage, Bottom, ProgressBar, ItemContainer} from "./styles"
require('dotenv').config()
const parseTorrent = require('parse-torrent')

const Item = ({showList, magnetLink})  => {

  useEffect(()=> {
    getTorrentInfo(magnetLink)
  }, [])

  const [name, setName] = useState("It: Chapter 2");
  const [posterSrc, setPosterSrc] = useState("");

  let size = 700.0;
  let doneSize = 254.2;
  let status = "Downloading";

  const getPoster = async(url) => {
    const api_call = await fetch(url);
    const data = await api_call.json();
    setPosterSrc("http://image.tmdb.org/t/p/w200//" +data.results[0].poster_path);
    }

  const getTorrentInfo = (link) => {
    const torrentData = parseTorrent(link)
    const regex = /(.*) ([12][09]\d\d)[ \n]/;
    let stringVal = torrentData.name
    stringVal = stringVal.replace(/\./g,' ')
    stringVal = stringVal.match(regex)[0];
    stringVal = stringVal.substring(0, stringVal.length - 1);
    var lastIndex = stringVal.lastIndexOf(" ");
    
    stringVal = stringVal.substring(0, lastIndex);

    const posterQuery = "https://api.themoviedb.org/3/search/movie?api_key="+ process.env.REACT_APP_MOVIE_KEY +"&query="+(stringVal.replace(/ /g,"%20"));
    getPoster(posterQuery)
    setName(stringVal)
  }

  return (
      <ItemStyle showList={showList} posterSrc={posterSrc}>
        <ItemContainer showList={showList}>
          <Name showList={showList}>{name}</Name>
          {/* <Size>{size.toFixed(1)}MB</Size>*/}
          <Bottom showList={showList}>
            <div>{status}</div>
            <Percentage>({((doneSize/size)*100).toFixed(2)}%)</Percentage>
            <ProgressBar percent={((doneSize/size)*100).toFixed(2)}>
                <div></div>
            </ProgressBar>
          </Bottom>
        </ItemContainer>
      </ItemStyle>
  );
};

export default Item