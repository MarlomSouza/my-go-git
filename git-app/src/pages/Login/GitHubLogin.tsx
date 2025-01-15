import { useAuth } from "../../context/AuthContext";

const GitHubLogin = () => {
  const {login}  = useAuth();

  return (
    <div className="flex items-center justify-center min-h-screen ">
      <div className="p-8  rounded-lg shadow-lg">
        <h1 className="mb-6 text-2xl font-bold text-center text-gray-800">Welcome to GitHub Repos Viewer</h1>
        <p className="mb-4 text-center text-gray-600">Log in to explore your private and public repositories.</p>
        <button
          onClick={login}
          className="w-full px-6 py-3 text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-4 focus:ring-blue-300"
        >
          Log in with GitHub
        </button>
      </div>
    </div>
  );
};

export default GitHubLogin;
