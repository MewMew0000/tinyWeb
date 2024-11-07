import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import axios from '../axios';

const router = useRouter();

export const useAuthStore = defineStore('auth', () => {
    const token = ref<string | null> (localStorage.getItem('token'));
    const isAuthenticated = computed(() => token.value !== null && token.value !== '');
    const login = async (username: string, password: string) => {
        try {
            const response = await axios.post<{ token: string }>('/auth/login', {username, password});
            token.value = (response.data as { token: string }).token;
            localStorage.setItem('token', token.value || '');
        }
        catch (error) {
            console.error('login error: ', error);
        }
    }
    const register = async (username: string, password: string) => {
        try {
            const response = await axios.post<{ token: string }>('/auth/register', {username, password});
            token.value = response.data.token;
            localStorage.setItem('token', token.value || '');
        } catch (error) {
            console.error('register error: ', error);
        }
    };

    const logout = () => {
        token.value = null;
        localStorage.removeItem('token');
    }

    return {
        token,
        isAuthenticated,
        login,
        register,
        logout
    };
});