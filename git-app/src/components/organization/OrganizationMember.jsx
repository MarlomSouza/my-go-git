import React from 'react';

const OrganizationMember = ({ member }) => {
  return (
    <div className="flex items-center p-4 border border-gray-300 rounded-lg max-w-md mx-auto bg-gray-50 shadow">
      
      <img 
        src={member.avatar_url} 
        alt={`${member.login}'s avatar`} 
        className="w-16 h-16 rounded-full border border-gray-200"
      />
     
      <div className="ml-4">
        <h2 className="text-lg font-semibold text-gray-800">{member.login}</h2>
        <a 
          href={member.html_url} 
          target="_blank" 
          rel="noopener noreferrer" 
          className="text-blue-600 hover:underline"
        >
          View Profile
        </a>
        
      </div>
    </div>
  );
};

export default OrganizationMember;