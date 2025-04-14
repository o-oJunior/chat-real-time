import ButtonClose from "@/components/button/close"
import Card from "@/components/card/card"
import Modal from "@/components/modal/modal"
import { useAppSelector } from "@/redux/hook"
import { clearNotification, useWebSocket } from "@/redux/websocket/slice"
import React, { useEffect, useState } from "react"
import { useDispatch } from "react-redux"

interface NotificationItem {
  message: string
}

const Notification = ({ children }: { children: React.ReactNode }) => {
  const [messages, setMessages] = useState<NotificationItem[]>([])
  const { notification } = useAppSelector(useWebSocket)
  const dispatch = useDispatch()

  useEffect(() => {
    if (notification.message !== "") {
      setMessages((prev) => {
        const updated = [{ message: notification.message }, ...prev]
        return updated.slice(0, 3)
      })

      setTimeout(() => {
        setMessages((prev) => prev.filter((item) => item.message !== notification.message))
        dispatch(clearNotification())
      }, 5000)
    }
  }, [notification])

  const onClose = (index: number) => {
    messages.splice(index, 1)
    setMessages([...messages])
  }

  return (
    <div>
      <div>{children}</div>
      {messages.length > 0 && (
        <div className="fixed flex flex-col gap-2 right-0 mr-2 bottom-0 mb-2">
          {messages.map((item, index) => (
            <div className="flex flex-wrap w-64" key={index}>
              <Card includeButtonClose={true} onClose={() => onClose(index)}>
                <span dangerouslySetInnerHTML={{ __html: item.message }}></span>
              </Card>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default Notification
