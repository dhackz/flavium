import styled from "styled-components"

export const ListStyle = styled.div`
    background-image: linear-gradient(to bottom right, #3f4a4a, #47392f);
    min-height: 1000px;
`;

export const ItemStyle = styled.li`
    box-sizing:border-box;
    list-style: none;
    
    display: grid;
    grid-template-columns: 25% 25% 25% 25%;
    color: white;
    width: 100%;
    padding-top: 16px;
    padding-bottom: 14px;
    padding: 10px;
    border-top: 1px solid rgba(35,42,42, 1);
    border-bottom: 1px solid rgba(35,42,42, 1);
    background-image: linear-gradient(to right, #404a4a, #43413c);
`;