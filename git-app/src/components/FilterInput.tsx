
const FilterInput = ({ filter, onFilterChange } : { filter : any,onFilterChange: any }) => {
    return (
        <input
            type="text"
            placeholder="Filter repositories"
            value={filter}
            onChange={onFilterChange}
            className="p-2 border rounded-lg w-full"
        />
    );
};

export default FilterInput;