import Loading from "../Loading";
import { Organization } from "../../models/types";
import OrganizationDetails from "./OrganizationDetails";

const OrganizationList = ({
  organizations,
  isLoading,
}: {
  organizations: Organization[];
  isLoading: boolean;
}) => {
  return (
    <div className="w-full rounded-lg bg-white p-4 shadow-lg">
      <h1 className="mb-4 text-2xl font-bold">Organizations</h1>
      {isLoading ? (
        <Loading />
      ) : (
        organizations.map((org, index) => (
          <div key={`${index}-org-detail`} className="flex flex-col gap-2">
            <OrganizationDetails org={org} />
          </div>
        ))
      )}
    </div>
  );
};

export default OrganizationList;
