import { configureStore } from "@reduxjs/toolkit"
import userReducer from "./user/reducer"
import webSocketReducer from "./websocket/reducer"

const store = configureStore({
  reducer: {
    user: userReducer,
    websocket: webSocketReducer,
  },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export default store
