import React from 'react';
import { StyledHeader, StyledLogo, StyledLink, StyledImage } from "./styles"

const Header = () => (
    <StyledHeader>
        <StyledImage src={ require('./logo.png') } />
        <StyledLogo>Flavium</StyledLogo>
        <StyledLink href="https://app.plex.tv/">Browse</StyledLink>
    </StyledHeader>
  );

export default Header