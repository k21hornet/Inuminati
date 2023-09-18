import React, { useEffect, useState } from 'react'
import { useAppDispatch, useAppSelector } from '../app/hooks';
import axios from 'axios';
import { logout } from '../features/userSlice';
import { User } from '../Type';
import { useParams } from 'react-router-dom';
import heic2any from 'heic2any';
import AddAPhotoIcon from '@mui/icons-material/AddAPhoto';

const UserProfile = () => {
  const currentUser = useAppSelector((state) => state.user);
  const dispatch = useAppDispatch();
  const [profileUser, setProfileUser] = useState<User>();
  const [userIcon, setUserIcon] = useState(""); // アイコン更新時にuseeffect
  const { id } = useParams();

  useEffect(() => {
    //投稿したユーザー情報を得る
    const getPostUser = async (id:string|undefined) => {
      const res = await axios.get(`${process.env.REACT_APP_API_URL}/${id}`);
      setProfileUser(res.data);
    };
    getPostUser(id)
    
  }, [id, userIcon]);

  const userLogout =async () => {
    if(window.confirm("ログアウトしますか？")) {
      try {
        await axios.post(`${process.env.REACT_APP_API_URL}/logout`);
        dispatch(logout());
      }catch(err) {
        console.log(err);
      }
    }
  }

  /* アイコン画像の投稿を行う */
  const handleFileChange = async(event: React.ChangeEvent<HTMLInputElement>) => {
    if(event.target.files) {
      const selectedFile = event.target.files[0];

      // アップロードする画像のデータ
      const formData = new FormData();
      
      // HEIC形式の画像の場合にはJPEG形式に変換して追加する
      if (selectedFile.type === 'image/heic') {
        console.log("heicだお")
        const jpegFile: any = await heic2any({
          blob: selectedFile,
          toType: 'image/jpeg'
        });

        formData.append('icon', jpegFile);
      } else {
        formData.append('icon', selectedFile);
      }
  
      try {
        // 画像をアップロード
        const res = await axios.post(`${process.env.REACT_APP_API_URL}/${id}/upicon`, formData);
        setUserIcon(res.data.icon);

      } catch (err) {
        console.log(err);
      }
    }
  };

  return (
    <div className='w-full flex flex-col md:px-12 px-2 pt-4 m-2'>
      <div className='w-full flex items-center'>
        <div className='w-28 h-28 bg-slate-200 rounded-full mr-8 flex justify-center items-center'>
          <img src={profileUser?.icon} className="w-full h-full object-cover rounded-full overflow-hidden" alt='' />
        </div>
        {currentUser?.id === profileUser?.id ? (
            <label className='my-3 relative top-9 right-16' htmlFor='file'>
            <span className='bg-slate-600 cursor-pointer hover:bg-slate-800 p-1 text-white rounded-full flex items-center'>
              <AddAPhotoIcon/>
            </span>
            <input 
            type='file' 
            id="file" 
            accept='.png, .jpeg, .jpg, .heic'
            style={{display: "none"}}
            onChange={handleFileChange}
            required
          />
          </label>
          ) : (
            <></>
          ) }
        

        <div className='w-1/2'>
          <div className='flex'>
            <h1 className='text-2xl'>{profileUser?.username}</h1>
          </div>
          {currentUser?.id === profileUser?.id ? (
            <div className="" onClick={()=>alert("Comming soon!")}>
              ユーザー設定
            </div>
          ) : (
            <div onClick={()=>alert("Comming soon!")}>
              フォローする
            </div>
          ) }
          {currentUser?.id === profileUser?.id ? (
            <div className="w-24 bg-slate-700 hover:bg-slate-800 text-white px-2 py-1 rounded cursor-pointer" onClick={userLogout}>
              ログアウト
            </div>
          ) : (
            <></>
          ) }
        </div>
        
      </div>
    </div>
  )
}

export default UserProfile