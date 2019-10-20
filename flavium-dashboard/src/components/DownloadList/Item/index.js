import React, {useState, useEffect} from 'react';
import { ItemStyle, Name, Percentage, Bottom, ProgressBar, ItemContainer} from "./styles"
import Details from "./Details"
require('dotenv').config()

const Item = ({showList, download, setIsListExpanded, setIndexOfExpanded, isExpanded, index})  => {

  useEffect(()=> {

    const fetchDataFromName = async (stringVal) => {
      setName(stringVal)
      const api_call = await fetch("https://api.themoviedb.org/3/search/movie?api_key="+ process.env.REACT_APP_MOVIE_KEY +"&query="+(stringVal.replace(/ /g,"%20")));
      const data = await api_call.json();
      if(data.results.length>0){
        setId(data.results[0].id)
        const api_call2 = await fetch("https://api.themoviedb.org/3/movie/"+data.results[0].id+"?api_key="+process.env.REACT_APP_MOVIE_KEY);
        const data2 = await api_call2.json();
        setPosterSrc("http://image.tmdb.org/t/p/w200//" +data2.poster_path)
        setDescription(data2.overview)
        setVoteAverage(data2.vote_average)
        getTrailerLink(data2.id);
        setReleaseDate(data2.release_date);
        setGenres(data2.genres.map(genre => genre.name));
        setBudget(data2.budget)
        setRevenue(data2.revenue)
        setRuntime(data2.runtime)
        setProgress(parseInt(download.Done.slice(0,download.Done.length-1)));
        setDoneSize(download.Have);
        setStatus(download.Status);
      }
    }
    if(download.Name !== undefined){
      const albinsBetterRegex = /^(([a-zA-Z]+) +(& +)?)+/;
      let stringVal = download.Name
      stringVal = stringVal.replace(/\./g,' ')
      stringVal = stringVal.match(albinsBetterRegex)[0];
      fetchDataFromName(stringVal)
    }
    
  }, [download])

  const [name, setName] = useState("It: Chapter 2");
  const [id, setId] = useState("")
  const [posterSrc, setPosterSrc] = useState("");
  const [youtubeId, setYoutubeId] = useState("");
  const [description, setDescription] = useState("");
  const [voteAverage, setVoteAverage] = useState(0.0);
  const [genres, setGenres] = useState([]);
  const [releaseDate, setReleaseDate] = useState("");
  const [budget, setBudget] = useState("");
  const [revenue, setRevenue] = useState("")
  const [runtime, setRuntime] = useState("")
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
    if(!showList){
      if(!isExpanded){
        setIsListExpanded(true)
        setIndexOfExpanded(index)
      }else{
        setIsListExpanded(false);
        setIndexOfExpanded(null)
      }
    }
  }
  if(isExpanded){
    details = <Details name={name} description={description} youtubeId={youtubeId} voteAverage={voteAverage} genres={genres} budget={budget}releaseDate={releaseDate} revenue={revenue} runtime={runtime}/>;
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