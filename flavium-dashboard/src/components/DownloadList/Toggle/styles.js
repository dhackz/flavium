import styled from "styled-components";


export const ToggleBtn = styled.div`
    display: inline-block;
    margin-right: 30px;
    padding: 20px;
    margin-top 5px;
    border:  ${props => props.selected ? "1px solid lightgreen" : "none"};
    color:  ${props => props.selected ? "lightgreen" : "white"};
`;




//probably wont need these
export const ListBtn = styled.div`
    display: inline-block;
    background: ${props => props.selected ? "white" : "palevioletred"};
`;

export const CardBtn = styled.div`
margin-left: 20px;
    background: ${props => props.selected ? "white" : "palevioletred"};
    display: inline-block;
`;
