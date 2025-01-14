import React from 'react';
import OrganizationDetails from './organizationDetails';

const OrganizationList = ({ organizations }) => {
    return (
        <>
            {organizations.map((org, index) => (
                <OrganizationDetails key={`${index}-org-detail`} org={org} />
            ))}
        </>
    );
};

export default OrganizationList;
