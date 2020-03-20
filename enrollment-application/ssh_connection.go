package main

import (
	client "github.com/jairsjunior/go-ssh-client-tunnel/clientv2"
	"github.com/sirupsen/logrus"
)

func serve(sshServerHost string, sshServerPort int, user string, password string, localRDPHost string, localRDPPort int, tunneltoHost string, tunneltoPort int) {
	localRDPEndpoint := client.Endpoint{
		Host: localRDPHost,
		Port: localRDPPort,
	}
	tunneltoEndpoint := client.Endpoint{
		Host: tunneltoHost,
		Port: tunneltoPort,
	}
	sshServerEndpoint := client.Endpoint{
		Host: sshServerHost,
		Port: sshServerPort,
	}
	logrus.Infof("Connecting to remote server: SSHServerHost: %s | SSHServerPort: %d | user: %s | password: %s", sshServerHost, sshServerPort, user, password)
	logrus.Infof("Local routing to: Host: %s | Port: %d", localRDPHost, localRDPPort)
	logrus.Infof("Tunneling to: Host: %s | Port: %d", tunneltoHost, tunneltoPort)

	err := client.CreateConnectionRemoteV2(user, password, localRDPEndpoint, tunneltoEndpoint, sshServerEndpoint)
	if err == nil {
		logrus.Errorf("Error at create connection: ", err)
	}
}
