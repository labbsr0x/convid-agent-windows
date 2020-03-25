import React, { useEffect } from 'react';

import { Button, TextField } from "@material-ui/core";

import { useTranslation } from 'react-i18next';


// The Model is defined here and not in a separeted file because of
// babel issues that I couldn't fix (Tiago Stutz)
function EnrollmentFormModel(defaultAddress, defaultAccountID) {

    const [sealed] = React.useState(process.env.REACT_APP_SEALED)
    const [submitEnabled, setSubmitEnabled] = React.useState(false)
    const [serverAddress, setServerAddress] = React.useState(defaultAddress || process.env.REACT_APP_SERVER_ADDRESS)
    const [accountId, setAccountId] = React.useState(defaultAccountID || process.env.REACT_APP_ACCOUNT_ID)

    useEffect(_ => {
        setSubmitEnabled(serverAddress && accountId)
    }, [serverAddress, accountId])

    return {
        sealed,
        submitEnabled,
        serverAddress, setServerAddress,
        accountId, setAccountId,
    }
}

function EnrollmentForm({ enroll, defaultAddress, defaultAccountID, autoSubmit }) {

    const { t } = useTranslation();

    const {
        sealed,
        submitEnabled,
        serverAddress, setServerAddress,
        accountId, setAccountId
    } = EnrollmentFormModel(defaultAddress, defaultAccountID)

    const onEnroll = () => {
        if (serverAddress && accountId) {
            enroll(serverAddress, accountId)
        }
    }

    if (autoSubmit) {
        onEnroll()
    }

    return <>
        <div>

            <div name="form">
                <div>
                    {!sealed && <TextField style={{ width: "300px" }} defaultValue={serverAddress} label={t("Server Address")} onChange={e => setServerAddress(e.target.value)} variant="standard" />}
                </div>
                <div>
                    {!sealed && <TextField style={{ width: "300px" }} defaultValue={accountId} type="text" onChange={e => setAccountId(e.target.value)} label={t("Organization Login")} variant="standard" />}
                </div>
                <div>
                    {sealed && <h3 style={{ width: "330px" }} >{t("Click the button below to enable it to be remotely accesed")}</h3>}
                </div>
                <div className="button-area">
                    <Button variant="outlined" disabled={!submitEnabled} onClick={_ => onEnroll()}>{t("Register this machine")}</Button>
                </div>
            </div>

        </div>
    </>
}


export default EnrollmentForm