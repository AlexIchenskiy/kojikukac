import ReactDOM from 'react-dom/client'
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import './index.scss'
import Home from './pages/Home/Home';
import Auth from './pages/Auth/Auth';

ReactDOM.createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Routes>

      <Route exact path='/home' element={<Home />} />
    </Routes>
  </BrowserRouter>
)
