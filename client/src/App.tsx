import React, { useState, useEffect } from "react";
import Web3 from 'web3'
import { getNonce, postSignature, register as apiRegister } from "./api/api";
import Auth from "./Auth";
import Register from "./Register"

const myWindow = window as any

const App: React.FC = () => {
  const [isWeb3Active, setIsWeb3Active] = useState<boolean>(false)
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false)
  const [wasRegistrationSuccessful, setWasRegistrationSuccessful] = useState<boolean | null>()


  const handleWeb3 = async () => {
    if (myWindow.ethereum) {
      myWindow.web3 = new Web3(myWindow.ethereum)
      try {
        await myWindow.ethereum.enable()
        setIsWeb3Active(true)
      } catch (error) {
        console.error(error)
        setIsWeb3Active(false)
      }
    } else if (myWindow.web3) {
      myWindow.web3 = new Web3(myWindow.web3.currentProvider)
      setIsWeb3Active(true)
    } else {
      window.alert('no metamask')
      setIsWeb3Active(false)
    }
  }

  useEffect(() => {
    handleWeb3()
  }, [handleWeb3])
  console.log("isweb3active", isWeb3Active)

  const handleLogin = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    // console.log("handleLogin e:", e)
    const coinbase = await myWindow.web3.eth.getCoinbase(console.log)
    interface Params {
      pb: string
    }
    console.log(coinbase)
    const params: Params = {
      pb: coinbase
    }
    const resp = await getNonce(params)
    console.log(resp)
    const nonceResp = resp.data.nonce
    const userSignature = await myWindow.web3.eth.personal.sign(nonceResp, coinbase, console.log)
    console.log(userSignature)
    interface Data {
      pb: string,
      sig: string
    }
    const data: Data = {
      pb: coinbase,
      sig: userSignature
    }
    const isAuthResp = await postSignature(data)
    console.log('isauth login resp', isAuthResp.data.authenticated)
    setIsAuthenticated(isAuthResp.data.authenticated)
  }

  const handleLogout = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setIsAuthenticated(false)
  }

  console.log('state isauthenticated: ', isAuthenticated)

  const handleRegister = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    if (myWindow.ethereum) {
      myWindow.web3 = new Web3(myWindow.ethereum)
      try {
        await myWindow.ethereum.enable()
      } catch (error) {
        console.error(error)
      }
    } else if (myWindow.web3) {
      myWindow.web3 = new Web3(myWindow.web3.currentProvider)
    } else {
      console.log('no metamask')
    }
    const coinbase = await myWindow.web3.eth.getCoinbase(console.log)
    try {
      interface Data {
        pb: string
      }
      const data: Data = {
        pb: coinbase
      }
      const resp = await apiRegister(data)
      console.log("register resp: ", resp)
      setWasRegistrationSuccessful(true)
    } catch (err) {
      console.error(err)
      console.log("after api register error", err.response.status)
      setWasRegistrationSuccessful(false)
    }
  }

  console.log("was registration successful", wasRegistrationSuccessful)

  return (
    <div className="App">
      <Auth isAuthenticated={isAuthenticated} handleLogin={handleLogin} handleLogout={handleLogout} />
      <Register handleRegister={handleRegister} />
    </div>
  );
};

export default App;
