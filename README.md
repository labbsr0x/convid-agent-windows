# Convid Agente for Windows

Windows Agent GUI for Convid remote access

## Exposing the Machine

To run, first install **wails** following the instructions at https://wails.app/

To bring the SSH Test server up, run `docker-compose up`

After installing wails cli above, go to the folder **enrollment-application** and run `wails serve` to start the Go part of the application
Then go to the **enrollment-application/frontend** folder and num `npm start` to run the React part.
The browser will open with the React Application "tied" to the Go backend.

You also need an Accounts server up. You will put the Account service URL in the Application, on the "Server Address" field on the main screen.

To test, you put the Accounts server URL and a Company Identification that already exists on the Accounts database and click the Button. It should invoke the Accounts API to get the machine identification and SSH configuration and then try to connect to the SSH.



