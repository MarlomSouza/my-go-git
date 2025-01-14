import React from 'react';
import StarCount from './StarCount';

const RepositoryDetails = ({ repo }) => {
  return (
    <li className="border-b py-2 flex justify-between items-center">
      <span>{repo.name}</span>
      <StarCount count={repo.stargazers_count} />
      <span
        className={`px-2 py-1 text-sm rounded ${repo.private ? "bg-red-200 text-red-700" : "bg-green-200 text-green-700"}`}
      >
        {repo.private ? "Private" : "Public"}
      </span>
    </li>
  );
};

export default RepositoryDetails;