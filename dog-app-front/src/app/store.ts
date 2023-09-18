import { configureStore } from "@reduxjs/toolkit";
import userReducer from "../features/userSlice";

// reduxのstore
export const store = configureStore({
  reducer: userReducer,
})

//dispatchの型
export type AppDispatch = typeof store.dispatch
export type RootState = ReturnType<typeof store.getState>