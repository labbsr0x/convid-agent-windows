import React from 'react';

import { useTranslation } from 'react-i18next';

import logo from './logo.png';
import logoBB from './logo-bb.png';
import iconHeart from './iconHeart.png';

import EnrollmentForm from "./components/EnrollmentForm";
import './App.css';

function App() {

  const { t } = useTranslation()

  return (
    <div id="app" className="App">
      <div>
        <img src={logo} alt="BB" className="logo" />
      </div>
      <div>
        <EnrollmentForm />
      </div>
      <footer>
        {t("Handcrafted with")} <img src={iconHeart} alt="Love" className="heartIcon" /> by <img src={logoBB} alt="Banco do Brasil" className="logoFooter" />
      </footer>
    </div>
  );
}

export default App;
