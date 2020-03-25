package main

import (
	"math"
	"math/rand"
	"time"

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
	logrus.Infof("Connecting to remote server: SSHServerHost: %s | SSHServerPort: %d | user: %s | password: %s", "bla.bla.bla", sshServerPort, user, password)
	logrus.Infof("Local routing to: Host: %s | Port: %d", localRDPHost, localRDPPort)
	logrus.Infof("Tunneling to: Host: %s | Port: %d", tunneltoHost, tunneltoPort)

	isConnected := make(chan bool)
	at := 0.0

	for {
		go client.CreateConnectionRemoteV2(user, password, localRDPEndpoint, tunneltoEndpoint, sshServerEndpoint, isConnected)

		v := <-isConnected

		if !v {

			if at < 4 {
				r := rand.Intn(10) + int(math.Exp2(at))
				time.Sleep(time.Duration(r) * time.Second)
				logrus.Warningf("Error connecting to SSH... retrying in %d seconds.", r)
				at++
				continue
			} else {
				logrus.Warning("Could not connect to SSH. Max attempts reached.")
				agentInstance.runtime.Events.Emit("ConnectionError")
				break
			}

		} else {
			agentInstance.runtime.Events.Emit("ConnectionSucceed")
			logrus.Infof("===>>>> CONNECTED")
			break
		}

	}
}
