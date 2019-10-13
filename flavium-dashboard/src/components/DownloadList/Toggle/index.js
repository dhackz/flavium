import React from 'react';
import {ToggleBtn, ToggleStyle, Image} from "./styles";

const Toggle = ({setShowList, showList}) => {

    const listClick = () => {
        setShowList(true);
      };
    const cardClick = () => {
        setShowList(false);
    };

    return (
        <ToggleStyle>
            <ToggleBtn onClick={() => cardClick()} selected={!showList}><Image src="./grid.png" selected={!showList} /></ToggleBtn>
            <ToggleBtn onClick={() => listClick()} selected={showList}isListBtn={true}><Image src="./iconListBlack.png" selected={showList} /></ToggleBtn>
        </ToggleStyle>
    )
};

export default Toggle