import { useEffect, useState } from 'react';

import { Member } from '../../models/types';
import api from '../../services/api';
import Loading from '../Loading';
import OrganizationMember from './OrganizationMember';

const OrganizationMemberList = ({ orgName }: {orgName: string}) => {
  const [members, setMembers] = useState<Member[]>([]);
  const [isLoading, setIsLoading] = useState(true);

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
    <div className="w-full  bg-white shadow-lg rounded-lg p-4">
      <h1 className="text-2xl font-bold mb-4">{orgName} members</h1>
      {isLoading ? <Loading /> :
        (
          members.map((member) => (
            <>
              <OrganizationMember key={member.id} member={member} />
            </>

          ))
        )
      }

    </div>
  );
};

export default OrganizationMemberList;