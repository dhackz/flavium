import React from 'react';
import {DetailsStyle, Title, Grid, Description, Trailer} from "./styles"

import ReactPlayer from 'react-player'

const Details = ({name, description, youtubeId, voteAverage}) => {

    const url = "https://www.youtube.com/watch?v=" + youtubeId
    return (
        <DetailsStyle>
            <Grid>
                <Trailer>
                    <ReactPlayer width="500px" height="350px" url={url} controls={true} playing={false}/>
                </Trailer>
                <div>
                    <Description>
                    <Title>{name}</Title>
                        {description}
                    </Description>
                </div>  
            </Grid>
        </DetailsStyle>
    )
};

export default Details