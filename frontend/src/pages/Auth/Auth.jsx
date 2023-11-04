import { useState } from 'react';

import './Auth.scss';

const Auth = () => {
  const [flipped, setFlipped] = useState(false);
  const [loginData, setLoginData] = useState({
    email: "",
    password: "",
  });
  const [registerData, setRegisterData] = useState({
    name: "",
    surname: "",
    email: "",
    password: "",
  });

  const onLoginSubmit = (e) => {
    e.preventDefault();

    console.log(loginData);
  }

  const onRegisterSubmit = (e) => {
    e.preventDefault();

    console.log(registerData);
  }

  return (
    <div className='auth-container'>
      <div className='auth-container-flip'>
        <div className={`auth-container-flip-inner ${flipped ? 'flipped' : ''}`}>

          <div className='auth-container-flip-inner-login'>
            <div className='auth-container-flip-inner-login-title'>Log in</div>
            <form>

              <label>E-mail</label>
              <input type='text' value={loginData.email} onChange={(e) => setLoginData({...loginData, email: e.target.value})}></input>
              <label>Password</label>
              <input type='password' value={loginData.password} onChange={(e) => setLoginData({...loginData, password: e.target.value})}></input>
              <div className='auth-container-flip-inner-login-button'>
                <input type='submit' value="Submit" onClick={onLoginSubmit} />
              </div>
              <div className='auth-container-flip-inner-login-redirect' onClick={() => setFlipped(true)}>
                Create your profile &rarr;
              </div>

            </form>
          </div>

          <div className='auth-container-flip-inner-register'>
            <div className='auth-container-flip-inner-register-title'>Sign up</div>
            <form>

              <label>First name</label>
              <input type='text' value={registerData.name} onChange={(e) => setRegisterData({...registerData, name: e.target.value})}></input>
              <label>Last name</label>
              <input type='text' value={registerData.surname} onChange={(e) => setRegisterData({...registerData, surname: e.target.value})}></input>
              <label>E-mail</label>
              <input type='text' value={registerData.email} onChange={(e) => setRegisterData({...registerData, email: e.target.value})}></input>
              <label>Password</label>
              <input type='password' value={registerData.password} onChange={(e) => setRegisterData({...registerData, password: e.target.value})}></input>
              <div className='auth-container-flip-inner-register-button'>
                <input type='submit' value="Submit" onClick={onRegisterSubmit} />
              </div>
              <div className='auth-container-flip-inner-register-redirect' onClick={() => setFlipped(false)}>
                Log in to existing profile &rarr;
              </div>
              
            </form>
          </div>

        </div>
      </div>
    </div>
  )
}

export default Auth;