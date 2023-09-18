// slice... storeの中のstate + reducer + action creator、redux toolkitより。
import { createSlice } from "@reduxjs/toolkit";
import { InitialUserState } from "../interface/types";

const initialState: InitialUserState = {
  user: null,
}

// スライス作成、
// actionはdispatchによって送信されたもの
export const userSlice = createSlice({
  name: "user",
  initialState: initialState,
  reducers: {
    login: (state, action) => {
      state.user = action.payload
    },
    logout: (state) => {
      state.user = null
    },
  },
})

export const {login, logout} = userSlice.actions
export default userSlice.reducer