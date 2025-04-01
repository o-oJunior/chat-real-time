import { addNotification } from "@/redux/websocket/slice"
import { useDispatch } from "react-redux"

const MESSAGE_ERROR = {
  statusCode: 500,
  error: "Erro na conexão com o servidor, tente novamente mais tarde!",
}

type wsOnMessage = ((this: WebSocket, ev: MessageEvent) => any) | null

class WebSocketService {
  private static instance: WebSocketService
  private readonly BASE_URL: string
  private socket: WebSocket | null = null
  private reconnectInterval = 5000
  private isManualClose = false

  constructor() {
    if (!process.env.NEXT_PUBLIC_URL_WS) {
      throw new Error("A variável de ambiente NEXT_PUBLIC_URL_WS não está definida.")
    }
    this.BASE_URL = process.env.NEXT_PUBLIC_URL_WS
  }

  public static getInstance(): WebSocketService {
    if (!WebSocketService.instance) {
      WebSocketService.instance = new WebSocketService()
    }
    return WebSocketService.instance
  }

  private isConnected(): boolean {
    return this.socket !== null && this.socket.readyState === WebSocket.OPEN
  }

  public connectWebSocket(userID: string, onMessageCallback: wsOnMessage) {
    if (this.isConnected()) {
      console.log("Já tem uma conexão ativa com websocket")
      return
    }
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      console.warn("WebSocket já está conectado.")
      return
    }

    this.isManualClose = false
    this.socket = new WebSocket(`${this.BASE_URL}?userId=${userID}`)

    this.socket.onopen = () => {
      console.log(`Conectado ao WebSocket`)
    }

    this.socket.onmessage = onMessageCallback

    this.socket.onerror = (error) => {
      console.error("Erro no WebSocket:", error)
    }

    this.socket.onclose = () => {
      console.warn("Conexão WebSocket fechada.")
      if (!this.isManualClose) {
        console.log(`Tentando reconectar em ${this.reconnectInterval / 1000}s...`)
        setTimeout(() => this.connectWebSocket(userID, onMessageCallback), this.reconnectInterval)
      }
    }
  }

  public sendMessage(message: string): void {
    if (this.isConnected()) {
      this.socket?.send(message)
    } else {
      console.warn("WebSocket não está conectado.")
    }
  }

  public closeWebSocket() {
    if (this.socket) {
      this.isManualClose = true
      this.socket.close()
      console.log("🔌 WebSocket fechado manualmente.")
    }
  }

  public getSocket(): WebSocket | null {
    return this.socket
  }
}

export default WebSocketService
