import Modal from "@/components/modal/modal"
import { useAppSelector } from "@/redux/hook"
import { clearNotification, useWebSocket } from "@/redux/websocket/slice"
import React, { useEffect, useState } from "react"
import { useDispatch } from "react-redux"

const Notification = ({ children }: { children: React.ReactNode }) => {
  const [message, setMessage] = useState("")
  const [isModalOpen, setIsModalOpen] = useState(false)
  const { notification } = useAppSelector(useWebSocket)
  const dispatch = useDispatch()

  useEffect(() => {
    if (notification.message !== "") {
      setMessage(notification.message)
      setIsModalOpen(true)
      setInterval(() => {
        setMessage("")
        setIsModalOpen(false)
        dispatch(clearNotification())
      }, 5000)
    }
  }, [notification])

  const onClose = () => {
    setIsModalOpen(false)
  }

  return (
    <div>
      <div>{children}</div>
      <Modal
        isOpen={isModalOpen}
        fullScreen={false}
        xPosition="right"
        yPosition="bottom"
        positionClose="right"
        onClose={onClose}
      >
        {message !== "" && (
          <div>
            <span dangerouslySetInnerHTML={{ __html: message }}></span>
          </div>
        )}
      </Modal>
    </div>
  )
}

export default Notification
