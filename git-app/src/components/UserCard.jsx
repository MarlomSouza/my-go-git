import React from 'react';

const UserCard = ({ user }) => (
  <div className="w-full bg-white shadow-lg rounded-lg p-6 mb-6">
    <div className="flex items-center mb-4">
      <img
        src={user.avatar_url}
        alt="Avatar"
        className="w-20 h-20 rounded-full mr-4"
      />
      <div>
        <h2 className="text-xl font-bold">{user.name}</h2>
        <p className="text-gray-600">@{user.login}</p>
        {user.email && <p className="text-gray-700 mt-1">{user.email}</p>}
      </div>
    </div>
    <div className="grid grid-cols-2 gap-4 mb-4">
      <div>
        <p className="text-gray-800"><strong>Followers:</strong> {user.followers}</p>
        <p className="text-gray-800"><strong>Following:</strong> {user.following}</p>
      </div>
      <div>
        <p className="text-gray-800"><strong>Account Created:</strong> {new Date(user.created_at).toLocaleDateString()}</p>
        <p className="text-gray-800"><strong>Last Updated:</strong> {new Date(user.updated_at).toLocaleDateString()}</p>
      </div>
    </div>
    <div className="mb-4">
      <p className="text-gray-800"><strong>Two-Factor Authentication:</strong> {user.two_factor_authentication ? 'Enabled' : 'Disabled'}</p>
    </div>
    <div className="flex justify-between">
      <span className="bg-blue-100 text-blue-600 px-3 py-1 rounded-full">
        Public Repos: {user.public_repos}
      </span>
      <span className="bg-green-100 text-green-600 px-3 py-1 rounded-full">
        Private Repos: {user.total_private_repos}
      </span>
    </div>
  </div>
);

export default UserCard;