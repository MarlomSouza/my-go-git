import { Organization } from '../../models/types';
import Loading from '../Loading';
import OrganizationDetails from './OrganizationDetails';

const OrganizationList = ({ organizations, isLoading } : {organizations: Organization[], isLoading: boolean} ) => {
    return (
        <div className="w-full bg-white  shadow-lg rounded-lg p-4">
            <h1 className="text-2xl font-bold mb-4">Organizations</h1>
            {isLoading ? <Loading /> :
                organizations.map((org, index) => (
                    <div key={`${index}-org-detail`} className="flex flex-col gap-2">
                        <OrganizationDetails org={org} />
                    </div>
                ))
            }

        </div>
    );
}

export default OrganizationList;
