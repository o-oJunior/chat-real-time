import React from "react"
import ReactDOM from "react-dom"

type positionX = "left" | "center" | "right"
type positionY = "top" | "center" | "bottom"
type positionClose = "left" | "right"

type Props = {
  children: React.ReactNode
  isOpen: boolean
  fullScreen?: boolean
  xPosition?: positionX
  yPosition?: positionY
  positionClose?: positionClose
  onClose?: () => void
}

const Modal = ({
  isOpen,
  xPosition = "center",
  yPosition = "center",
  positionClose = "left",
  fullScreen = true,
  onClose,
  children,
}: Props) => {
  if (!isOpen) return null
  const xClasses = {
    left: "left-0 ml-2",
    center: "left-1/2 transform -translate-x-1/2",
    right: "right-0 mr-2",
  }

  const yClasses = {
    top: "top-0 mt-2",
    center: "top-1/2 transform -translate-y-1/2",
    bottom: "bottom-0 mb-2",
  }
  let classAddittional = "inset-0 flex items-center justify-center bg-black bg-opacity-50"
  let classIcon = "w-6 h-6"
  let classChildren = "h-full pt-5 px-5 pb-5"
  let padding = "p-6"
  if (!fullScreen) {
    classAddittional = `border border-gray-300 ${xClasses[xPosition]} ${yClasses[yPosition]}`
    classIcon = "w-4 h-4"
    classChildren = "h-full py-8 px-20"
    padding = "p-2"
  }

  return ReactDOM.createPortal(
    <div className={`fixed ${classAddittional}`}>
      <div className={`bg-white h-5/6 rounded-lg shadow-lg max-w-lg w-full ${padding}`}>
        <button
          onClick={onClose}
          className={`text-gray-600 hover:text-gray-900 focus:outline-none float-${positionClose}`}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            className={classIcon}
          >
            <line x1="18" y1="6" x2="6" y2="18" />
            <line x1="6" y1="6" x2="18" y2="18" />
          </svg>
        </button>
        <div className={classChildren}>{children}</div>
      </div>
    </div>,
    document.body,
  )
}

export default Modal
