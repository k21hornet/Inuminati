import Sidebar from '../components/Sidebar';
import Timeline from '../components/Timeline';

const Home = () => {

  return (
    <div className='flex h-screen bg-slate-50'>
      <Sidebar/>
      <div className='md:ml-60 w-full mb-24 md:mb-0 overflow-scroll'>
        <Timeline user={null}/>
      </div>
    </div>
  )
}

export default Home