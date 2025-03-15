import React from "react"
import ListItem from "./listItem"
import Card from "../card/card"

export type InviteStatus = "none" | "pending" | "added" | "received"

export interface IUsers {
  username: string
  description: string
  userIdInviter?: string
  inviteStatus: InviteStatus
}

export interface Item extends IUsers {
  id: string
}

export type Options = {
  text: string
  function: (item: Item, index: number) => void
}

export type Dropdown = {
  isVisible: boolean
  indexVisible: undefined | number
  options: Options[]
}

const invite = {
  added:
    "M96 128a128 128 0 1 1 256 0A128 128 0 1 1 96 128zM0 482.3C0 383.8 79.8 304 178.3 304l91.4 0C368.2 304 448 383.8 448 482.3c0 16.4-13.3 29.7-29.7 29.7L29.7 512C13.3 512 0 498.7 0 482.3zM472 200l144 0c13.3 0 24 10.7 24 24s-10.7 24-24 24l-144 0c-13.3 0-24-10.7-24-24s10.7-24 24-24z",
  none: "M96 128a128 128 0 1 1 256 0A128 128 0 1 1 96 128zM0 482.3C0 383.8 79.8 304 178.3 304l91.4 0C368.2 304 448 383.8 448 482.3c0 16.4-13.3 29.7-29.7 29.7L29.7 512C13.3 512 0 498.7 0 482.3zM504 312l0-64-64 0c-13.3 0-24-10.7-24-24s10.7-24 24-24l64 0 0-64c0-13.3 10.7-24 24-24s24 10.7 24 24l0 64 64 0c13.3 0 24 10.7 24 24s-10.7 24-24 24l-64 0 0 64c0 13.3-10.7 24-24 24s-24-10.7-24-24z",
  pending:
    "M224 0a128 128 0 1 1 0 256A128 128 0 1 1 224 0zM178.3 304l91.4 0c20.6 0 40.4 3.5 58.8 9.9C323 331 320 349.1 320 368c0 59.5 29.5 112.1 74.8 144L29.7 512C13.3 512 0 498.7 0 482.3C0 383.8 79.8 304 178.3 304zM352 368a144 144 0 1 1 288 0 144 144 0 1 1 -288 0zm144-80c-8.8 0-16 7.2-16 16l0 64c0 8.8 7.2 16 16 16l48 0c8.8 0 16-7.2 16-16s-7.2-16-16-16l-32 0 0-48c0-8.8-7.2-16-16-16z",
  received:
    "M96 128a128 128 0 1 1 256 0A128 128 0 1 1 96 128zM0 482.3C0 383.8 79.8 304 178.3 304l91.4 0C368.2 304 448 383.8 448 482.3c0 16.4-13.3 29.7-29.7 29.7L29.7 512C13.3 512 0 498.7 0 482.3zM471 143c9.4-9.4 24.6-9.4 33.9 0l47 47 47-47c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9l-47 47 47 47c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0l-47-47-47 47c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l47-47-47-47c-9.4-9.4-9.4-24.6 0-33.9z",
}

export const initialValueDropdown = {
  isVisible: false,
  indexVisible: undefined,
  options: [],
}

type Props = {
  users: IUsers[]
  userIdLogged: string
  text: string
  dropdown: Dropdown
  handleInvite: (item: Item, index: number) => Promise<void>
}

const ListUser = ({ users, userIdLogged, text, dropdown, handleInvite }: Props) => {
  return (
    <ListItem
      list={users}
      text={text}
      styleList="min-w-[300px]"
      renderItem={(item: Item, index: number) => {
        const statusInvite = userIdLogged === item.userIdInviter ? "pending" : "received"
        const statusKey = item.inviteStatus ? statusInvite : "none"
        return (
          <Card>
            <div className="flex flex-row justify-between w-full">
              <div className="flex flex-col gap-1">
                <span>{item.username}</span>
                <span className="text-sm text-gray-400">{item.description}</span>
              </div>
              <div className="flex cursor-pointer" onClick={() => handleInvite(item, index)}>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-8 h-8"
                  stroke="currentColor"
                  viewBox="0 0 640 512"
                >
                  <path fill="#787878" strokeLinecap="round" strokeLinejoin="round" d={invite[statusKey]} />
                </svg>
              </div>
            </div>
            {dropdown.isVisible && dropdown.indexVisible === index && (
              <div className="absolute bg-white shadow-md rounded mt-2 p-2">
                <ul>
                  {dropdown.options.map((value, i) => (
                    <li
                      onClick={() => value.function(item, index)}
                      key={i}
                      className="p-2 hover:bg-gray-200 cursor-pointer"
                    >
                      {value.text}
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </Card>
        )
      }}
    />
  )
}

export default ListUser
