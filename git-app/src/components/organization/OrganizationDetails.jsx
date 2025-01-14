import React from 'react';
import { useAuth } from '../../context/AuthContext';

const OrganizationDetails = ({ org }) => {
  const { setOrganization } = useAuth();

  return (
    <div className="flex items-center p-4 border border-gray-300 rounded-lg bg-gray-50 shadow h-auto">
      <img
        src={org.avatar_url}
        alt={`${org.login}'s avatar`}
        className="w-16 h-16 rounded-full border border-gray-200"
      />
      <div className="ml-4">
        <h2 className="text-lg font-semibold text-gray-800">{org.login}</h2>
        <p className="text-gray-600">{org.description || 'No description provided'}</p>
        <button
          className="text-blue-600 hover:underline"
          onClick={() => setOrganization(org)}
        >
          View Details
        </button>

      </div>
    </div>
  );
}

export default OrganizationDetails;
