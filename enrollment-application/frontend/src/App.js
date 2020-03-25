import React from 'react';

import { useTranslation } from 'react-i18next';

import { copyTextToClipboard } from "./helpers";

import logo from './logo.png';
import logoLabbs from './logo-labbs.png';
import iconHeart from './iconHeart.png';
import loadingIcon from './loading.gif'
import copyIcon from './copy.png';

import { Button } from "@material-ui/core";
import EnrollmentForm from "./components/EnrollmentForm";
import './App.css';

function AppModel() {

  const { t } = useTranslation()

  const [error, setError] = React.useState("")
  const [busy, setBusy] = React.useState(false)
  const [machineId, setMachineId] = React.useState("")
  const [address, setAddress] = React.useState("")
  const [accountID, setAccountID] = React.useState("")
  const [connected, setConnected] = React.useState(false)
  const [autoSubmit, setAutoSubmit] = React.useState(false)

  React.useEffect(() => {
    setBusy(true)
    window.backend.doLoadConfig().then((ret) => {
      if (!ret.error) {
        setAddress(ret.address)
        setAccountID(ret.accountID)
      } else {
        setError(ret.error)
      }
      setBusy(false)
    }).catch(setError)
  }, []);

  window.wails.Events.On("ConnectionSucceed", _ => {
    setConnected(true)
    setError(false)
    setBusy(false)
  })
  window.wails.Events.On("ConnectionError", _ => {
    setConnected(false)
    setBusy(false)
    setError("Timeout connecting to server")
  })

  const enroll = (address, accountID) => {
    setBusy(true)
    setAddress(address)
    setAccountID(accountID)
    window.backend.doRegister(address, accountID).then(ret => {
      if (!ret.error) {
        setMachineId(ret.machineId)
        setError(false)
      } else {
        setMachineId("")

        if (ret["error"].indexOf("such host") !== -1 || ret["error"].indexOf("refused") !== -1) {
          setError("Unreachable host")
        } else if (ret["error"].indexOf("404") !== -1) {
          setError(t("Company not found with ID") + " " + accountID)
        } else {
          setError(ret["error"])
        }
      }
    });
  }

  const tryAgain = () => {
    setAutoSubmit(true)
    setError(false)
    setMachineId("")
    setAddress("")
    setAccountID("")
  }

  return {
    connected,
    error, setError,
    busy,
    machineId,
    address,
    accountID,
    enroll,
    autoSubmit,
    tryAgain
  }
}

function App() {

  const { t } = useTranslation()

  const {
    error,
    connected,
    busy,
    machineId,
    accountID,
    address,
    enroll,
    autoSubmit,
    tryAgain
  } = AppModel()

  return (
    <div id="app" className="App">
      <div>
        <img src={logo} alt="BB" className="logo" />
      </div>
      {!busy && <>
        <div className="content-area">

          {!machineId && !error && !connected && <EnrollmentForm enroll={enroll} autoSubmit={autoSubmit} defaultAddress={address} defaultAccountID={accountID} />}

          {machineId && !error && connected && <div className="machineid-area">
            <div>{t("Machine successfully registered")}</div>
            <h1>{machineId} <img src={copyIcon} alt="Copy" onClick={_ => copyTextToClipboard(machineId)} className="copy-button" title={t("Copy to clipboard")} /></h1>
            <div>{t("Take a picture or write down a note of the above code because it will be requested when remotely accessing this machine")}</div>
          </div>}

          {error && <div className="machineid-area">
            <h1>{t("Error registering")}</h1>
            <div>{t(error)}</div>
            <div style={{ marginTop: "1rem" }}><Button onClick={_ => tryAgain()} variant="outlined">{t("Try again")}</Button></div>
          </div>}
        </div>
      </>}
      {busy && <>
        <div className="loading-area">
          <h1>{t("Your machine is being registered")}...</h1>
          <img src={loadingIcon} alt="Loading" className="loadingIcon" />
        </div>
      </>}
      <footer>
        {t("Handcrafted with")} <img src={iconHeart} alt="Love" className="heartIcon" /> by <img src={logoLabbs} alt="Labbs" className="logoFooter" />
      </footer>
    </div>
  );
}

export default App;
