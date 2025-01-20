import { Repository } from "../models/types";
import StarCount from "./StarCount";

const RepositoryDetails = ({ repo }: { repo: Repository }) => {
  return (
    <div className="flex items-center justify-between border-b py-4 transition-colors hover:bg-gray-50">
      <span className="flex-grow text-lg font-medium">{repo.name}</span>

      <div className="ml-auto flex items-center space-x-1">
        <StarCount count={repo.stargazers_count} />
      </div>

      <span
        className={`ml-3 rounded-full px-3 py-1 text-xs font-semibold ${repo.private ? "bg-red-200 text-red-700" : "bg-green-200 text-green-700"}`}
      >
        {repo.private ? "Private" : "Public"}
      </span>
      <a
        href={repo.html_url}
        target="_blank"
        rel="noopener noreferrer"
        className="ml-4 flex items-center text-blue-600 hover:underline"
      >
        Go to Repo
        <svg
          className="ml-1 h-4 w-4"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d="M14 3h7v7m0 0L10 21l-7-7 11-11z"
          />
        </svg>
      </a>
    </div>
  );
};

export default RepositoryDetails;
