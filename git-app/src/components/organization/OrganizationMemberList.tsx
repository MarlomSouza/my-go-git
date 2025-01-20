import { useEffect, useState } from "react";

import Loading from "../Loading";
import { Member } from "../../models/types";
import OrganizationMember from "./OrganizationMember";
import api from "../../services/api";

const OrganizationMemberList = ({ orgName }: { orgName: string }) => {
  const [members, setMembers] = useState<Member[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchOrganizationMembers = async () => {
      try {
        const response = await api.get(`/organization/${orgName}/members`);
        setMembers(response.data);
      } catch (err) {
        console.error("Error fetching organization members:", err);
      } finally {
        setIsLoading(false);
      }
    };
    if (orgName) {
      fetchOrganizationMembers();
    }
  }, [orgName]);

  return (
    <div className="w-full rounded-lg bg-white p-4 shadow-lg">
      <h1 className="mb-4 text-2xl font-bold">{orgName} members</h1>
      {isLoading ? (
        <Loading />
      ) : (
        members.map((member) => (
          <>
            <OrganizationMember key={member.id} member={member} />
          </>
        ))
      )}
    </div>
  );
};

export default OrganizationMemberList;
