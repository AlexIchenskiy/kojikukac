import axios from "axios";
import Header from "../../components/header/Header";
import Maps from "../../components/maps/Maps";
import "./Home.scss";

const Home = () => {
  /*axios
    .get("http://localhost:8080/api/user")
    .then((res) => {
      console.log(res);
    })
    .catch((err) => {
      console.error(err);
    });*/

  return (
    <>
      <Header />
      <Maps />
    </>
  );
};

export default Home;
