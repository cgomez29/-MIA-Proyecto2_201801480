import React from "react";
import yaml from "js-yaml";
import { FilePicker } from "react-file-picker";
import axios from 'axios';
import {Button, Grid, Paper} from "@material-ui/core";
import url from "../../config";
import Loading from '../../components/Loading'

export default class Bulkload extends React.Component {
    state = {
        loading: false
    }

    paperStyle = {padding: 25, height:'30vh', width:300, margin: "20px auto"}
    btnSumbitStyle = {margin: '15px 0'};

    load={};
    handleFileChange = file => {
        const reader = new FileReader();
        reader.readAsText(file);
        reader.onload = e => {
            try {
                const doc = yaml.load(e.target.result);
                this.load = JSON.stringify(doc);
                this.onSubmit();
            } catch (e) {
                console.log(e);
            }
        };

        this.setState({ title: file.name });
    };

    instance = axios.create({
        withCredentials: true,
    })

    onSubmit = async () => {
        this.setState({
            loading: true
        })
        await this.instance.post(`${url}/bulkload`
            ,this.load
            ,{headers: {
                    'Content-Type': 'application/json',
                }})
            .then(res => {
                console.log(res)
                this.setState({
                    loading: false
                })
            }).catch(err =>{
                console.log(err)
                this.setState({
                    loading: false
                })
            })
    };
    render() {
        if (this.state.loading)
            return <Loading/>

        return (
            <div className="container">
                <Paper elevation={10} style={this.paperStyle}>
                    <Grid align="center">
                        <h2> Carga Masiva </h2>
                    </Grid>
                    <FilePicker
                        extensions={["yaml"]}
                        onChange={this.handleFileChange}
                        onError={errMsg => console.log(errMsg)}
                    >
                        <Button type="submit"
                                variant="contained"
                                color="primary"
                                style={this.btnSumbitStyle}
                                fullWidth
                        >
                            Seleccionar archivo
                        </Button>
                    </FilePicker>
                </Paper>
            </div>
        )
    }

}