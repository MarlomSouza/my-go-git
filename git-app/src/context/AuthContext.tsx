import React, { ReactNode, createContext, useContext, useEffect, useState } from 'react';

import { Organization } from '../models/types';
import api from '../services/api';

interface AuthContextProps {
  isAuthenticated: boolean;
  login: () => void;
  logout: () => Promise<void>;
  organization: Organization | null;
  setOrganization: React.Dispatch<React.SetStateAction<Organization | null>>;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [organization, setOrganization] = useState<Organization | null>(null);

  useEffect(() => {
    // Check if access_token cookie exists
    const token = document.cookie
      .split('; ')
      .find((row) => row.startsWith('access_token='))
      ?.split('=')[1];

    setIsAuthenticated(!!token); // User is authenticated if token exists
  }, []);

  const login = () => {
    window.location.href = api.defaults.baseURL + '/login/github';
  };

  const logout = async () => {
    await api.post('/logout');
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

export const useAuth = (): AuthContextProps => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};