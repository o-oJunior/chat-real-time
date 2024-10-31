import React from "react"

type AlertType = "error" | "warning" | "success"

interface AlertModalProps {
  type: AlertType
  message: string
  onClose: () => void
}

const icons = {
  error: (
    <svg
      className="w-12 h-12 text-red-500 mx-auto"
      fill="none"
      stroke="currentColor"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M6 18L18 6M6 6l12 12"></path>
    </svg>
  ),
  warning: (
    <svg
      className="w-12 h-12 text-yellow-500 mx-auto"
      fill="none"
      stroke="currentColor"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth="2"
        d="M12 9v2m0 4h.01M12 3a9 9 0 110 18 9 9 0 010-18z"
      ></path>
    </svg>
  ),
  success: (
    <svg
      className="w-12 h-12 text-green-500 mx-auto"
      fill="none"
      stroke="currentColor"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M5 13l4 4L19 7"></path>
    </svg>
  ),
}

const AlertModal: React.FC<AlertModalProps> = ({ type, message, onClose }) => {
  const icon = icons[type]

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg p-6 max-w-sm w-full">
        <div className="text-center mb-4">{icon}</div>
        <div className="text-center mb-6">
          <p className="text-lg font-semibold">{message}</p>
        </div>
        <div className="text-center">
          <button
            onClick={onClose}
            className="px-4 py-2 bg-blue-500 text-white font-semibold rounded hover:bg-blue-600"
          >
            OK
          </button>
        </div>
      </div>
    </div>
  )
}

export default AlertModal
