import { ReactNode } from "react";
import { useAuth } from "./context/AuthContext";
import Dashboard from "./pages/Dashboard/Dashboard";
import GitHubLogin from './pages/Login/GitHubLogin';

interface ContainerProps {
  children: ReactNode;
}

const Container = ({ children }: ContainerProps) => {
  return <section className="flex flex-col bg-gray-100">{children}</section>;
};

interface LayoutProps {
  children: ReactNode;
}

export const Layout = ({ children }: LayoutProps) => {
  return (
    <Container>
      <MainContainer>{children}</MainContainer>
    </Container>
  );
};

interface MainContainerProps {
  children: ReactNode;
}

const MainContainer = ({ children }: MainContainerProps) => {
  return <main className="flex flex-col">{children}</main>;
};

export default function App() {
  const { isAuthenticated } = useAuth();

  return <Layout>{isAuthenticated ? <Dashboard /> : <GitHubLogin />}</Layout>;
}