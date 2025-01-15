import { useAuth } from '../../context/AuthContext';
import { Organization } from '../../models/types';

const OrganizationDetails = ({org}: {org: Organization}) => {
  const { organization, setOrganization } = useAuth();
  const isSelected = organization && organization.login === org.login;

  return (
    <div className={`relative flex items-center p-4 border ${isSelected ? 'border-green-500' : 'border-gray-300'} rounded-lg bg-gray-50 shadow h-auto`}>
      {isSelected && (
        <button
          className="absolute top-2 right-2 text-gray-500 hover:text-gray-700"
          onClick={() => setOrganization(null)}
        >
          &times;
        </button>
      )}
      <img
        src={org.avatar_url}
        alt={`${org.login}'s avatar`}
        className="w-16 h-16 rounded-full border border-gray-200"
      />
      <div className="ml-4">
        <h2 className="text-lg font-semibold text-gray-800">{org.login}</h2>
        <p className="text-gray-600">{org.description || 'No description provided'}</p>
        <button
          className="text-blue-600 hover:underline mt-2"
          onClick={() => setOrganization(org)}
        >
          View Details
        </button>
      </div>
    </div>
  );
};

export default OrganizationDetails;