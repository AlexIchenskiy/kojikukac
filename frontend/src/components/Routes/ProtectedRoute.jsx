import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../../context/authContext";

export const ProtectedRoute = () => {
  const { token } = useAuth();

  if (!token) {
    return <Navigate to="/auth" />;
  }

  return <Outlet />;
};