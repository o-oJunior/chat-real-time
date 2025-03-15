import React from "react"
import ReactDOM from "react-dom"

type Props = {
  children: React.ReactNode
  isOpen?: boolean
  onClose?: () => void
}
const Modal = ({ isOpen, onClose, children }: Props) => {
  if (!isOpen) return null
  return ReactDOM.createPortal(
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white h-5/6 p-6 rounded-lg shadow-lg max-w-lg w-full">
        <button onClick={onClose} className="text-gray-600 hover:text-gray-900 focus:outline-none float-left">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
            className="w-6 h-6"
          >
            <line x1="18" y1="6" x2="6" y2="18" />
            <line x1="6" y1="6" x2="18" y2="18" />
          </svg>
        </button>
        <div className="h-full pt-12 px-5 pb-5 z-50">{children}</div>
      </div>
    </div>,
    document.body,
  )
}

export default Modal
