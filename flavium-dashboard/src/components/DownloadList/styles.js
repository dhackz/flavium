import styled from "styled-components"

export const ListStyle = styled.div`
    padding: ${props => props.showList ? "0" : "20px"};
    background-image: linear-gradient(to bottom right, #3f4a4a, #47392f);
    display: ${props => props.showList ? "block" : "grid"};
    grid-template-columns: ${props => props.showList ? "25% 25% 25% 25%" : "repeat(auto-fill, minmax(20px, 200px))"};
    grid-column-gap: 20px;
    grid-row-gap: 20px;
    
`;

export const Header = styled.div`
    text-align: right;
    width: 100%;
    background-image: linear-gradient(to right, #404a4a, #43413c);
`;

export const LargeText = styled.div`
    text-align: left;
    padding: 20px;
    font-size: 22px;
    float:left;
`;

export const Status = styled.p`
    margin: 0;
    font-size: 1rem; font-weight: 400;
`;

export const Toggle = styled.div`
    float:right;
    margin: 0 0 1.5rem;

    input{
        visibility:hidden;
    }
    input + label {
        margin: 0; padding: .75rem 2rem; box-sizing: border-box;
        position: relative; display: inline-block;
        font-size: 1rem; line-height: 140%; font-weight: 600; text-align: center;
        
    }
    input:hover + label { background-color: green;}
    input:checked + label {
        background-color: red;
        color: #FFF;
        z-index: 1;
    }
    input:focus + label {@include focusOutline;}

    @include breakpoint(800) {
        input + label {
            padding: .75rem .25rem;
        }
    }
`;

export const Fieldset = styled.div`
    margin: 0; padding: 2rem; 
`;

export const ItemStyle = styled.li`
    box-sizing:border-box;
    list-style: none;
    
    display: grid;
    grid-template-columns: 50% 50%;
    color: white;
    padding-top: 16px;
    padding-bottom: 14px;
    padding: 10px;
    border-top: 1px solid rgba(35,42,42, 1);
    border-bottom: 1px solid rgba(35,42,42, 1);
    background-image: linear-gradient(to right, #404a4a, #43413c);
`;