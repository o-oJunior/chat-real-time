import { createSlice, PayloadAction } from "@reduxjs/toolkit"

export interface IUser {
  id: number
  username: string
  firstName: string
  lastName: string
  email: string
  status: string
  description: string
  createAt: string
  expiresAt: number
}

export interface IUserData {
  user: IUser
}

export const initialValueUser: IUser = {
  id: 0,
  username: "",
  firstName: "",
  lastName: "",
  email: "",
  status: "",
  description: "",
  createAt: "",
  expiresAt: 0,
}

const initialState: IUserData = {
  user: initialValueUser,
}

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    addUserData: (state, action: PayloadAction<IUser>) => {
      state.user = action.payload
    },
    userLogout: (state) => {
      state.user = initialValueUser
    },
  },
})

export const { addUserData, userLogout } = userSlice.actions
export const useUser = (state: any) => state.user as IUserData
