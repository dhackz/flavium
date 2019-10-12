import styled from "styled-components"

export const InputArea = styled.div`
    color: white;
    background-image: linear-gradient(to right, #404a4a, #43413c);
    padding-top: 20px;
    padding-bottom: 20px;
`;

export const StyledInput = styled.input`
    vertical-align: middle;
    height: 30px;
    width: 30%;
    font-family: 'Raleway', sans-serif;
    padding-left: 10px;
    opacity: 0.9;
    &:hover,
    &:focus{
        opacity: 1;
    }
`;

export const StyledButton = styled.button`
    vertical-align: middle;
    height: 36px;
    width: 10%
    background-color: #cc7b19;
    border: none;
    color: white;
    font-family: 'Raleway', sans-serif;
    &:hover{
        background-color: #8a571b;
    }
`;