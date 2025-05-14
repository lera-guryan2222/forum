import React, { createContext, useContext, useState, useEffect } from 'react';
import axios from 'axios';

const AuthContext = createContext(null);

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};

export const AuthProvider = ({ children }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const [authToken, setAuthToken] = useState(localStorage.getItem('token'));

    const api = axios.create({
        baseURL: 'http://localhost:8080/api/auth',
        withCredentials: true,
        headers: {
            'Content-Type': 'application/json',
        }
    });

    useEffect(() => {
        checkAuth();
    }, []);

    const checkAuth = async () => {
        try {
            if (!authToken) {
                setLoading(false);
                return;
            }

            const response = await api.get('/verify', {
                headers: { Authorization: `Bearer ${authToken}` }
            });

            setIsAuthenticated(true);
            setUser(response.data.user);
        } catch (error) {
            console.error('Auth check failed:', error);
            logout();
        } finally {
            setLoading(false);
        }
    };

    const login = async (email, password) => {
        try {
            const response = await api.post('/login', { email, password });
            
            localStorage.setItem('token', response.data.access_token);
            setAuthToken(response.data.access_token);
            setIsAuthenticated(true);
            setUser(response.data.user);
            return { success: true };
        } catch (error) {
            console.error('Login failed:', error);
            return {
                success: false,
                error: error.response?.data?.message || 'Login failed'
            };
        }
    };

    const register = async (username, email, password) => {
        try {
            const response = await api.post('/register', { 
                username, 
                email, 
                password 
            });
            
            if (response.data.access_token) {
                localStorage.setItem('token', response.data.access_token);
                setAuthToken(response.data.access_token);
                setIsAuthenticated(true);
                setUser(response.data.user);
                return { success: true };
            }
            return { success: false, error: 'No token received' };
        } catch (error) {
            console.error('Registration failed:', error);
            return {
                success: false,
                error: error.response?.data?.message || 'Registration failed'
            };
        }
    };

    const logout = () => {
        localStorage.removeItem('token');
        setAuthToken(null);
        setIsAuthenticated(false);
        setUser(null);
    };

    const value = {
        isAuthenticated,
        user,
        loading,
        authToken,
        login,
        register,
        logout
    };

    return (
        <AuthContext.Provider value={value}>
            {!loading && children}
        </AuthContext.Provider>
    );
};