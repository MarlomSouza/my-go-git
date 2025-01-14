import React from 'react';

const FilterInput = ({ filter, onFilterChange }) => {
    return (
        <input
            type="text"
            placeholder="Filter repositories"
            value={filter}
            onChange={onFilterChange}
            className="mb-4 p-2 border rounded w-full"
        />
    );
};

export default FilterInput;