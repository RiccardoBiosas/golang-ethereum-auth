import axios from "axios";

const BASE_URL = "http://localhost:8080/api/";
const LOGIN_ENDPOINT = "auth/login";
const REGISTER_ENDPOINT = "auth/register";



interface NonceEndpoint {
  pb: string
}

interface SignatureEndpoint {
  pb: string,
  sig: string
}

interface RegisterEndpoint {
  pb: string
}

const client = axios.create({
  baseURL: BASE_URL,
  headers: {
    // "Content-Type": "application/json",
    'Content-Type': 'text/plain; charset=utf-8'
  },
});

export const getNonce = async (data: NonceEndpoint) => {
  const resp = await client.get(LOGIN_ENDPOINT, { params: data });
  return resp;
};

export const postSignature = async (data: SignatureEndpoint) => {
  const resp = await client.post(LOGIN_ENDPOINT, data);
  return resp;
};

export const register = async (data: RegisterEndpoint) => {
  const resp = await client.post(REGISTER_ENDPOINT, data);
  return resp;
};
