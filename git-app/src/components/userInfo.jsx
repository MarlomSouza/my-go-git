
const UserInfo  = ({ user }) => {
  return (
    <div className="p-6 bg-white rounded-lg shadow-lg">
      <div className="flex items-center">
        <img
          src={user.avatar_url}
          alt={`${user.name}'s avatar`}
          className="w-16 h-16 rounded-full mr-4"
        />
        <div>
          <h2 className="text-2xl font-bold">{user.name}</h2>
          <p className="text-gray-500">@{user.login}</p>
          <p className="text-gray-600">{user.email}</p>
        </div>
      </div>
      <div className="mt-4">
        <p>
          <strong>Public Repositories:</strong> {user.public_repos}
        </p>
        <p>
          <strong>Private Repositories:</strong> {user.total_private_repos}
        </p>
        <p>
          <strong>Followers:</strong> {user.followers}
        </p>
        <p>
          <strong>Following:</strong> {user.following}
        </p>
        <p>
          <strong>Account Created:</strong> {new Date(user.created_at).toLocaleDateString()}
        </p>
        <p>
          <strong>Last Updated:</strong> {new Date(user.updated_at).toLocaleDateString()}
        </p>
        <p>
          <strong>Two-Factor Authentication:</strong>{" "}
          {user.two_factor_authentication ? "Enabled" : "Disabled"}
        </p>
      </div>
    </div>
  );
};

export default UserInfo;