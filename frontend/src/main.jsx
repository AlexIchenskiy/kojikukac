import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import "./index.scss";
import Home from "./pages/Home/Home";
import Auth from "./pages/Auth/Auth";
import Profile from "./pages/Profile/Profile";
import { ProtectedRoute } from "./components/Routes/ProtectedRoute";
import AuthProvider from "./provider/AuthProvider";
import { UnauthenticatedRoute } from "./components/Routes/UnathenticatedRoute";

ReactDOM.createRoot(document.getElementById("root")).render(
  <AuthProvider>
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
          path="/profile/:id"
          element={
            <>
              <ProtectedRoute />
              <Profile />
            </>
          }
        />
      </Routes>
    </BrowserRouter>
  </AuthProvider>
);
