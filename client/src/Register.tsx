import React from 'react'

interface Props {
    handleRegister: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
}

const Register: React.FC<Props> = ({ handleRegister }) => {

    return (
        <button type="button" onClick={handleRegister}>register</button>
    )
}

export default Register