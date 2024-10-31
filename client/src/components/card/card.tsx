import React, { ReactNode } from "react"

const Card = ({ children }: { children: ReactNode }) => {
  return <div className="flex border border-gray-300 rounded-lg p-5">{children}</div>
}

export default Card
