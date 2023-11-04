import { FaMapMarkedAlt } from 'react-icons/fa';

import './Home.scss';

const Home = () => {
  return (
    <>
      <header>
        <div className='logo-container'>
          <FaMapMarkedAlt className='logo-container-logo' />
          <span className='logo-container-title'>BurekParking</span>
        </div>
        <div className='user-container'>A</div>
      </header>
      <section>
      </section>
    </>
  );
}

export default Home;