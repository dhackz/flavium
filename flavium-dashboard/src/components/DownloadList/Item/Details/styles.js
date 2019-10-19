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

export const Title = styled.div`
    display: grid;
    grid-template-rows: 50% 50;
    grid-template-columns: none;
    text-align: left;
    font-size: 22px;
    margin-bottom:15px;
    display:inline-block;

    margin-right: auto;
    margin-left: 0;
`;

export const DescriptionStyle = styled.div`
    margin-top: 20px;
`;
export const Grid = styled.div`
    display: grid;
    grid-template-rows: 50% 50;
    grid-template-columns: none;
    margin-left: 20px;

    @media (min-width: 1200px){
        margin-left: 0;
        grid-template-rows: none;
        grid-template-columns: 50% 50%;
    }
`;
export const Description = styled.div`
    height: 220px;
    text-align: left;
    margin-top:0;
    margin-left: 10px;
    margin-right: 10px;
    @media (min-width: 1200px){
        margin-top:70px;
        margin-right:100px;
        margin-left: 0;
    }
`;

export const Trailer = styled.div`
    margin-top: 20px;
    margin-bottom: 20px;
    display:inline-block;
    
    @media (min-width: 1200px){
        margin-left: 15%;
    }
`;

export const Score = styled.div`
    margin-right: 0;
    margin-left: auto;
    font-size: 20px;
`;
export const Date = styled.div`
    margin-right: 0;
    margin-left: auto;
`;

export const Runtime = styled.div`
    margin-right: 0;
    margin-left: auto;
    font-size: 18px;
`;
export const Genres = styled.div`
    margin-right: auto;
    margin-left: 0;
    padding-bottom:10px;
    font-size: 18px;
`;

export const Budget = styled.span`
    color: red;
    margin-right: auto;
    margin-left: 0;
`;
export const Revenue = styled.span`
    color:lawngreen;
    margin-right: 0;
    margin-left: auto;
`;
export const HorizontalDiv = styled.div`
    display: flex;
`;