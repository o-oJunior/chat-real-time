import React from "react"
import ButtonClose from "../button/close"

type WithCloseButton = {
  includeButtonClose: true
  onClose: () => void
}

type WithoutCloseButton = {
  includeButtonClose: false
  onClose?: never
}

type Props = (WithCloseButton | WithoutCloseButton) & {
  children: React.ReactNode
}

const Card = ({ children, includeButtonClose, onClose }: Props) => {
  const styleIncludeButton = includeButtonClose ? "p-2" : ""
  return (
    <div className={`flex border border-gray-300 rounded-lg gap-2 ${styleIncludeButton}`}>
      {includeButtonClose && (
        <div className="fixed">
          <ButtonClose position="left" styleIcon="w-4 h-4" onClose={onClose} />
        </div>
      )}
      <div className="w-full p-5">{children}</div>
    </div>
  )
}

export default Card
