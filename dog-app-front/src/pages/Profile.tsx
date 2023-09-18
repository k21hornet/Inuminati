import Sidebar from '../components/Sidebar';
import UserProfile from '../components/UserProfile';
import Timeline from '../components/Timeline';
import { useParams } from 'react-router-dom';

const Profile = () => {
  const paramsId = useParams().id;

  return (
    <div className='flex'>
      <Sidebar/>
      <div className='md:ml-60 w-full h-screen bg-slate-50 mb-24 md:mb-0 overflow-scroll'>
        <UserProfile/>
        <Timeline user={paramsId}/>
      </div>
    </div>
  )
}

export default Profile