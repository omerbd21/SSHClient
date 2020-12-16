import React, {useCallback, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from "@material-ui/core/Button";
import sendCommand from "../api/sendCommand";

const useStyles = makeStyles((theme) => ({
    root: {
        '& > *': {
            margin: theme.spacing(1),
            width: '25ch',
        },
    },
    formGroup: {
        display: 'flex',
        justifyContent:'center',
        alignItems:'center',
        height: '100vh',
    }
}));

export default function BasicTextFields() {
    const classes = useStyles();
    const [ip, setIp] = useState("");
    const [command, setCommand] = useState("");
    const [output, setOutput] = useState("");
    const [isSending, setIsSending] = useState(false);
    const sendRequest = useCallback(async () => {
        if (isSending)
            return
        setIsSending(true)
        await sendCommand().then(result => setOutput(result))
        setIsSending(false)
    }, [isSending])
    return (
        <form className={classes.root} noValidate autoComplete="off">
            <span>
                <div className={classes.formGroup}>
                <TextField onChange={(event) => setIp(event.target.value)}
                           value={ip} id="standard-basic" label="IP" />
                <TextField onChange={(event) => setCommand(event.target.value)}
                           value={command} id="standard-basic" label="Command" />
                <Button onClick={sendRequest} variant="contained" color="primary">
                    Primary
                </Button>
                    <h2>{output ? output["output"]: "Loading"}</h2>
            </div>
            </span>

        </form>
    );
}