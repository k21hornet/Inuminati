import React, { useState } from 'react'
import Sidebar from '../components/Sidebar'
import { useNavigate } from 'react-router-dom';
import axios, { AxiosError } from 'axios';
import heic2any from 'heic2any';

const NewDog = () => {
  const navigate = useNavigate();
  const [imgUrl, setImgUrl] = useState("");
  const [caption, setCaption] = useState("");
  // ローディング中は投稿ボタンを薄く、かつ状態を表示
  const [isUploading, setIsUploading] = useState<boolean>(false);

  const handleSubmit = async() => {
    const post = {
      img: imgUrl,
      caption: caption,
    }
    try {
      await axios.post(`${process.env.REACT_APP_API_URL}/dogs`, post);
      navigate(-1); // 1ページ戻る
    }catch(err) {
      if (err instanceof AxiosError) {
        // 型ガード
        let msg; // エラーメッセージ
        msg = err.response?.data;
        if(msg.includes("img")) alert("画像をアップロードしてください");
        else alert(msg);
      }
      console.log(err);
    }
  }

  /* 画像の投稿を行う */
  const handleFileChange = async(event: React.ChangeEvent<HTMLInputElement>) => {
    if(event.target.files) {
      const selectedFile = event.target.files[0];

      setIsUploading(true);

      // アップロードする画像のデータ
      const formData = new FormData();
      
      // HEIC形式の画像の場合にはJPEG形式に変換して追加する
      if (selectedFile.type === 'image/heic') {
        console.log("heicだお")
        const jpegFile: any = await heic2any({
          blob: selectedFile,
          toType: 'image/jpeg'
        });

        formData.append('image', jpegFile);
      } else {
        formData.append('image', selectedFile);
      }
  
      try {
        // 画像をアップロード
        const res = await axios.post(`${process.env.REACT_APP_API_URL}/dogs/upload_s3`, formData);
        setImgUrl(res.data.url);
        console.log(res.data.url);

      } catch (err) {
        console.log(err);
      }
      setIsUploading(false);
    }
  };

  return (
    <div className='flex h-screen'>
      <Sidebar/>
      <div className='md:ml-60 w-full bg-slate-50 h-5/6 md:h-full'>
        <div className='w-full h-full p-10 flex flex-col'>
          <h1 className='text-3xl'>新規投稿作成</h1>
          <label className='my-3' htmlFor='file'>
            <span className='bg-slate-700 hover:bg-slate-800 text-white px-2 py-1 rounded'>画像を選択</span>
            <input 
            type='file' 
            id="file" 
            accept='.png, .jpeg, .jpg, .heic'
            style={{display: "none"}}
            onChange={handleFileChange}
            required
          />
          </label>

          <div className='h-3/5 flex justify-center items-center bg-black'>
            {imgUrl!=="" 
            ? <img src={imgUrl} className="max-w-full max-h-full object-cover" alt='読み込み失敗' /> 
            : (isUploading 
              ? <p className='text-slate-400'>画像をアップロード中...</p> 
              : <p className='text-slate-400'>アップした画像が表示されます</p>
            )}
          </div>

          <textarea 
            className='p-2 outline-none border my-3'
            placeholder='キャプションを入力' 
            value={caption} 
            onChange={(e)=>setCaption(e.target.value)}
          />

          <div className='flex'>
          {isUploading 
              ? <div className='w-24 text-xl text-center bg-slate-400 text-white px-2 py-1 rounded'>
                    投稿
                </div>
              : <div 
                  className='w-24 text-xl text-center bg-slate-700 hover:bg-slate-800 text-white px-2 py-1 rounded cursor-pointer' 
                  onClick={handleSubmit}>投稿
                </div>
          }
            <div 
              className='w-28 flex items-center justify-center text-center bg-slate-700 hover:bg-slate-800 text-white ml-4 px-2 py-1 rounded cursor-pointer' 
              onClick={()=>navigate(-1)}>キャンセル
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default NewDog