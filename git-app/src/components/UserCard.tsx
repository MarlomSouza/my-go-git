import LogOut from "./Logout";
import { User } from "../models/types";

const UserCard = ({ user }: { user: User }) => (
  <div className="relative w-full rounded-lg bg-white p-6 shadow-lg">
    <LogOut />
    <div className="mb-4 flex items-center">
      <img
        src={user.avatar_url}
        alt="Avatar"
        className="mr-4 h-20 w-20 rounded-full"
      />
      <div>
        <h2 className="text-xl font-bold">{user.name}</h2>
        <p className="text-gray-600">@{user.login}</p>
        {user.email && <p className="mt-1 text-gray-700">{user.email}</p>}
      </div>
    </div>
    <div className="mb-4 grid grid-cols-2 gap-4">
      <div>
        <p className="text-gray-800">
          <strong>Followers:</strong> {user.followers}
        </p>
        <p className="text-gray-800">
          <strong>Following:</strong> {user.following}
        </p>
      </div>
      <div>
        <p className="text-gray-800">
          <strong>Account Created:</strong>{" "}
          {new Date(user.created_at).toLocaleDateString()}
        </p>
        <p className="text-gray-800">
          <strong>Last Updated:</strong>{" "}
          {new Date(user.updated_at).toLocaleDateString()}
        </p>
      </div>
    </div>
    <div className="mb-4">
      <p className="text-gray-800">
        <strong>Two-Factor Authentication:</strong>{" "}
        {user.two_factor_authentication ? "Enabled" : "Disabled"}
      </p>
    </div>
    <div className="flex justify-between">
      <span className="rounded-full bg-green-200 px-3 py-1 text-green-700">
        Public Repos: {user.public_repos}
      </span>
      <span className="rounded-full bg-red-200 px-3 py-1 text-red-700">
        Private Repos: {user.total_private_repos}
      </span>
    </div>
  </div>
);

export default UserCard;
