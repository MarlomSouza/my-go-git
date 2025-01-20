import { Member } from "../../models/types";

const OrganizationMember = ({ member }: { member: Member }) => {
  return (
    <div className="my-4 flex items-center rounded-lg bg-gray-50 p-4 shadow hover:bg-gray-100">
      <img
        src={member.avatar_url}
        alt={`${member.login}'s avatar`}
        className="mr-4 h-12 w-12 rounded-full border border-gray-200"
      />
      <div className="flex-1">
        <p className="font-semibold text-gray-800">{member.login}</p>
        <a
          href={member.html_url}
          target="_blank"
          rel="noopener noreferrer"
          className="text-blue-600 hover:underline"
        >
          View Profile
        </a>
      </div>
    </div>
  );
};

export default OrganizationMember;
