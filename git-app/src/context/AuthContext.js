import React, { createContext, useEffect, useState } from "react";

import api from "../services/api";

const AuthContext = createContext();

const baseURL = process.env.REACT_APP_API_URL || "http://localhost:5000";
export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [organization, setOrganization] = useState(null);

  useEffect(() => {
    // Check if access_token cookie exists

    const token = document.cookie
      .split("; ")
      .find((row) => row.startsWith("access_token="))
      ?.split("=")[1];

    setIsAuthenticated(!!token); // User is authenticated
  }, []);

  const login = () => {
    window.location.href = baseURL + "/login/github";
  };

  const logout = async () => {
    await api.post("/logout");
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider
      value={{ isAuthenticated, login, logout, organization, setOrganization }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => React.useContext(AuthContext);
