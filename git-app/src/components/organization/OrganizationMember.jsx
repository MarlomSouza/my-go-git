import React from 'react';

const OrganizationMember = ({ member }) => {
  return (
    <li className="flex items-center bg-gray-50 p-4 rounded-lg shadow hover:bg-gray-100">
      <img
        src={member.avatar_url}
        alt={`${member.login}'s avatar`}
        className="w-12 h-12 rounded-full border border-gray-200 mr-4"
      />
      <div className="flex-1">
        <p className="text-gray-800 font-semibold">{member.login}</p>
        <a
          href={member.html_url}
          target="_blank"
          rel="noopener noreferrer"
          className="text-blue-600 hover:underline"
        >
          View Profile
        </a>
      </div>
    </li>
  );
};

export default OrganizationMember;