import React from 'react';
import StarCount from './StarCount';

const RepositoryDetails = ({ repo }) => {
  return (
    <li className="border-b py-4 flex justify-between items-center hover:bg-gray-50 transition-colors">

      <span className="font-medium text-lg flex-grow">{repo.name}</span>

      <div className="flex items-center space-x-1 ml-auto">
        <StarCount count={repo.stargazers_count} />
      </div>


      <span
        className={`px-3 py-1 text-xs font-semibold rounded-full ml-3 ${repo.private ? "bg-red-200 text-red-700" : "bg-green-200 text-green-700"}`}
      >
        {repo.private ? "Private" : "Public"}
      </span>
    </li>

  );
};

export default RepositoryDetails;