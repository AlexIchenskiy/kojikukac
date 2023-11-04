import { createContext, useContext } from "react";

const AuthContext = createContext(localStorage.getItem('token') || '');

export const useAuth = () => {
  return useContext(AuthContext);
};

export default AuthContext;