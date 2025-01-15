import React from "react";
import { useAuth } from "./context/AuthContext";
import Dashboard from './pages/Dashboard/Dashboard';
import GitHubLogin from './pages/Login/GitHubLogin';

const MainContainer = ({ children }) => {
  return (
    <main className="flex  flex-col">{children}</main>
  );
};

const Container = ({ children }) => {
  return <section className="flex flex-col  bg-gray-100">{children}</section>;
};

export const Layout = ({ children }) => {
  return (
    <Container>
      <MainContainer>{children}</MainContainer>
    </Container>
  );
};

export default function App() {
  const { isAuthenticated } = useAuth();

  return <Layout>{isAuthenticated ? <Dashboard /> : <GitHubLogin />}</Layout>;
  
}


