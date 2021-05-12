import React, {useEffect, useState} from 'react';
import Paper from '@material-ui/core/Paper';
import Typography from "@material-ui/core/Typography";
import EmojiEventsIcon from '@material-ui/icons/EmojiEvents';
import AttachMoneyIcon from '@material-ui/icons/AttachMoney';
import url from "../../config";
import axios from "axios";

const Home = () => {
    const [data, setData] = useState({
        capital: 0,
        gold: 0,
        silver: 0,
        bronze: 0,
    })
    // styles
    const paperStyle = {padding: 25, maxWidth: 350, margin: "20px auto"}

    const instance = axios.create({
        withCredentials: true,
    })

    useEffect(()=> {
        (
            async () => {
                await instance.get(`${url}/home/detalle`)
                    .then(res => {
                        setData(res.data)
                    })
            }
        )();
    },[])

    return (
        <>
            <Paper elevation={10} style={paperStyle}>
                <Typography variant="body2" gutterBottom>
                    CAPITAL
                </Typography>
                <AttachMoneyIcon/> Q.{data.capital}
            </Paper>
            <Paper elevation={10} style={paperStyle}>
                <Typography variant="body2" gutterBottom>
                    GOLD
                </Typography>
                <EmojiEventsIcon/> {data.gold}
            </Paper>
            <Paper elevation={10} style={paperStyle}>
                <Typography variant="body2" gutterBottom>
                    SILVER
                </Typography>
                <EmojiEventsIcon/><EmojiEventsIcon/> {data.silver}
            </Paper>
            <Paper elevation={10} style={paperStyle}>
                <Typography variant="body2" gutterBottom>
                    BRONZE
                </Typography>
                <EmojiEventsIcon/><EmojiEventsIcon/><EmojiEventsIcon/> {data.bronze}
            </Paper>
        </>
    );
};

export default Home;