package sshclient

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	Conn *ssh.Client
}

func Connect(addr, user, keyPath string) (*Client, error) {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // TODO: Implement proper host key verification
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	return &Client{Conn: client}, nil
}

func (c *Client) Run(cmd string) (string, error) {
	session, err := c.Conn.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	return strings.TrimSpace(string(output)), err
}
