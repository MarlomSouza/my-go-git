import React, { useCallback, useEffect, useState } from "react";

import OrganizationDetails from "../../components/organization/OrganizationDetails";
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


  // const handleMyRepos = () => {
  //   setOrganization({});
  //   fetchRepo();
  // }

  const fetchRepo = async () => {
    try {
      // if (privateRepo) {

      //   const response = await api.get('/repos/private');
      //   setRepos(response.data);
      // }
      // else {
      //   const response = await api.get('/repos/');
      //   setRepos(response.data);
      // }
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

  useEffect(() => {
    if (organization?.login) {
      fetchOrganizationsRepos()
    }

  }, [fetchOrganizationsRepos, organization])

  useEffect(() => {
    fetchRepo();
  }, []);

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
    fetchOrganizations()
  }, []);

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (!user) {
    return <p>Loading...</p>;
  }

  return (
    // <div className="flex flex-col items-center justify-center min-h-screen">
    //   <button onClick={logout} className="text-sm text-red-500">Logout</button>
    //   <div className="flex p-2  flex-row rounded-lg shadow-lg">
    //     <div className="flex flex-col rounded-lg shadow-lg">
    //       <UserCard user={user} />
    //       <div className="flex space-x-4 mb-4">
    //         <button className="px-4 py-2 rounded  bg-gray-200 text-gray-700" onClick={handleMyRepos}>
    //           View my repos
    //         </button>

    //       </div>
    //       {organizations.length > 0 && (<OrganizationList organizations={organizations} />)}
    //     </div>
    //     {!!organization && <OrganizationMemberList orgName={organization.login} />}
    //     <RepoList repos={repos} />

    //   </div>
    // </div>
    <div className="bg-gray-100 min-h-screen p-4">
      <div className="flex items-center justify-between  gap-2">
        <div className={organizations ? "w-3/4" : "w-full"}>
          <UserCard user={user} />
        </div>
        <div className="w-1/4  " >
          {/*  list of organization */}
          {organizations.map((org, index) => (
            <OrganizationDetails key={`${index}-org-detail`} org={org} />
          ))}
        </div>
      </div>
      {organization ? (
        <div className="flex items-center justify-between">
          <div className={organizations ? "w-3/4" : "w-full"}>
            <RepositoryList repos={repos} />
          </div>
          <div className="w-1/4  " >
            <OrganizationMemberList orgName={organization.login} />
          </div>
        </div>
      ) : <RepositoryList repos={repos} />}
    </div>
  );
};




export default Dashboard;
