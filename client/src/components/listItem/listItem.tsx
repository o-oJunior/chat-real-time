import React from "react"

type Props = {
  text: string
  list: any[]
  renderItem: (item: any) => React.ReactNode
}

const ListItem = ({ text, list, renderItem }: Props) => {
  return (
    <div>
      {list.length > 0 ? (
        <ul className="flex flex-col w-full gap-2">
          {list.map((item, index) => (
            <li className="flex-1 w-full justify-center" key={index}>
              {renderItem(item)}
            </li>
          ))}
        </ul>
      ) : (
        <span className="flex text-gray-400 justify-center">Nenhum {text} encontrado!</span>
      )}
    </div>
  )
}

export default ListItem
