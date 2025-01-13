import { useEffect, useState } from "react";

import api from "../services/api";

const RepoList = () => {
  const [repos, setRepos] = useState([]);
  const [privateRepo, setPrivateRepo] = useState(false);
    
    const fetchRepo = async () => {
      try {
            if (privateRepo) {
              
            const response = await api.get('/repos/private');
            setRepos(response.data);
        }
            else {
                const response = await api.get('/repos/');
            setRepos(response.data);
        }
        } catch (error) {
            console.error('Error fetching repositories:', error);
        }
  }
  

  useEffect(() => {
    fetchRepo();
  }, []);
  
  useEffect(() => {
    fetchRepo();
  }, [privateRepo]);


  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-100 p-6">
      <h2 className="text-2xl font-bold mb-4">Your Repositories</h2>
      <div className="flex items-center mb-4">
        <input
          type="checkbox"
          id="private"
          checked={privateRepo}
          onChange={() => setPrivateRepo(!privateRepo)}
          className="mr-2"
        />
        <label htmlFor="private">Show Private Repositories</label>
      </div>


        

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

export default RepoList;