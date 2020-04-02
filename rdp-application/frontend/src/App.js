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
  const [machineInfo, setMachineInfo] = React.useState("")
  const [connected, setConnected] = React.useState(false)
  const [localPort] = React.useState(process.env.REACT_APP_LOCAL_PORT ? process.env.REACT_APP_LOCAL_PORT : "43389")

  const [address, setAddress] = React.useState("")
  const [machineID, setMachineID] = React.useState("")
  const [withTotp, setWithTotp] = React.useState("")

  window.wails.Events.On("ConnectionSucceed", _ => {
    setConnected(true)
    setBusy(false)
  })
  window.wails.Events.On("ConnectionError", _ => {
    setConnected(false)
    setBusy(false)
    setError("Timeout connecting to server")
  })

  React.useEffect(() => {
    setBusy(true)
    window.backend.doLoadConfig().then((ret) => {
      if (!ret.error) {
        setAddress(ret.address)
        setMachineID(ret.machineID)
      } else {
        setError(ret.error)
      }
      setBusy(false)
    }).catch(setError)
  }, []);

  const enroll = (address, machineID, totpCode) => {
    setBusy(true)
    setAddress(address)
    setMachineID(machineID)
    if (totpCode) {
      window.backend.doConnectTotp(address, machineID, totpCode).then(ret => {
        if (!ret.error) {
          setMachineInfo(ret)
          setError(false)
        } else {
          setMachineInfo(null)
          if (ret["error"].indexOf("such host") !== -1 || ret["error"].indexOf("refused") !== -1) {
            setError("Unreachable host")
          } else if (ret["error"].indexOf("404") !== -1) {
            setError(t("Machine not found with ID") + " " + machineID)
          } else {
            setError(ret["error"])
          }
          setBusy(false)
        }
      }).catch(setError)
    }else{
      window.backend.doRegister(address, machineID).then(ret => {
        if (!ret.error) {
          if(ret.withTotp){
            setWithTotp(true)
            setError(false)
            setBusy(false)
          }else{
            setMachineInfo(ret)
            setError(false)
          }
        } else {
          setMachineInfo(null)
          if (ret["error"].indexOf("such host") !== -1 || ret["error"].indexOf("refused") !== -1) {
            setError("Unreachable host")
          } else if (ret["error"].indexOf("404") !== -1) {
            setError(t("Machine not found with ID") + " " + machineID)
          } else {
            setError(ret["error"])
          }
          setBusy(false)
        }
      }).catch(setError)
    }
  }

  const tryAgain = () => {
    setError(false)
    setAddress("")
    setMachineInfo(null)
    setWithTotp(false)
  }
  return {
    error,
    busy,
    machineInfo,
    localPort,
    connected,
    address,
    machineID,
    enroll,
    tryAgain,
    withTotp
  }
}

function App() {

  const { t } = useTranslation()

  const {
    error,
    busy,
    machineInfo,
    localPort,
    connected,
    address,
    machineID,
    enroll,
    tryAgain,
    withTotp,
  } = AppModel()

  const remoteMachineAddress = machineInfo ? "127.0.0.1:" + localPort : ""

  return (
    <div id="app" className="App">
      <div>
        <img src={logo} alt="BB" className="logo" />
      </div>
      {!busy && <>
        <div className="content-area">
          {!machineInfo && !error && !connected && <EnrollmentForm enroll={enroll} defaultAddress={address} defaultMachineID={machineID} defaultWithTotp={withTotp}/>}

          {machineInfo && !error && connected && <div className="machineid-area">
            <div>{t("Successfully connected")}</div>
            <h1>{remoteMachineAddress} <img src={copyIcon} alt="Copy" onClick={_ => copyTextToClipboard(remoteMachineAddress)} className="copy-button" title={t("Copy to clipboard")} /></h1>
            <div>{t("Open your Remote Desktop Application and use the following address to connect")}: {remoteMachineAddress}</div>
          </div>}

          {error && <div className="machineid-area">
            <h1>{t("Error registering")}</h1>
            <div>{t(error)}</div>
            <div style={{ marginTop: "1rem" }}><Button onClick={_ => tryAgain()} variant="outlined">{t("Try again")}</Button></div>
          </div>}
        </div>
      </>}
      {busy && <>
        {!machineInfo && <div className="loading-area">
          <h1>{t("Getting Account Information")}...</h1>
          <img src={loadingIcon} alt="Loading" className="loadingIcon" />
        </div>}

        {machineInfo && !connected && <div className="loading-area">
          <h1>{t("Connecting to RDP gateway")}...</h1>
          <img src={loadingIcon} alt="Loading" className="loadingIcon" />
        </div>}

      </>}
      <footer>
        {t("Handcrafted with")} <img src={iconHeart} alt="Love" className="heartIcon" /> by <img src={logoLabbs} alt="Labbs" className="logoFooter" />
      </footer>
    </div>
  );
}

export default App;
