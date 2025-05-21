// import React, { createContext, useContext, useState, useEffect } from 'react';
// import axios from 'axios';

// const AuthContext = createContext(null);

// export const useAuth = () => {
//     const context = useContext(AuthContext);
//     if (!context) {
//         throw new Error('useAuth must be used within an AuthProvider');
//     }
//     return context;
// };

// export const AuthProvider = ({ children }) => {
//     const [isAuthenticated, setIsAuthenticated] = useState(false);
//     const [user, setUser] = useState(null);
//     const [loading, setLoading] = useState(true);
//     const [authToken, setAuthToken] = useState(localStorage.getItem('token'));

//     const api = axios.create({
//         baseURL: 'http://localhost:8080',
//         withCredentials: true,
//         headers: {
//             'Content-Type': 'application/json',
//         }
//     });

//     // Добавляем перехватчик для логирования запросов
//     api.interceptors.request.use(request => {
//         console.log('Starting Request:', request);
//         return request;
//     });

//     // Добавляем перехватчик для логирования ответов
//     api.interceptors.response.use(
//         response => {
//             console.log('Response:', response);
//             return response;
//         },
//         error => {
//             console.error('Response Error:', error.response || error);
//             return Promise.reject(error);
//         }
//     );

//     useEffect(() => {
//         checkAuth();
//     }, []);

//     const checkAuth = async () => {
//         try {
//             if (!authToken) {
//                 setLoading(false);
//                 return;
//             }

//             const response = await api.get('/api/auth/verify', {
//                 headers: { Authorization: `Bearer ${authToken}` }
//             });

//             setIsAuthenticated(true);
//             setUser(response.data.user);
//         } catch (error) {
//             console.error('Auth check failed:', error);
//             logout();
//         } finally {
//             setLoading(false);
//         }
//     };

//     const login = async (email, password) => {
//         try {
//             console.log('Attempting login with:', { email });
//             const response = await api.post('/api/auth/login', { email, password });
//             console.log('Login response:', response.data);
            
//             localStorage.setItem('token', response.data.access_token);
//             setAuthToken(response.data.access_token);
//             setIsAuthenticated(true);
//             setUser(response.data.user);
//             return { success: true };
//         } catch (error) {
//             console.error('Login failed:', error.response || error);
//             return {
//                 success: false,
//                 error: error.response?.data?.message || 'Login failed'
//             };
//         }
//     };

//     const register = async (username, email, password) => {
//         try {
//             console.log('Attempting registration with:', { username, email });
//             const response = await api.post('/api/auth/register', { 
//                 username, 
//                 email, 
//                 password 
//             });
//             console.log('Registration response:', response.data);
            
//             if (response.data.access_token) {
//                 localStorage.setItem('token', response.data.access_token);
//                 setAuthToken(response.data.access_token);
//                 setIsAuthenticated(true);
//                 setUser(response.data.user);
//                 return { success: true };
//             }
//             return { success: false, error: 'No token received' };
//         } catch (error) {
//             console.error('Registration failed:', error.response || error);
//             return {
//                 success: false,
//                 error: error.response?.data?.message || 'Registration failed'
//             };
//         }
//     };

//     const logout = () => {
//         localStorage.removeItem('token');
//         setAuthToken(null);
//         setIsAuthenticated(false);
//         setUser(null);
//     };

//     const value = {
//         isAuthenticated,
//         user,
//         loading,
//         authToken,
//         login,
//         register,
//         logout
//     };

//     return (
//         <AuthContext.Provider value={value}>
//             {!loading && children}
//         </AuthContext.Provider>
//     );
// };


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
        baseURL: 'http://localhost:8081/api', // Добавлен /api в базовый URL
        withCredentials: true,
        headers: {
            'Content-Type': 'application/json',
        }
    });

    // Добавляем перехватчик для добавления токена в заголовки
    api.interceptors.request.use(config => {
        const token = localStorage.getItem('token');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
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

            const response = await api.get('/auth/verify'); // Убрали /api из URL, так как он уже в baseURL
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
        console.log('Attempting login with:', { email });
        const response = await api.post('/auth/login', { email, password });
        console.log('Login response:', response.data);

        if (response.data.access_token) {
            localStorage.setItem('token', response.data.access_token);
            setAuthToken(response.data.access_token);
            setIsAuthenticated(true);
            setUser(response.data.user);
            return { success: true };
        } else {
            return { success: false, error: 'No access token received' };
        }
    } catch (error) {
        console.error('Login failed:', error.response?.data || error);
        let errorMessage = 'Login failed';

        if (error.response) {
            // Ошибка от сервера
            errorMessage = error.response.data.message || errorMessage;
        } else if (error.request) {
            // Запрос был сделан, но ответ не получен
            errorMessage = 'No response received from server';
        } else {
            // Произошла ошибка при настройке запроса
            errorMessage = error.message;
        }

        return {
            success: false,
            error: errorMessage
        };
    }
};


   // В функции register
const register = async (username, email, password) => {
    try {
        console.log('Attempting registration with:', { username, email });
        const response = await api.post('/auth/register', {
            username: username,
            email: email,
            password: password
        });
        console.log('Registration response:', response.data);

        if (response.data.access_token) {
            localStorage.setItem('token', response.data.access_token);
            setAuthToken(response.data.access_token);
            setIsAuthenticated(true);
            setUser(response.data.user);
            return { success: true };
        } else {
            return { success: false, error: 'No access token received' };
        }
    } catch (error) {
        console.error('Registration failed:', error.response?.data || error);
        let errorMessage = 'Registration failed';

        if (error.response) {
            // Ошибка от сервера
            errorMessage = error.response.data.message || error.response.data.error || errorMessage;
        } else if (error.request) {
            // Запрос был сделан, но ответ не получен
            errorMessage = 'No response received from server';
        } else {
            // Произошла ошибка при настройке запроса
            errorMessage = error.message;
        }

        return {
            success: false,
            error: errorMessage
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
