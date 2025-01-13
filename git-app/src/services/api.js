import axios from 'axios';

const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL || "http://localhost:5000/api", // URL do seu backend
  timeout: 10000, // Timeout para as requisições
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true
});

// Interceptor para tratar respostas de erro
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.reject({ ...error.response.data });
  },
);

export default api;
