import styled from "styled-components";

export const ToggleBtn = styled.div`
    display: inline-block;
    padding: 15px;
    color:  ${props => props.selected ? "lightgreen" : "white"};
    background-color: ${props => props.selected ? "white" : "rgba(0,0,0,0)"};
`;


export const ToggleStyle = styled.div`
    margin-top 10px;
    margin-bottom 6px;
    margin-right: 30px;
    display: inline-block;
    border: 1px solid white;
`;

export const Image = styled.img`
    width: 22px;
    height: 22px;
    filter: ${props => props.selected ? "" : "brightness(0) invert(1)"};
`;