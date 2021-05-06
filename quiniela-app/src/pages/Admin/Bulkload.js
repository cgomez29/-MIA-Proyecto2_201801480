import React from "react";
import yaml from "js-yaml";
import { FilePicker } from "react-file-picker";
import axios from 'axios';
import {Button, Grid, Paper} from "@material-ui/core";
import url from "../../config";


export default class Bulkload extends React.Component {
    paperStyle = {padding: 25, height:'30vh', width:300, margin: "20px auto"}
    btnSumbitStyle = {margin: '15px 0'};

    load={};
    handleFileChange = file => {
        const reader = new FileReader();
        reader.readAsText(file);
        reader.onload = e => {
            try {
                const doc = yaml.load(e.target.result);
                //this.load = JSON.stringify(doc);
                this.load = doc
                //console.log(this.load);
                //console.log(doc['A2'])
                this.onSubmit();
            } catch (e) {
                console.log(e);
            }
        };

        this.setState({ title: file.name });
    };
    onSubmit = async () => {
        console.log(JSON.stringify(this.load['A2']))

        /*await axios.post(`${url}/bulkload`, JSON.stringify(this.load['A2']))
            .then(res => {
                console.log(res)
            }).catch(err =>{
                console.log(err)
            })*/
    };
    render() {
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