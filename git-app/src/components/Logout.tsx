import { useState } from "react";
import { useAuth } from "../context/AuthContext";

const LogOut = () => {
  const { logout } = useAuth();
  const [showTooltip, setShowTooltip] = useState(false);
  return (
    <div
      onClick={logout}
      onMouseEnter={() => setShowTooltip(true)}
      onMouseLeave={() => setShowTooltip(false)}
      className="absolute right-4 top-4 cursor-pointer text-red-500 hover:text-red-600"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        className="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="2"
          d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1m0-10V5m0 10v1m0-10V5"
        />
      </svg>

      {showTooltip && (
        <div className="absolute right-10 top-0 rounded bg-gray-700 px-2 py-1 text-xs text-white">
          Logout
        </div>
      )}
    </div>
  );
};

export default LogOut;
