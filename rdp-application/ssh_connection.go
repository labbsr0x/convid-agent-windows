package main

import (
	"math/rand"
	"time"

	client "github.com/jairsjunior/go-ssh-client-tunnel/clientv2"
	"github.com/sirupsen/logrus"
)

func connect(sshServerHost string, sshServerPort int, user string, password string, localRDPHost string, localRDPPort int, tunneltoHost string, tunneltoPort int) {
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

	agentInstance.runtime.Events.Emit("ConnectionSucceed")
	for {
		err := client.CreateConnectionLocalV2(user, password, localRDPEndpoint, tunneltoEndpoint, sshServerEndpoint)
		logrus.Infof("===>>>> CONNECTD: %s", err)
		if err != nil {
			agentInstance.runtime.Events.Emit("ConnectionError")
			r := rand.Intn(10)
			time.Sleep(time.Duration(r) * time.Microsecond)
			logrus.Warningf("Error connecting to SSH... retrying in %d seconds.", r)
			continue
		}
		break
	}
}
