import { useAuth } from "../../context/AuthContext";

const GitHubLogin = () => {
  const { login } = useAuth();

  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="rounded-lg p-8 shadow-lg">
        <h1 className="mb-6 text-center text-2xl font-bold text-gray-800">
          Welcome to GitHub Repos Viewer
        </h1>
        <p className="mb-4 text-center text-gray-600">
          Log in to explore your private and public repositories.
        </p>
        <button
          onClick={login}
          className="w-full rounded-md bg-blue-600 px-6 py-3 text-white hover:bg-blue-700 focus:outline-none focus:ring-4 focus:ring-blue-300"
        >
          Log in with GitHub
        </button>
      </div>
    </div>
  );
};

export default GitHubLogin;
