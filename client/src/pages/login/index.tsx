import { useState } from "react"
import Auth from "./components/auth"
import CreateAccount from "./components/createAccount"

const Login = () => {
  const [modeAuth, setModeAuth] = useState<boolean>(true)

  const toggleAuthentication = (event: boolean) => {
    setModeAuth(event)
  }

  return (
    <>
      {modeAuth ? (
        <Auth toggleAuthentication={toggleAuthentication} />
      ) : (
        <CreateAccount toggleAuthentication={toggleAuthentication} />
      )}
    </>
  )
}

export default Login
