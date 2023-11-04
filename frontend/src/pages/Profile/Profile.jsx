import { useParams } from "react-router-dom";
import Header from "../../components/header/Header";

const Profile = () => {
  let { id } = useParams();

  return (
    <>
      <Header />
      <div>{ id }</div>
    </>
  )
}

export default Profile;