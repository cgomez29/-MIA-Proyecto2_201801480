import React from 'react';

const Home = (props: {name: string}) =>{
    return (
        <React.Fragment>
            {props.name ? 'Hi ' + props.name : 'You are not logged in'}
        </React.Fragment>
    );
};

export default Home;