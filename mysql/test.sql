USE golang_ethereum_auth;

Create table `users` (
    id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    public_key VARCHAR(42) UNIQUE NOT NULL,
    nonce varchar(38) NOT NULL
);