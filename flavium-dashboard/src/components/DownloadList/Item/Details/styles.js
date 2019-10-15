import styled from "styled-components"
export const DetailsStyle = styled.div`

position: absolute;
display: inline-block;
top: auto;
left: 0;
width: 100%;
background-color: rgba(0,0,0, 0.34);
margin-top: 10px; 
`;

export const Title = styled.h2`
    text-align: left;
    margin-top: 30px;
`;

export const Grid = styled.div`
    display: grid;
    grid-template-rows: 50% 50;
    grid-template-columns: none;
    margin-left: 20px;

    @media (min-width: 768px){
        margin-left: 0;
        grid-template-rows: none;
        grid-template-columns: 50% 50%;
    }
`;
export const Description = styled.div`
    height: 220px;
    text-align: left;
    margin-top:0;
    @media (min-width: 768px){
        margin-top:100px;
        margin-right:100px;
    }
    
    margin-right:0;
`;

export const Trailer = styled.div`
    margin-top: 20px;
    margin-bottom: 20px;

    margin-left: 0;
    @media (min-width: 768px){
        margin-left: 15%;
    }
`;