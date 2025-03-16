import React from "react"

type Props = {
  text: string
  list: any[]
  styleList?: string
  renderItem: (item: any, index: number) => React.ReactNode
}

const ListItem = ({ text, list, styleList, renderItem }: Props) => {
  return (
    <div>
      {list.length > 0 ? (
        <ul className="flex flex-row flex-wrap w-full gap-2">
          {list.map((item, index) => (
            <li className={`flex-1 justify-center ${styleList}`} key={index}>
              {renderItem(item, index)}
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
