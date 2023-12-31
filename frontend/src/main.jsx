import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import "./index.scss";
import Home from "./pages/Home/Home";
import Auth from "./pages/Auth/Auth";
import Profile from "./pages/Profile/Profile";
import { ProtectedRoute } from "./components/Routes/ProtectedRoute";
import AuthProvider from "./provider/AuthProvider";
import { UnauthenticatedRoute } from "./components/Routes/UnathenticatedRoute";
import { RecoilRoot } from "recoil";

ReactDOM.createRoot(document.getElementById("root")).render(
  <AuthProvider>
    <RecoilRoot>
      <BrowserRouter>
        <Routes>
          <Route index element={<Navigate to="/auth" />} />
          <Route
            exact
            path="/auth"
            element={
              <>
                <UnauthenticatedRoute />
                <Auth />
              </>
            }
          />
          <Route
            exact
            path="/home"
            element={
              <>
                <ProtectedRoute /> <Home />
              </>
            }
          />
          <Route
            exact
            path="/profile"
            element={
              <>
                <ProtectedRoute />
                <Profile />
              </>
            }
          />
        </Routes>
      </BrowserRouter>
    </RecoilRoot>
  </AuthProvider>
);
