import React from 'react';

const Pagination = ({ currentPage, totalPages, onPageChange, itemsPerPage, onItemsPerPageChange }) => {
    return (
        <div className="flex justify-between items-center">
            <div>
                <label htmlFor="itemsPerPage" className="mr-2">Items per page:</label>
                <select
                    id="itemsPerPage"
                    value={itemsPerPage}
                    onChange={onItemsPerPageChange}
                    className="p-2 border rounded-lg"
                >
                    <option value={5}>5</option>
                    <option value={10}>10</option>
                    <option value={20}>20</option>
                </select>
            </div>
            <span>Page {currentPage} of {totalPages}</span>
            <div>
                <button
                    onClick={() => onPageChange(currentPage - 1)}
                    disabled={currentPage === 1}
                    className="px-4 py-2 bg-gray-300 rounded disabled:opacity-50"
                >
                    Previous
                </button>
                <button
                    onClick={() => onPageChange(currentPage + 1)}
                    disabled={currentPage === totalPages}
                    className="ml-1 px-4 py-2 bg-gray-300 rounded disabled:opacity-50"
                >
                    Next
                </button>
            </div>
        </div >
    );
};

export default Pagination;