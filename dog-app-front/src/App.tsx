import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import './App.css';
import Home from './pages/Home';
import Signup from './pages/Signup';
import Login from './pages/Login';
import { useAppDispatch, useAppSelector } from './app/hooks';
import { login, logout } from "./features/userSlice";
import Loading from './components/Loading';
import Profile from './pages/Profile';
import NewDog from './pages/NewDog';
import Post from './pages/Post';
import Favorite from './pages/Favorite';

function App() {
  const currentUser = useAppSelector((state) => state.user);
  const dispatch = useAppDispatch();
  const [isLoadingGetCUser, setIsLoadingGetCUser] = useState<boolean>(true);

  // アプリを起動したときに/csrfにアクセスし、Tokenをもらう
  useEffect(()=> {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      //axios.get<型>("path")
      const {data} = await axios.get(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      // headerの名前をつけてtokenを付与
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }

    const getUserId = async () => {
      //axios.get<型>("path")
      const {data} = await axios.get(
        `${process.env.REACT_APP_API_URL}/user`
      )
      if(data.user_id) {
        dispatch(login({
          id: data.user_id,
        }))
      }else {
        dispatch(logout());
      }
      setIsLoadingGetCUser(false);
    }
    getCsrfToken()
    getUserId()
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={isLoadingGetCUser ? (<Loading/>) : (
          currentUser ? <Home/> : <Navigate to="/login"/>
        )}/>
        <Route path="/new" element={isLoadingGetCUser ? (<Loading/>) : (
          currentUser ? <NewDog/> : <Navigate to="/login"/>
        )}/>
        <Route path="/dogs/:id" element={isLoadingGetCUser ? (<Loading/>) : (
          currentUser ? <Post/> : <Navigate to="/login"/>
        )}/>
        <Route path="/likes" element={isLoadingGetCUser ? (<Loading/>) : (
          currentUser ? <Favorite/> : <Navigate to="/login"/>
        )}/>
        <Route path="/user/:id" element={isLoadingGetCUser ? (<Loading/>) : (
          currentUser ? <Profile/> : <Navigate to="/login"/>
        )}/>
        <Route path="/signup" element={isLoadingGetCUser ? (<Loading/>) : (
          currentUser ? <Navigate to="/"/> : <Signup/>
        )}/>
        <Route path="/login" element={isLoadingGetCUser ? (<Loading/>) : (
          currentUser ? <Navigate to="/"/> : <Login/>
        )}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
