# Convid Agente for Windows

Windows Agent GUI for Convid remote access

## Exposing the Machine

To run, first install **wails** following the instructions at https://wails.app/

To bring the SSH Test server up, run `docker-compose up`

- After installing **wails cli** above, go to the folder **enrollment-application** and run `wails serve` to start the GO part of the application
- Then go to the **enrollment-application/frontend** folder and num `npm start` to run the React part.
The browser will open with the React Application "tied" to the Go backend.

You also need an Accounts server up. You will put the Account service URL in the Application, on the "Accounts Server Address" field on the main screen.

To test, you put the Accounts server URL and a Company Identification that already exists on the Accounts database and click the Button. It should invoke the Accounts API to get the machine identification and SSH configuration and then try to connect to the SSH.



## Building and Packaging

### Enrollment Application (Gonzaga)

To package a instalation for an specific AccountID you can use the environment variables on the `wails build` command:
- `REACT_APP_ACCOUNT_ID`: AccountId related to the application
- `REACT_APP_SERVER_ADDRESS`: Accounts Backend API server
- `REACT_APP_SEALED`: so the application does not show Server Address on the main screen

**Example:** `REACT_APP_ACCOUNT_ID=12345678 REACT_APP_SERVER_ADDRESS=https://my-accounts-backend-server.com REACT_APP_SEALED=true wails build`

### RDP Application (Chiquinho)

To package a instalation for an specific AccountID you can use the environment variables on the `wails build` command:
- `REACT_APP_SERVER_ADDRESS`: Accounts Backend API server
- `REACT_APP_SEALED`: so the application does not show Server Address and AccountId fields on the main screen

**Example:** `REACT_APP_SERVER_ADDRESS=https://my-accounts-backend-server.com REACT_APP_SEALED=true wails build`