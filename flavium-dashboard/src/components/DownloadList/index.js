import React, {useState, useEffect} from 'react';
import { ListStyle, LargeText, Header} from "./styles"
import Item from "./Item"
import ItemColumns from "./ItemColumns"
import Toggle from "./Toggle"

const DownloadList = ({postListener}) => {
  
  const [showList, setShowList] = useState(false);
  const [currentDownloads, setDownloads] = useState([]);
  const [isListExpanded, setIsListExpanded] = useState(false);
  const [indexOfExpanded, setIndexOfExpanded] = useState(null);


  useEffect(() => {
      fetchData()
  },[postListener]);

  const fetchData = async () => {
      const result = await fetch("http://localhost:8080/v1/torrent", {
          method: 'GET',
          credentials: 'include',
          headers: {
              "Access-Control-Allow-Credentials":"true"
          },
      });
      const json = await result.json()
      await setDownloads(json.torrents);
  }

  let itemColumns = null;
  if(showList){
    itemColumns = <ItemColumns />
  }

  return (
    <div>
      <Header>
        <LargeText>Currently downloading:</LargeText> 
        <Toggle setShowList={setShowList} showList={showList}/>
      </Header>
      {itemColumns}
      {currentDownloads && <ListStyle showList={showList}>
          {currentDownloads.map((item,key) => {
              let isExpanded = false;
              if(isListExpanded && key===indexOfExpanded){
                isExpanded=true;
              }
              return (
                <Item
                  download={item}
                  showList={showList}
                  key={key}
                  setIsListExpanded={setIsListExpanded}
                  setIndexOfExpanded={setIndexOfExpanded}
                  isExpanded={isExpanded}
                  index={key}
                />
              );
            })}
      </ListStyle>}
    </div>
  )
};

export default DownloadList