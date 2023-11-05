import { useState } from "react";
import Header from "../../components/header/Header";

import "./Profile.scss";
import axios, { AxiosHeaders } from "axios";

const Profile = () => {
  const [userData, setUserData] = useState(null);

  let config = {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    },
  };

  axios
<<<<<<< HEAD
    .get("http://localhost:8080/api/user/profile")
=======
    .get("http://localhost:8080/api/user", config)
>>>>>>> 05953691ac0d66747b2f46ee65c54047c7692068
    .then((res) => {
      setUserData(res);
    })
    .catch((err) => {
      console.error(err);
    });

  return (
    <>
      <Header />
      {userData && (
        <section className="data">
          <div className="data-user">
            <div className="data-user-title">Who am I</div>
            <div className="data-user-block">
              <div className="data-user-block-general">
                <div className="data-user-block-general-heading">Firstname</div>
                <div className="data-user-block-general-subheading">
                  {userData.firstname}
                </div>
                <div className="data-user-block-general-heading">Lastname</div>
                <div className="data-user-block-general-subheading">
                  {userData.lastname}
                </div>
                <div className="data-user-block-general-heading">Email</div>
                <div className="data-user-block-general-subheading">
                  {userData.email}
                </div>
              </div>
              <div className="data-user-block-secondary">
                <div className="data-user-block-general-heading">User ID</div>
                <div className="data-user-block-general-subheading">
                  {userData.ID}
                </div>
                <div className="data-user-block-general-heading">
                  Creation date
                </div>
                <div className="data-user-block-general-subheading">
                  {new Date(userData.createdAt).toUTCString()}
                </div>
              </div>
            </div>
          </div>
          <div className="data-history">
            <div className="data-history-title">Reservation history</div>
            <div>No reservations yet :(</div>
          </div>
        </section>
      )}
    </>
  );
};

export default Profile;
