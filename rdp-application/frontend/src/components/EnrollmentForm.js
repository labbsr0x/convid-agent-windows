import React, { useEffect } from 'react';

import { Button, TextField } from "@material-ui/core";

import { useTranslation } from 'react-i18next';


// The Model is defined here and not in a separeted file because of
// babel issues that I couldn't fix (Tiago Stutz)
function EnrollmentFormModel(defaultAddress, defaultMachineID, defaultWithTotp) {

    const [sealed] = React.useState(process.env.REACT_APP_SEALED)
    const [submitEnabled, setSubmitEnabled] = React.useState(false)
    const [serverAddress, setServerAddress] = React.useState(defaultAddress || process.env.REACT_APP_SERVER_ADDRESS)
    const [machineID, setMachineID] = React.useState(defaultMachineID)
    const [totpCode, setTotpCode] = React.useState("")
    const [withTotp] = React.useState(defaultWithTotp)

    useEffect(_ => {
        setSubmitEnabled(serverAddress && machineID && (withTotp ? totpCode : true))
    }, [serverAddress, machineID, withTotp, totpCode])

    return {
        sealed,
        submitEnabled,
        serverAddress, setServerAddress,
        machineID, setMachineID,
        totpCode, setTotpCode,
        withTotp,
    }
}

function EnrollmentForm({ enroll, defaultAddress, defaultMachineID, defaultWithTotp }) {

    const { t } = useTranslation();

    const {
        sealed,
        submitEnabled,
        serverAddress, setServerAddress,
        machineID, setMachineID,
        totpCode, setTotpCode,
        withTotp,
    } = EnrollmentFormModel(defaultAddress, defaultMachineID, defaultWithTotp)

    const onEnroll = () => {
        if (serverAddress && machineID && (withTotp ? totpCode : true)) {
            enroll(serverAddress, machineID, totpCode)
        }
    }

    return <>
        <div>
            <div name="form">
                <div>
                    {!sealed && <TextField style={{ width: "300px" }} defaultValue={serverAddress} label={t("Accounts Server Address")} onChange={e => setServerAddress(e.target.value)} variant="standard" />}
                </div>
                <div>
                    <TextField style={{ width: "300px" }} defaultValue={machineID} type="text" onChange={e => setMachineID(e.target.value.trim())} label={t("Remote Machine Code")} variant="standard" />
                </div>
                <div>
                    {withTotp && <TextField style={{ width: "300px" }} defaultValue={totpCode} type="text" onChange={e => setTotpCode(e.target.value.trim())} label={t("TOTP Code")} variant="standard" />}
                </div>
                <div className="button-area">
                    <Button variant="outlined" disabled={!submitEnabled} onClick={_ => onEnroll()}>{t("Connect to remote computer")}</Button>
                </div>
            </div>

        </div>
    </>
}


export default EnrollmentForm