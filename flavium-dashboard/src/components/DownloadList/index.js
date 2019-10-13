import React, {useState, useEffect} from 'react';
import { ListStyle, LargeText, Header} from "./styles"
import Item from "./Item"
import ItemColumns from "./ItemColumns"
import Toggle from "./Toggle"
import downloads from "./mockDownloads.json"

const DownloadList = () => {
  
  const [showList, setShowList] = useState(false);

  const [currentDownloads, setDownloads] = useState([]);
  
  useEffect(() => {
    //TODO: Get real current downloads instead of mock data
    setDownloads(downloads.downloads);
  }, [])
  

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
      <ListStyle showList={showList}>
      
          {currentDownloads.map((item,key) => {
              const { magnetLink } = item;

              return (
                <Item
                  magnetLink={magnetLink}
                  showList={showList}
                  key={key}
                />
              );
            })}
      </ListStyle>
    </div>
  )
};

export default DownloadList