import { useState } from "react";
import { Repository } from "../models/types";
import FilterInput from "./FilterInput";
import Pagination from "./Pagination";
import RepositoryDetails from "./RepositoryDetails";

const RepositoryList = ({ repos }: { repos: Repository[] }) => {
  const [filter, setFilter] = useState("");
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [itemsPerPage, setItemsPerPage] = useState<number>(5);

  const filteredRepos = repos.filter((repo) =>
    repo.name.toLowerCase().includes(filter.toLowerCase()),
  );

  const totalPages = Math.ceil(filteredRepos.length / itemsPerPage);
  const paginatedRepos = filteredRepos.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage,
  );

  const handleFilterChange = (e: any ) => {
    setFilter(e.target.value);
    setCurrentPage(1);
  };

  const handlePageChange = (page : number) => {
    setCurrentPage(page);
  };

  const handleItemsPerPageChange = (e: any) => {
    setItemsPerPage(Number(e.target.value));
    setCurrentPage(1);
  };

  return (
    <div className="flex flex-1 flex-col justify-items-stretch gap-2 rounded-lg bg-white p-4 shadow">
      <h3 className="mb-2 text-lg font-semibold">Repositories</h3>
      <FilterInput filter={filter} onFilterChange={handleFilterChange} />
      <div className="flex-grow">
        {paginatedRepos.map((repo, index) => (
          <RepositoryDetails key={`${index}-repo`} repo={repo} />
        ))}
      </div>

      <div>
        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={handlePageChange}
          itemsPerPage={itemsPerPage}
          onItemsPerPageChange={handleItemsPerPageChange}
        />
      </div>
    </div>
  );
};

export default RepositoryList;
