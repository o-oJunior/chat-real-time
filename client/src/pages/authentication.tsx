import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { addUserData, initialValueUser, useUser } from "../redux/user/slice"
import { useDispatch } from "react-redux"
import { useAppSelector } from "@/redux/hook"
import Sidebar from "@/components/sidebar/sidebar"
import AlertModal, { AlertProps, initialValueAlert } from "@/components/modal/alert"
import WebSocketService from "@/api/v1/websocket"
import { addNotification } from "@/redux/websocket/slice"

const Authentication = ({ children }: { children: React.ReactNode }) => {
  const [alert, setAlert] = useState<AlertProps>(initialValueAlert)
  const { user } = useAppSelector(useUser)
  const dispatch = useDispatch()
  const router = useRouter()

  useEffect(() => {
    const nowMS = new Date().getTime()
    const userLocalStorage = localStorage.getItem("user")
    let userData = { ...initialValueUser, expiresAt: 0 }
    if (userLocalStorage && userLocalStorage !== "") {
      userData = JSON.parse(userLocalStorage)
    }
    const expiresAt = userData.expiresAt * 1000
    if (nowMS >= expiresAt) {
      localStorage.removeItem("user")
      router.push("/login")
    } else {
      const webSocketService = WebSocketService.getInstance()
      webSocketService.connectWebSocket(userData.id, handleMessageWebSocket)
      dispatch(addUserData(userData))
    }
    validatePageNotFound()
  }, [router.pathname])

  const handleMessageWebSocket = (ev: MessageEvent) => {
    const data = JSON.parse(ev.data)
    if (data.type === "notification") {
      dispatch(addNotification(data))
    }
  }

  const validatePageNotFound = () => {
    if (router.pathname == "/404") {
      return router.push("/")
    }
  }
  return (
    <>
      {user.username !== "" ? (
        <div className="flex row w-full">
          <div className="mr-4">{router.pathname !== "/404" && <Sidebar />}</div>
          <div className="py-5 w-full pr-5">{children}</div>
        </div>
      ) : (
        <div>
          <div>{children}</div>
          <AlertModal
            type={alert.type}
            message={alert.message}
            isOpen={alert.modalOpen}
            onClose={() => setAlert(initialValueAlert)}
          />
        </div>
      )}
    </>
  )
}

export default Authentication
