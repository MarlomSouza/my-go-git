import React, { createContext, useEffect, useState } from "react";

import api from "../services/api";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        // Check if access_token cookie exists
        
        const token = document.cookie
            .split("; ")
            .find((row) => row.startsWith("access_token="))
            ?.split("=")[1];

        
        if (token) {
            setIsAuthenticated(true); // User is authenticated
        }
    }, []);

    const login = () => {
        window.location.href = process.env.REACT_APP_API_URL + "/login/github";
    };

    const logout = async () => {
         await api.post("/logout");
        setIsAuthenticated(false);
    };

    return (
        <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => React.useContext(AuthContext);
