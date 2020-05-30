import React from 'react';
import {ToggleBtn, ToggleStyle, Image} from "./styles";

const Toggle = ({setShowList, showList, setIsListExpanded}) => {

    const listClick = () => {
        setShowList(true);
        setIsListExpanded(false);
      };
    const cardClick = () => {
        setShowList(false);
        setIsListExpanded(true);
    };

    return (
        <ToggleStyle>
            <ToggleBtn onClick={() => cardClick()} selected={!showList}><Image src="./grid.png" selected={!showList} /></ToggleBtn>
            <ToggleBtn onClick={() => listClick()} selected={showList}isListBtn={true}><Image src="./iconListBlack.png" selected={showList} /></ToggleBtn>
        </ToggleStyle>
    )
};

export default Toggle