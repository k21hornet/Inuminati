// 公式ドキュメンと参照
import { useSelector } from "react-redux";
import { TypedUseSelectorHook, useDispatch } from "react-redux";
import { AppDispatch, RootState } from "./store";

// 型をつける
export const useAppDispatch: () => AppDispatch = useDispatch
export const useAppSelector:  TypedUseSelectorHook<RootState> = useSelector