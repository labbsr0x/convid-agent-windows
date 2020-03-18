import { useState } from 'react'

export default () => {
    const [serverAddress, setServerAddress] = useState("")
    const [accountEmail, setAccountEmail] = useState("")

    const enroll = () => {

    }

    return {
        serverAddress, setServerAddress,
        accountEmail, setAccountEmail
    }
}