import { Link } from 'react-router-dom';
import { FaMapMarkedAlt } from 'react-icons/fa';

import './Header.scss';

const Header = () => {
  return (
    <header>
      <Link to='/home' className='logo-container'>
        <FaMapMarkedAlt className='logo-container-logo' />
        <span className='logo-container-title'>BurekParking</span>
      </Link>
      <div className='user-container'>A</div>
    </header>
  )
}

export default Header;