import React from 'react';

import {Button, TextField} from "@material-ui/core";

import { useTranslation } from 'react-i18next';

function EnrollmentForm(){

    const { t } = useTranslation();


    return <>
        <div>
            <div name="form">
                <div>
                    <TextField label={t("Server Address")} variant="standard" />
                </div>
                <div>
                    <TextField  type="email" label={t("Account E-mail")} variant="standard" />
                </div>
                <br/><br/>
                <div>
                    <Button variant="outlined">{t("Register this machine")}</Button>
                </div>
            </div>
        </div>
    </>
}

export default EnrollmentForm