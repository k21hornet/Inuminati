import React from 'react'
import Sidebar from '../components/Sidebar'

const Favorite = () => {
  return (
    <div className='flex'>
      <Sidebar/>
      <div className='md:ml-60 w-full bg-slate-50'>
        Favorite
      </div>
    </div>
  )
}

export default Favorite