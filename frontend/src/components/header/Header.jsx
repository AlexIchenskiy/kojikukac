import { FaMapMarkedAlt } from 'react-icons/fa';

import './Header.scss';

const Header = () => {
  return (
    <header>
      <div className='logo-container'>
        <FaMapMarkedAlt className='logo-container-logo' />
        <span className='logo-container-title'>BurekParking</span>
      </div>
      <div className='user-container'>A</div>
    </header>
  )
}

export default Header;