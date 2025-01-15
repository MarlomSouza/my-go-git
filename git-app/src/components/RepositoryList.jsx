import React, { useState } from 'react';

import FilterInput from './FilterInput';
import Pagination from './Pagination';
import RepositoryDetails from './RepositoryDetails';

const RepositoryList = ({ repos }) => {
  const [filter, setFilter] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(5);

  const filteredRepos = repos.filter(repo =>
    repo.name.toLowerCase().includes(filter.toLowerCase())
  );

  const totalPages = Math.ceil(filteredRepos.length / itemsPerPage);
  const paginatedRepos = filteredRepos.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage
  );

  const handleFilterChange = (e) => {
    setFilter(e.target.value);
    setCurrentPage(1);
  };

  const handlePageChange = (page) => {
    setCurrentPage(page);
  };

  const handleItemsPerPageChange = (e) => {
    setItemsPerPage(Number(e.target.value));
    setCurrentPage(1);
  };

  return (
    <div className="bg-white shadow rounded-lg p-4 flex flex-col flex-1 justify-items-stretch gap-2">
      <h3 className="text-lg font-semibold mb-2">Repositories</h3>
      <FilterInput filter={filter} onFilterChange={handleFilterChange} />
      <div className='flex-grow'>
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