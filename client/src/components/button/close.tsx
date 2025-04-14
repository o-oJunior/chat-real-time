import React from "react"

type Props = {
  position: string
  styleIcon: string
  onClose: () => void
}

const ButtonClose = ({ position, styleIcon, onClose }: Props) => {
  return (
    <button
      onClick={onClose}
      className={`text-gray-600 hover:text-gray-900 focus:outline-none float-${position}`}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        className={styleIcon}
      >
        <line x1="18" y1="6" x2="6" y2="18" />
        <line x1="6" y1="6" x2="18" y2="18" />
      </svg>
    </button>
  )
}

export default ButtonClose
