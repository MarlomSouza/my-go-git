import React, { useEffect, useState } from 'react';

import api from '../../services/api';
import OrganizationMember from './OrganizationMember';

const OrganizationMemberList = ({ orgName }) => {
  const [members, setMembers] = useState([]);

  useEffect(() => {
    const fetchOrganizationMembers = async () => {
      try {
        const response = await api.get(`/organization/${orgName}/members`);
        setMembers(response.data);
      } catch (err) {
        console.error('Error fetching organization members:', members, err);
      }
    };
    if (orgName) {
      fetchOrganizationMembers();
    }
  }, [members, orgName]);

  return (
    <div className="flex  flex-col">
      <h2 className="text-2xl font-bold mb-4">{orgName} members</h2>
      {members.map((member) => (
        <OrganizationMember key={member.id} member={member} />
      ))}
    </div>
  );
};

export default OrganizationMemberList;
