import React, {useCallback, useState} from 'react';
import {makeStyles} from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from "@material-ui/core/Button";
import sendCommand from "../api/sendCommand";
import Paper from "@material-ui/core/Paper";

const useStyles = makeStyles((theme) => ({
    formGroup: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
    },
    paper: {
        height: "30vh",
        width: "100vh",
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
}));

export default function BasicTextFields() {
    const classes = useStyles();
    const [ip, setIp] = useState("");
    const [command, setCommand] = useState("");
    const [output, setOutput] = useState("");
    const [isSending, setIsSending] = useState(false);
    const sendRequest = useCallback(async () => {
        if (ip === "" || command ==="") {
            alert("You didn't enter both values")
        }
        else {
            if (isSending)
                return
            setIsSending(true)
            await sendCommand(ip, command).then(result => setOutput(result))
            setIsSending(false)
        }
    }, [isSending, command, ip])
    return (
            <div className={classes.formGroup}>
                <div><Paper className={classes.paper}>
                    <div><TextField onChange={(event) => setIp(event.target.value)}
                                     value={ip} id="standard-basic" label="IP"/>
                        <TextField onChange={(event) => setCommand(event.target.value)}
                                   value={command} id="standard-basic" label="Command"/></div>
                    <div><Button onClick={sendRequest} variant="contained" color="primary">
                        Get Output
                    </Button></div>
                </Paper></div>
                <div><Paper className={classes.paper}>
                    <h1>{ip}:</h1>
                    {output ? output["output"] : "Loading"}</Paper></div>
            </div>
    );
}