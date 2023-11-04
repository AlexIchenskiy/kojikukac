import ReactDOM from 'react-dom/client'
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import './index.scss'
import Home from './pages/Home/Home';
import Auth from './pages/Auth/Auth';

ReactDOM.createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Routes>
      <Route index element={<Navigate to='/auth' />} />
      <Route exact path='/auth' element={<Auth />} />
      <Route exact path='/home' element={<Home />} />
    </Routes>
  </BrowserRouter>
)
