import styled from "styled-components"

export const ItemStyle = styled.li`
    box-sizing:border-box;
    padding: 10px;
    list-style: none;
    display: grid;
    grid-template-columns: 25% 25% 25% 25%;
    color: white;
    width: 100%;
    background: rgba(0, 0, 0, 0.2);
    &:nth-child(odd){
        background: rgba(0, 0, 0, 0.3);
    }
      
`;