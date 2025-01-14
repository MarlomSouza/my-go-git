import React from 'react';
import OrganizationDetails from './OrganizationDetails';

const OrganizationList = ({ organizations, isLoading }) => {
    return (
        <div className="max-w-4xl mx-auto bg-white shadow-lg rounded-lg p-6">
            <h1 className="text-2xl font-bold mb-4">Organizations</h1>
            {isLoading ? <p>Loading...</p> : (<ul className="space-y-4">
                {organizations.map((org, index) => (
                    <li key={`${index}-org-detail`}>
                        <OrganizationDetails org={org} />
                    </li>
                ))}
            </ul>)}

        </div>
    );
};

export default OrganizationList;
