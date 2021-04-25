import React from 'react';
import CardG from '../components/CardG'

const Home = (props: {name: string}) =>{
    let menu
    if(props.name) {
        menu = (
            <React.Fragment >
                <CardG/>
            </React.Fragment>
        )
    } else {
        menu = (
            <div>'You are not logged in'</div>
        )
    }


    return (
        <React.Fragment>
            { menu }
        </React.Fragment>
    );
};

export default Home;