import React from 'react';
import {ToggleBtn} from "./styles";

const Toggle = ({setShowList, showList}) => {

    const listClick = () => {
        setShowList(true);
      };
    const cardClick = () => {
        setShowList(false);
    };

    return (
        <>
            <ToggleBtn onClick={() => listClick()} selected={showList}>List</ToggleBtn>
            <ToggleBtn onClick={() => cardClick()} selected={!showList}>Cards</ToggleBtn>
        </>
    )
};

export default Toggle