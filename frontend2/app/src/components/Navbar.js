import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
    return (
        <nav>
            <ul>
                <li><Link to="/login">Login</Link></li>
                <li><Link to="/register">Register</Link></li>
                <li><Link to="/chat">Chat</Link></li>
                <li><Link to="/messages">Messages</Link></li>
            </ul>
        </nav>
    );
};

export default Navbar;
