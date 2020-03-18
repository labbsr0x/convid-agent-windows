import React from 'react';

import { useTranslation } from 'react-i18next';

import logo from './logo.png';
import logoBB from './logo-bb.png';
import iconHeart from './iconHeart.png';
import loadingIcon from './loading.gif'
import copyIcon from './copy.png';

import EnrollmentForm from "./components/EnrollmentForm";
import './App.css';

function AppModel() {

  const [busy, setBusy] = React.useState(false)
  const [machineId, setMachinedId] = React.useState("")

  const enroll = () => {
    setBusy(true)
    setTimeout(_ => { setMachinedId("JAX99357"); setBusy(false) }, 3000)
  }
  return {
    busy,
    machineId,
    enroll
  }
}

function App() {

  const { t } = useTranslation()

  const {
    busy,
    machineId,
    enroll
  } = AppModel()

  return (
    <div id="app" className="App">
      <div>
        <img src={logo} alt="BB" className="logo" />
      </div>
      {!busy && <>
        <div>
          {!machineId && <EnrollmentForm enroll={enroll} />}
          {machineId && <div className="machineid-area">
            <div>{t("Machine successfully registered")}</div>
            <h1>{machineId} <img src={copyIcon} alt="Copy" className="copy-button" /></h1>
            <div>{t("Take a picture or write down a note of the above code because it will be requested when remotely accessing this machine")}</div>
          </div>}
        </div>
      </>}
      {busy && <>
        <div>
          <h1>{t("Your machine is being registered")}...</h1>
          <img src={loadingIcon} alt="Loading" className="loadingIcon" />
        </div>
      </>}
      <footer>
        {t("Handcrafted with")} <img src={iconHeart} alt="Love" className="heartIcon" /> by <img src={logoBB} alt="Banco do Brasil" className="logoFooter" />
      </footer>
    </div>
  );
}

export default App;
