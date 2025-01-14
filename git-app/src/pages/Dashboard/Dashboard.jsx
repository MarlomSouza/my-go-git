import React, { useCallback, useEffect, useState } from "react";

import Loading from "../../components/Loading";
import OrganizationList from "../../components/organization/OrganizationList";
import OrganizationMemberList from "../../components/organization/OrganizationMemberList";
import RepositoryList from "../../components/RepositoryList";
import UserCard from "../../components/UserCard";
import { useAuth } from "../../context/AuthContext";
import api from "../../services/api";

const Dashboard = () => {
  const [user, setUser] = useState(null);
  const [organizations, setOrganizations] = useState([]);
  const [error, setError] = useState(null);
  const { organization } = useAuth();
  const [repos, setRepos] = useState([]);

  const fetchRepo = async () => {
    try {
      const response = await api.get('/repos/');
      setRepos(response.data);
    } catch (error) {
      console.error('Error fetching repositories:', error);
    }
  }

  const fetchOrganizationsRepos = useCallback(async () => {
    try {
      const response = await api.get(`/organization/${organization.login}/repos`);
      setRepos(response.data);
    } catch (error) {
      console.error('Error fetching repositories:', error);
    }
  }, [organization])

  const fetchUser = async () => {
    try {
      const response = await api.get("/repos/user");
      setUser(response.data);
    } catch (error) {
      setError(error.message);
    }
  }

  const fetchOrganizations = async () => {
    try {
      const response = await api.get("/organization/");
      setOrganizations(response.data);
    } catch (error) {
      setError(error.message);
    }
  }

  useEffect(() => {
    fetchUser();
    fetchRepo();
    fetchOrganizations()
  }, []);

  useEffect(() => {
    if (organization?.login) {
      fetchOrganizationsRepos()
    }
    else {
      fetchRepo();
    }

  }, [fetchOrganizationsRepos, organization])

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (!user) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <Loading />
      </div>
    );
  }

  return (

    <div className="bg-gray-100 min-h-screen p-4">
      <div className="flex items-center justify-between gap-2">
        <div className={organizations.length ? "w-3/4" : "w-full"}>
          <UserCard user={user} />
        </div>
        {organizations.length && (
          <div className="w-1/4">
            <OrganizationList organizations={organizations} />
          </div>
        )}
      </div>
      <div className="flex items-center justify-between gap-2 mt-4">
        <div className={organization ? "w-3/4" : "w-full"}>
          <RepositoryList repos={repos} />
        </div>
        {organization && (
          <div className="w-1/4">
            <OrganizationMemberList orgName={organization.login} />
          </div>
        )}
      </div>
    </div>
  );
};




export default Dashboard;
