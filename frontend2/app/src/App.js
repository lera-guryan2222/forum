import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Login from './components/Auth/Login';
import Register from './components/Auth/Register';
import PostList from './components/Posts/PostList';
import CreatePost from './components/Posts/CreatePost';
import MainLayout from './components/Layout/MainLayout';
import Chat from './components/Forum/Chat';
import { AuthProvider } from './components/Auth/AuthContext';
import styled from 'styled-components';

const AppContainer = styled.div`
  text-align: center;
  min-height: 100vh;
  background-color:rgb(64, 61, 61);
`;

const App = () => {
    const onPostCreated = () => {
        console.log('Post was created');
    };

    return (
        <AuthProvider>
            <AppContainer>
                <Router>
                    <MainLayout>
                        <Routes>
                            <Route path="/login" element={<Login />} />
                            <Route path="/register" element={<Register />} />
                            <Route path="/posts" element={
                                <div>
                                    <CreatePost onPostCreated={onPostCreated} />
                                    <PostList />
                                </div>
                            } />
                            <Route path="/chat" element={<Chat />} />
                            <Route path="/" element={<Navigate to="/posts" />} />
                        </Routes>
                    </MainLayout>
                </Router>
            </AppContainer>
        </AuthProvider>
    );
};

export default App;