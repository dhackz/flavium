import styled from "styled-components"
export const ItemContainer = styled.div`
    background: ${props => props.showList ? "" : "linear-gradient(black, rgba(0,0,0,0.2), rgba(0,0,0,0.2), black);"}
    height: ${props => props.showList ? "60px" : "320px"}
    display: ${props => props.showList ? "grid" : "block"}
    grid-template-columns: ${props => props.showList ? "50% 50%" : "100%"};
    background-repeat: no-repeat;
`;

export const ItemStyle = styled.li`
    color: white;
    list-style: none;
    background: rgba(0, 0, 0, 0.2);
    background-size: cover;
    ${props => { if(props.showList) { return(
        '&:nth-child(odd){background: rgba(0, 0, 0, 0.3)};'+
        'padding: 20px;'+
        'box-shadow: none;'+
        'height: 60px;'
    )} else { return(
        'background-image: url(' + props.posterSrc+ ');'+
        'padding: 0;'+
        'box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);'+
        'height: 320px;'
    )}}}
    transition: all .2s ease-in-out;
    &:hover{
        ${props => props.showList ? "" : "transform: scale(1.1);"}
    }
`;

export const ProgressBar = styled.div`
    margin-bottom: 10px;
    margin-left: 10px;
    margin-right: 10px;
    margin-top: 5px;

    background-color: black;
    border-radius: 13px;
    /* (height of inner div) / 2 + padding */
    padding: 4px;

    div{
        background-color: orange;
        width: ${props => props.percent}%;
        /* Adjust with JavaScript */
        height: 12px;
        border-radius: 10px;
    }
`;

export const Name = styled.div`
    font-weight: 900;
    padding-top:  ${props => props.showList ? "24px" : "10px"};
    padding-left: 10px;
    overflow: hidden;
    white-space: pre-wrap; /* css-3 */
    white-space: -moz-pre-wrap !important; /* Mozilla, since 1999 */
`;


export const Size = styled.div`
    font-weight: 900;
    padding-left: 10px;
`;
export const Percentage = styled.div`
    font-weight: 900;
`;

export const Bottom = styled.div`
    position:  ${props => props.showList ? "static" : "relative"};
    bottom:  ${props => props.showList ? "0px" : "-200px"};

`;