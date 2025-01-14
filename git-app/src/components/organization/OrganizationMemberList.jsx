import React, { useEffect, useState } from 'react';

import api from '../../services/api';
import Loading from '../Loading';
import OrganizationMember from './OrganizationMember';

const OrganizationMemberList = ({ orgName }) => {
  const [members, setMembers] = useState([]);
  const [isloading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchOrganizationMembers = async () => {
      try {
        const response = await api.get(`/organization/${orgName}/members`);
        setMembers(response.data);
      } catch (err) {
        console.error('Error fetching organization members:', err);
      } finally {
        setIsLoading(false);
      }
    };
    if (orgName) {
      fetchOrganizationMembers();
    }
  }, [orgName]);



  return (
    <div className="max-w-4xl mx-auto bg-white shadow-lg rounded-lg p-6">
      <h1 className="text-2xl font-bold mb-4">{orgName} Members</h1>
      {isloading ? <Loading /> : <ul className="space-y-4">
        {members.map((member) => (
          <OrganizationMember key={member.id} member={member} />
        ))}
      </ul>}

    </div>
  );
};

export default OrganizationMemberList;