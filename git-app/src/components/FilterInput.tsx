const FilterInput = ({
  filter,
  onFilterChange,
}: {
  filter: any;
  onFilterChange: any;
}) => {
  return (
    <input
      type="text"
      placeholder="Filter repositories"
      value={filter}
      onChange={onFilterChange}
      className="w-full rounded-lg border p-2"
    />
  );
};

export default FilterInput;
