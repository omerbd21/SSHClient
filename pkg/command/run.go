package command

import (
	"VInstaller/pkg/hostkeys"
	"bytes"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
)


// The function runs a command on a remote linux host using ssh.
func RunCommand(ip string, command string ) string {
	hostKey, err := hostkeys.GetHostKey(ip)
	key, err := ioutil.ReadFile(os.Getenv("HOME")+"\\.ssh\\id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	// Connect to ssh server
	conn, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer conn.Close()
	// Create a session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()
	var output bytes.Buffer
	session.Stdout = os.Stdout
	session.Stdout = &output
	session.Stdin = os.Stdin
	err = session.Run(command)
	if err != nil{
		log.Fatal(err)
	}
	return output.String()


}
