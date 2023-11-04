import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../../context/authContext";

export const UnauthenticatedRoute = () => {
  const { token } = useAuth();

  if (token) {
    return <Navigate to="/home" />;
  }

  return <Outlet />;
};