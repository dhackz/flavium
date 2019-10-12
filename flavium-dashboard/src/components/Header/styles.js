import styled from "styled-components"

export const StyledHeader = styled.div`
    width: 100%;
    background: #181c1c;
    color: white;
    text-align: left;
    height: 60px;
    background-image: linear-gradient(to right, #252c2c , #1e2424, #252c2c);
    font-size: 20px;
    line-height: 60px;
    padding-left: 20px;
    box-sizing:border-box;  /** so padding doesnt cause horizontal overflow **/
    -moz-box-sizing:border-box; 
    -webkit-box-sizing:border-box;
    -ms-box-sizing:border-box;
`;

export const StyledLink = styled.a`
    text-decoration:none;
    color: white;
    vertical-align: middle;
    margin-left: 100px;
`;

export const StyledLogo = styled.a`
    text-decoration:none;
    color: white;
    vertical-align: middle;
`;


export const StyledImage = styled.img`
    width: 40px;
    margin-bottom:3px;
    vertical-align: middle;
    margin-right: 10px;
    filter: invert(100%);
`;