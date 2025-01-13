import React, { useEffect, useState } from "react";

import RepoList from "../../components/repoList";
import UserInfo from "../../components/userInfo";
import { useAuth } from "../../context/AuthContext";
import api from "../../services/api";

const Dashboard = () => {
  const [user, setUser] = useState(null);
  const [error, setError] = useState(null);
  const {logout} = useAuth();
  


  const fetchUser = async () => {
    try {
      const response = await api.get("/repos/user");
      setUser(response.data);
    } catch (error) {
      setError(error.message);
    }
  }

  useEffect(() => {
    fetchUser();
  }, []);

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (!user) {
    return <p>Loading...</p>;
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="p-8 bg-white rounded-lg shadow-lg">
        <button onClick={logout} className="text-sm text-red-500"> Logout</button>
        <UserInfo user={user} />
        <RepoList />
      </div>
    </div>
  );
};




export default Dashboard;
