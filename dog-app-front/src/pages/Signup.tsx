import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
import axios, { AxiosError } from "axios";
import { useAppDispatch } from "../app/hooks";
import { login } from "../features/userSlice";

const Signup = () => {
  const navigate = useNavigate();
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [passwordConfirmation, setPasswordConfirmation] = useState<string>("");
  const dispatch = useAppDispatch();
  const [loginButtonText, setLoginButtonText] = useState<string>("新規登録");
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const userSignup = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (password !== passwordConfirmation) {
      alert("確認用パスワードが一致しません");
    } else {
      setIsLoading(true);

      const user = {
        username: username,
        email: email,
        password: password,
      };

      try {
        const res = await axios.post(
          `${process.env.REACT_APP_API_URL}/signup`,
          user
        );
        if (res.status === 201) {
          await axios.post(`${process.env.REACT_APP_API_URL}/login`, user);
          const {data} = await axios.get(`${process.env.REACT_APP_API_URL}/user`);
          dispatch(
            login({
              id: data.user_id,
            })
          );
          navigate("/");
        }
      } catch (err) {
        console.log(err);
        // エラー
        if (err instanceof AxiosError) {
          // 型ガード
          let msg = ""; // エラーメッセージ
          if (err.response?.data.message) {
            msg = err.response?.data.message;
          } else {
            // 403はString
            msg = err.response?.data;
          }
          if(msg.includes("csrf")) alert("csrfトークンが無効です。\nリロードしてやり直してください。")
          else if(msg.includes("valid email format")) alert("emailのフォーマットが正しくありません")
          else if(msg.includes("Duplicate")) alert("emailはすでに使用されています")
          else if(msg.includes("limited min 6 max 30 char")) alert("パスワードは6文字以上30文字以下です")
          else alert(msg);
        }
      }
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (isLoading) setLoginButtonText("ログイン中...");
    else setLoginButtonText("新規登録");
  }, [isLoading]);

  return (
    <div className="flex justify-center items-center flex-col min-h-screen bg-slate-50">
      <div className="md:shadow-lg px-6 py-6 md:bg-white mb-24 w-full md:w-112">
        <form className="flex flex-col" onSubmit={userSignup}>
          <p className="text-3xl text-center mb-3">新規登録</p>
          <label className="text-xs">ユーザー名</label>
          <input
            type="text"
            className="border w-full mb-3 px-2 py-1 text-lg rounded-sm"
            value={username}
            required
            onChange={(e) => {
              setUsername(e.target.value);
            }}
          />
          <label className="text-xs">メールアドレス</label>
          <input
            type="email"
            className="border w-full mb-3 px-2 py-1 text-lg rounded-sm"
            value={email}
            required
            onChange={(e) => {
              setEmail(e.target.value);
            }}
          />
          <label className="text-xs">パスワード</label>
          <input
            type="password"
            className="border w-full mb-3 px-2 py-1 text-lg rounded-sm"
            value={password}
            required
            onChange={(e) => {
              setPassword(e.target.value);
            }}
          />
          <label className="text-xs">パスワード確認用</label>
          <input
            type="password"
            className="border w-full mb-3 px-2 py-1 text-lg rounded-sm"
            value={passwordConfirmation}
            required
            onChange={(e) => {
              setPasswordConfirmation(e.target.value);
            }}
          />
          <button 
            className="text-xl mb-3 bg-slate-700 hover:bg-slate-800 text-white px-2 py-1 rounded" 
            type="submit"
          >{loginButtonText}</button>
        </form>
        <p>
          <Link className="text-blue-700 underline" to="/login">アカウントをお持ちの方はこちら</Link>
        </p>
        <p>
          ※ChromeまたはSafariで開いてください
        </p>
      </div>
    </div>
  );
};

export default Signup;
