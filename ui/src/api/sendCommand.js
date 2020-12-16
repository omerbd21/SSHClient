async function sendCommand() {
   // Default options are marked with *
    return await fetch("http://localhost:8080/runcommand", {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        headers: {
            'Content-Type': 'application/json',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify({"ip": "192.168.2.148", "command": "ls"}) // body data type must match "Content-Type" header
    }).then( async response => await response.json())
}
export default sendCommand;