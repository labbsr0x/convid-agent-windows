import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import pt_BR from './locales/pt-BR'
import en_US from './locales/en-US'

i18n
    .use(initReactI18next) // passes i18n down to react-i18next
    .init({
        resources: {
            en: en_US,
            pt: pt_BR
        },
        lng: "pt",

        // keySeparator: false, // we do not use keys in form messages.welcome

        interpolation: {
            escapeValue: false // react already safes from xss
        }
    });

export default i18n;