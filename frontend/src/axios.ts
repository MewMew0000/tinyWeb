import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:3000/api',
});

instance.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        if (!config.headers) {
            config.headers = {};
        }
        config.headers.Authorization = token;
        console.log(token)
    }
    return config;
});

export default instance;
