import React from 'react';
import { ListStyle } from "./styles"
import Item from "./Item"
import ItemColumns from "./ItemColumns"

const DownloadList = () => (
    <div>
      <ItemColumns />
      <ListStyle>
          {/* TODO: Get current downloads */}
          <Item infoHash={"E1083C831FE2151DF607B109C42D10766AE9700B"} size={700.0} doneSize={254.2} status={"Downloading"} started={"2019-10-05"}/>
          <Item infoHash={"E1083C831FE2151DF607B109C42D10766AE9700B"} size={700.0} doneSize={254.2} status={"Downloading"} started={"2019-10-05"}/>
          <Item infoHash={"E1083C831FE2151DF607B109C42D10766AE9700B"} size={700.0} doneSize={254.2} status={"Downloading"} started={"2019-10-05"}/>
          <Item infoHash={"E1083C831FE2151DF607B109C42D10766AE9700B"} size={700.0} doneSize={254.2} status={"Downloading"} started={"2019-10-05"}/>
      </ListStyle>
    </div>
  );

export default DownloadList