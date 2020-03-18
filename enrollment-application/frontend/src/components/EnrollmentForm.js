import React from 'react';

import { Button, TextField } from "@material-ui/core";

import { useTranslation } from 'react-i18next';


// The Model is defined here and not in a separeted file because of
// babel issues that I couldn't fix (Tiago Stutz)
function EnrollmentFormModel() {

    const [serverAddress, setServerAddress] = React.useState("")
    const [accountEmail, setAccountEmail] = React.useState("")

    return {
        serverAddress, setServerAddress,
        accountEmail, setAccountEmail,
    }
}

function EnrollmentForm({ enroll }) {

    const { t } = useTranslation();

    const {
        serverAddress, setServerAddress,
        accountEmail, setAccountEmail
    } = EnrollmentFormModel()

    return <>
        <div>

            <div name="form">
                <div>
                    <TextField defaultValue={serverAddress} label={t("Server Address")} onChange={e => setServerAddress(e.target.value)} variant="standard" />
                </div>
                <div>
                    <TextField defaultValue={accountEmail} type="email" onChange={e => setAccountEmail(e.target.value)} label={t("Account E-mail")} variant="standard" />
                </div>
                <div className="button-area">
                    <Button variant="outlined" onClick={_ => enroll(serverAddress, accountEmail)}>{t("Register this machine")}</Button>
                </div>
            </div>

        </div>
    </>
}

export default EnrollmentForm