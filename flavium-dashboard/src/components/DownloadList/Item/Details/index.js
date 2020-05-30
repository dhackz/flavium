import React from 'react';
import {DetailsStyle, Title, Grid, Description, Trailer, DescriptionStyle, Score, HorizontalDiv, Genres, Runtime, Budget, Revenue} from "./styles"

import ReactPlayer from 'react-player'

const Details = ({name, description, youtubeId, voteAverage, genres, releaseDate, tagline, budget, revenue, runtime}) => {

    const url = "https://www.youtube.com/watch?v=" + youtubeId
    return (
        <DetailsStyle>
            <Grid>
                <Trailer>
                    <ReactPlayer width="500px" height="350px" url={url} controls={true} playing={false}/>
                </Trailer>
                <Description>
                    <HorizontalDiv>
                        <Title><div>{name} ({releaseDate})</div></Title>
                        <Score>{"IMDb: " + voteAverage}</Score>
                    </HorizontalDiv>
                    <HorizontalDiv>
                        <Genres>{genres.map((genre) => (genre + " "))}</Genres>
                        <Runtime>{runtime} min</Runtime>
                    </HorizontalDiv>
                        <div>Result: ${(revenue-budget).toLocaleString()} <Revenue>(+${revenue.toLocaleString()})</Revenue>/<Budget>(-${budget.toLocaleString()})</Budget></div>
                    <DescriptionStyle>{description}</DescriptionStyle>
                </Description>
            </Grid>
        </DetailsStyle>
    )
};

export default Details