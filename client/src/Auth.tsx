import React from 'react'

interface Props {
    isAuthenticated: boolean,
    handleLogin: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void,
    handleLogout: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
}


const Auth: React.FC<Props> = ({ isAuthenticated, handleLogin, handleLogout }) => {
    if (!isAuthenticated) {
        return (
            <button type="button" onClick={handleLogin}>login</button>
        )
    }
    return <button type="button" onClick={handleLogout}>log out</button>


}

export default Auth