import { createSlice, PayloadAction } from "@reduxjs/toolkit"

export interface IWebSocket {
  message: string
  type: string
}

export interface IWebSocketData {
  notification: IWebSocket
}

export const initialValueWebSocket: IWebSocket = {
  message: "",
  type: "",
}

const initialState: IWebSocketData = {
  notification: initialValueWebSocket,
}

export const webSocketSlice = createSlice({
  name: "websocket",
  initialState,
  reducers: {
    addNotification: (state, action: PayloadAction<IWebSocket>) => {
      state.notification = action.payload
    },
    clearNotification: (state) => {
      state.notification = initialValueWebSocket
    },
  },
})

export const { addNotification, clearNotification } = webSocketSlice.actions
export const useWebSocket = (state: any) => state.websocket as IWebSocketData
