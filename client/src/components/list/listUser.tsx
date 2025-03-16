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

const invite = {
  added: { text: "Excluir contato", style: "" },
  pending: { text: "Cancelar", style: "text-sm text-red-600 hover:text-red-800" },
  none: { text: "Enviar convite", style: "text-blue-600 hover:text-blue-800" },
}

type Props = {
  users: IUsers[]
  userIdLogged: string
  text: string
  handleInvite: (item: Item, index: number, inviteStatus: InviteStatus) => Promise<void>
}

const ListUser = ({ users, userIdLogged, text, handleInvite }: Props) => {
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
                <span className="flex flex-wrap">{item.username}</span>
                <span className="text-sm text-gray-400">{item.description}</span>
              </div>
              {statusKey === "received" ? (
                <div className="flex flex-row gap-2">
                  <svg
                    className="w-7 h-7 cursor-pointer"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 512 512"
                    onClick={() => handleInvite(item, index, "none")}
                  >
                    <path
                      fill="#787878"
                      d="M256 48a208 208 0 1 1 0 416 208 208 0 1 1 0-416zm0 464A256 256 0 1 0 256 0a256 256 0 1 0 0 512zM175 175c-9.4 9.4-9.4 24.6 0 33.9l47 47-47 47c-9.4 9.4-9.4 24.6 0 33.9s24.6 9.4 33.9 0l47-47 47 47c9.4 9.4 24.6 9.4 33.9 0s9.4-24.6 0-33.9l-47-47 47-47c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0l-47 47-47-47c-9.4-9.4-24.6-9.4-33.9 0z"
                    />
                  </svg>
                  <svg
                    className="w-7 h-7 cursor-pointer"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 512 512"
                    onClick={() => handleInvite(item, index, "added")}
                  >
                    <path
                      fill="#76D976"
                      d="M256 48a208 208 0 1 1 0 416 208 208 0 1 1 0-416zm0 464A256 256 0 1 0 256 0a256 256 0 1 0 0 512zM369 209c9.4-9.4 9.4-24.6 0-33.9s-24.6-9.4-33.9 0l-111 111-47-47c-9.4-9.4-24.6-9.4-33.9 0s-9.4 24.6 0 33.9l64 64c9.4 9.4 24.6 9.4 33.9 0L369 209z"
                    />
                  </svg>
                </div>
              ) : (
                <div className="flex flex-col gap-1 text-center">
                  {statusKey === "pending" && <span className="text-gray-600">Solicitado</span>}
                  <span
                    className={`cursor-pointer ${invite[statusKey].style}`}
                    onClick={() => handleInvite(item, index, "none")}
                  >
                    {invite[statusKey].text}
                  </span>
                </div>
              )}
            </div>
          </Card>
        )
      }}
    />
  )
}

export default ListUser
