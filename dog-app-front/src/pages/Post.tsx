import React, { useEffect, useState } from 'react'
import Sidebar from '../components/Sidebar'
import axios from 'axios';
import { useAppSelector } from '../app/hooks';
import { useNavigate, useParams } from 'react-router-dom';
import { Dog, User } from '../Type';
import { Link } from 'react-router-dom';

const Post = () => {
  const currentUser = useAppSelector((state) => state.user);
  const { id } = useParams();
  const [post, setPost] = useState<Dog>();
  const [postUser, setPostUser] = useState<User>();
  const [postDate, setPostDate] = useState<Date>()
  const nav = useNavigate();

  const deleteDog = async() => {
    if (currentUser?.id === post?.user_id) {
      const check = window.confirm("削除しますか？");
      if (check) {
        await axios.delete(`${process.env.REACT_APP_API_URL}/dogs/${id}`);
        nav(-1);
      }
    } else {
      alert("自分の投稿のみ削除できます。");
    }
  }

  useEffect(() => {
    // 投稿を得る
    const getPost = async () => {
      const res = await axios.get(`${process.env.REACT_APP_API_URL}/dogs/${id}`);
      setPost(res.data);
      console.log(res.data)
      getPostUser(res.data.user_id); // 投稿したユーザーは誰か
      setPostDate(new Date(res.data.created_at))
    };
    getPost();
    //投稿したユーザー情報を得る
    const getPostUser = async (id: number) => {
      const res = await axios.get(`${process.env.REACT_APP_API_URL}/${id}`);
      setPostUser(res.data);
    };
    
  }, [id]);

  return (
    <div className='flex h-screen'>
      <Sidebar/>
      <div className='z-10 w-full flex justify-center p-4 bg-black/50' onClick={(e)=>nav(-1) /* 外側をクリックしてブラウザバック */}>
        <div className='w-11/12 flex md:flex-row flex-col-reverse justify-between bg-slate-100' onClick={(e)=>e.stopPropagation() /* ブラウザバックを防ぐ */}>

          <div className='md:w-4/12 md:h-full h-1/4 p-4'>
            <Link to={`/user/${post?.user_id}`}>
              <div className='flex items-center py-2'>
              <div className='w-12 h-12 bg-slate-200 rounded-full mr-2 flex justify-center items-center'>
                <img src={postUser?.icon} className="w-full h-full rounded-full object-cover" alt='' />
              </div>
                <h1 className='text-2xl'>{postUser?.username}</h1>
              </div>
            </Link>

            <hr className='bg-slate-300 h-0.5 mb-4'/>

            <div className=' text-slate-400 text-sm'>作成日 {postDate?.toLocaleDateString()}</div>
            <p className='whitespace-pre'>{post?.caption}</p>
            <div className='text-blue-800' onClick={deleteDog}>投稿を削除する</div>
          </div>

          <div className='md:w-8/12 md:h-full h-3/4 flex justify-center items-center bg-black'>
            <img 
              src={`${post?.img}`} alt="画像読み込み失敗"
              className='max-w-full max-h-full object-cover'
            />
          </div>
        </div>
      </div>
    </div>
  )
}

export default Post