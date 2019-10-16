import React, {useState, useEffect} from 'react';
import { ItemStyle, Name, Percentage, Bottom, ProgressBar, ItemContainer} from "./styles"
import Details from "./Details"
require('dotenv').config()

const Item = ({showList, download, setIsListExpanded, setIndexOfExpanded, isExpanded, index})  => {

  useEffect(()=> {
    let stringVal = download.Name;
    stringVal = stringVal.replace(/\./g,' ')
    stringVal = stringVal.substring(0, stringVal.length - 1);
    var lastIndex = stringVal.lastIndexOf(" ");
    
    stringVal = stringVal.substring(0, lastIndex);
    const getPoster = async(url) => {
      const api_call = await fetch(url);
      const data = await api_call.json();
      if(data.results){
        setPosterSrc("http://image.tmdb.org/t/p/w200//" +data.results[0].poster_path);
        setDescription(data.results[0].overview)
        setVoteAverage(data.results[0].vote_average)
        getTrailerLink(data.results[0].id);
      }
    }

    getPoster("https://api.themoviedb.org/3/search/movie?api_key="+ process.env.REACT_APP_MOVIE_KEY +"&query="+(stringVal.replace(/ /g,"%20")))
    
    setName(stringVal)
    setProgress(parseInt(download.Done.slice(0,download.Done.length-1)));
    setDoneSize(download.Have);
    setStatus(download.Status);
  }, [download])

  const [name, setName] = useState("It: Chapter 2");
  const [posterSrc, setPosterSrc] = useState("");
  const [description, setDescription] = useState("");
  const [youtubeId, setYoutubeId] = useState("");
  const [voteAverage, setVoteAverage] = useState(0.0);
  const [progress, setProgress] = useState(0);
  const [doneSize, setDoneSize] = useState("0 kb");
  const [status, setStatus] = useState("N/A");

  const getTrailerLink = async(movieId) => {
    const api_call = await fetch("http://api.themoviedb.org/3/movie/"+movieId+"/videos?api_key="+process.env.REACT_APP_MOVIE_KEY)
    const data = await api_call.json();
    setYoutubeId(data.results[0].key)
  }

  let details = null;
  const handleClick = () => {
    if(!isExpanded){
      setIsListExpanded(true)
      setIndexOfExpanded(index)
    }else{
      setIsListExpanded(false);
      setIndexOfExpanded(null)
    }
  }
  if(isExpanded){
    details = <Details name={name} description={description} youtubeId={youtubeId} voteAverage={voteAverage}/>;
  }

  return (
    <>
      <ItemStyle showList={showList} posterSrc={posterSrc} isExpanded={isExpanded} onClick={(event) => {handleClick(event)}}>
        <ItemContainer showList={showList} id="item-container">
          <Name showList={showList}>{name}</Name>
          {/* <Size>{size.toFixed(1)}MB</Size>*/}
          <Bottom showList={showList}>
            <div>{status}</div>
            <Percentage>({(progress).toFixed(2)}%, {doneSize})</Percentage>
            <ProgressBar percent={(progress).toFixed(2)}>
                <div></div>
            </ProgressBar>
          </Bottom>
        </ItemContainer>
          {details}
      </ItemStyle>
          </>
  );
};

export default Item