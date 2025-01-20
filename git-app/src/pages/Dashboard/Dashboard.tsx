import { Organization, Repository, User } from "../../models/types";
import { useCallback, useEffect, useState } from "react";

import Loading from "../../components/Loading";
import OrganizationList from "../../components/organization/OrganizationList";
import OrganizationMemberList from "../../components/organization/OrganizationMemberList";
import RepositoryList from "../../components/RepositoryList";
import UserCard from "../../components/UserCard";
import api from "../../services/api";
import { useAuth } from "../../context/AuthContext";

const Dashboard = () => {
  const [user, setUser] = useState<User | null>(null);
  const [organizations, setOrganizations] = useState<Organization[]>([]);
  const [error, setError] = useState<any | null>(null);
  const { organization } = useAuth();
  const [repos, setRepos] = useState<Repository[]>([]);

  const fetchRepo = async () => {
    try {
      const response = await api.get("/repos/");
      setRepos(response.data);
    } catch (error) {
      console.error("Error fetching repositories:", error);
    }
  };

  const fetchOrganizationsRepos = useCallback(async () => {
    try {
      const response = await api.get(
        `/organization/${organization?.login}/repos`,
      );
      setRepos(response.data);
    } catch (error) {
      console.error("Error fetching organization repositories:", error);
    }
  }, [organization]);

  const fetchUser = async () => {
    try {
      const response = await api.get("/repos/user");
      setUser(response.data);
    } catch (error: any) {
      setError(error.message);
    }
  };

  const fetchOrganizations = async () => {
    try {
      const response = await api.get("/organization/");
      setOrganizations(response.data);
    } catch (error: any) {
      setError(error.message);
    }
  };

  useEffect(() => {
    fetchUser();
    fetchRepo();
    fetchOrganizations();
  }, []);

  useEffect(() => {
    if (organization?.login) {
      fetchOrganizationsRepos();
    } else {
      fetchRepo();
    }
  }, [fetchOrganizationsRepos, organization]);

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (!user) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <Loading />
      </div>
    );
  }

  return (
    <div className="min-h-screen p-4">
      <div className="flex flex-wrap items-stretch justify-between gap-2">
        <div className="flex flex-grow">
          <UserCard user={user} />
        </div>
        {organizations.length > 0 && (
          <div className="flex flex-1">
            <OrganizationList organizations={organizations} isLoading={false} />
          </div>
        )}
      </div>
      <div className="mt-2 flex flex-wrap items-stretch justify-between gap-2">
        <div className="flex flex-grow">
          <RepositoryList repos={repos} />
        </div>
        {organization && (
          <div className="flex flex-grow">
            <OrganizationMemberList orgName={organization.login} />
          </div>
        )}
      </div>
    </div>
  );
};

export default Dashboard;
