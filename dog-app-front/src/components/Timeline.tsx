import axios from 'axios'
import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { Dog } from '../Type';

const Timeline = (props: any) => {
  const [dogPosts, setDogPosts] = useState<Dog[]>([]);
  const {user} = props; // nullなら全ての投稿、falseならユーザーの投稿

  useEffect(() => {
    const getAllPosts = async() => {
      const res = user
        ? await axios.get(`${process.env.REACT_APP_API_URL}/dogs/${user}/user`)
        : await axios.get(`${process.env.REACT_APP_API_URL}/dogs/allusers`) 

      if(res.status===200)
        setDogPosts(res.data);
    }
    getAllPosts();
  }, [user])

  return (
    <div className='w-full md:px-8 md:py-4 md:p-2 p-0.5'>
      <div className='grid md:grid-cols-4 grid-cols-3 gap-0'>
        {dogPosts.map((post) => (
          <div className='md:m-0.5 m-px'>
            <Link to={`/dogs/${post.id}`}>
              <img src={`${post?.img}`} alt="画像読み込み失敗" className="aspect-square object-cover"/>
              
            </Link>
          </div>
        ))}
      </div>
    </div>
  )
}

export default Timeline