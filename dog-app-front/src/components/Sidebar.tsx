import React from 'react'
import { Link } from 'react-router-dom'
import HomeIcon from '@mui/icons-material/Home';
import AddBoxIcon from '@mui/icons-material/AddBox';
import FavoriteIcon from '@mui/icons-material/Favorite';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import WorkspacePremiumIcon from '@mui/icons-material/WorkspacePremium';
import { useAppSelector } from '../app/hooks';

const Sidebar = () => {
  const currentUser = useAppSelector((state) => state.user);

  return (
    <div className='bg-slate-900 text-slate-100 md:w-60 w-full md:h-screen fixed md:top-0 bottom-0 p-6'>
        
      <Link to="/" className='md:flex md:text-2xl mb-12 mt-2 hidden'>イヌミナティ</Link>

      <ul className='sidebarList md:block flex w-full justify-between'>
        <li className='mb-6'>
          <Link to="/" className='flex text-center text-lg'>
            <HomeIcon className=''/>
            <span className='hidden md:block md:ml-2'>ホーム</span>
          </Link>
        </li>

        <li className='mb-6'>
          <Link to="/new" className='flex text-center text-lg'>
            <AddBoxIcon className=''/>
            <span className='hidden md:block md:ml-2 '>新規作成</span>
          </Link>
        </li>

        <li className='mb-6'>
          <Link to="/" className='flex text-center text-lg' onClick={()=>{alert("犬コンテスト近日開催予定!!")}}>
            <WorkspacePremiumIcon className=''/>
            <span className='hidden md:block md:ml-2'>犬コンテスト</span>
          </Link>
        </li>

        <li className='mb-6'>
          <Link to="/" className='flex text-center text-lg' onClick={()=>{alert("お気に入り実装予定")}}>
            <FavoriteIcon className=''/>
            <span className='hidden md:block md:ml-2'>お気に入り</span>
          </Link>
        </li>

        <li className='mb-6'>
          <Link to={`/user/${currentUser?.id}`} className='flex text-center text-lg'>
            <AccountCircleIcon className=''/>
            <span className='hidden md:block md:ml-2'>プロフィール</span>
          </Link>
        </li>
      </ul>

    </div>
  )
}

export default Sidebar