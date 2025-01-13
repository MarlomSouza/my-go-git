import React, { useEffect, useState } from "react";

import { useAuth } from "../../context/AuthContext";
import api from "../../services/api";

const Dashboard = () => {
  const [user, setUser] = useState(null);
  const [error, setError] = useState(null);
  const {logout} = useAuth();


  const fetchUser = async () => {
    try {
      const response = await api.get("/login/user");
      setUser(response.data);
    } catch (error) {
      setError(error.message);
    }
  }

  useEffect(() => {
    fetchUser();
  }, []);

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (!user) {
    return <p>Loading...</p>;
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="p-8 bg-white rounded-lg shadow-lg">
        <button onClick={logout} className="text-sm text-red-500"> Logout</button>
        <h1 className="text-2xl font-bold">Welcome, {user.name}!</h1>
        <p>Your email: {user.email}</p>
        <p>Repositories: {user.repoCount}</p>

        <RepoList />
      </div>
    </div>
  );
};





const RepoList = () => {
    const [repos, setRepos] = useState([]);
    
    const fetchRepo = async () => {
        try {
            const response = await api.get('/repos/');
            setRepos(response.data);
        } catch (error) {
            console.error('Error fetching repositories:', error);
        }
    }

  useEffect(() => {
    fetchRepo();
  }, []);

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-100 p-6">
      <h2 className="text-2xl font-bold mb-4">Your Repositories</h2>
      <ul className="w-full max-w-4xl bg-white p-4 rounded-lg shadow-md">
        {repos.length === 0 ? (
          <p>No repositories found.</p>
        ) : (
          repos.map((repo, index) => (
            <li key={index} className="border-b border-gray-200 py-4">
              <div className="flex items-start justify-between">
                <h3 className="text-xl font-semibold">{repo.name}</h3>
                <span
                  className={`text-sm ${
                    repo.private ? 'text-red-500' : 'text-green-500'
                  }`}
                >
                  {repo.private ? 'Private' : 'Public'}
                </span>
              </div>
              <p className="text-gray-600">{repo.description}</p>
              <p className="text-sm text-gray-500">
                Last Update: {new Date(repo.updated_at).toLocaleDateString()}
              </p>
            </li>
          ))
        )}
      </ul>
    </div>
  );
};

export default Dashboard;
