import API_V1_INVITE from "@/api/v1/invite"
import API_V1_USER from "@/api/v1/user"
import Card from "@/components/card/card"
import ListItem from "@/components/listItem/listItem"
import Alert, { AlertProps, initialValueAlert } from "@/components/modal/alert"
import Modal from "@/components/modal/modal"
import Pagination, { initialValuePagination, TApiPagination } from "@/components/pagination/pagination"
import Search from "@/components/search/search"
import { IResponse } from "@/interfaces/response"
import { useAppSelector } from "@/redux/hook"
import { useUser } from "@/redux/user/slice"
import Head from "next/head"
import React, { ChangeEvent, FormEvent, use, useEffect, useRef, useState } from "react"

type TSearch = {
  contacts: string
  users: string
}

const initialValueSearch = {
  contacts: "",
  users: "",
}

type InviteStatus = "none" | "pending" | "accepted"
interface IUsers {
  username: string
  description: string
  userIdInviter?: string
  inviteStatus: InviteStatus
}

interface Item extends IUsers {
  id: string
}

type Options = {
  text: string
  function: (item: Item, index: number) => void
}

type Dropdown = {
  isVisible: boolean
  indexVisible: undefined | number
  options: Options[]
}

const invite = {
  accepted:
    "M96 128a128 128 0 1 1 256 0A128 128 0 1 1 96 128zM0 482.3C0 383.8 79.8 304 178.3 304l91.4 0C368.2 304 448 383.8 448 482.3c0 16.4-13.3 29.7-29.7 29.7L29.7 512C13.3 512 0 498.7 0 482.3zM472 200l144 0c13.3 0 24 10.7 24 24s-10.7 24-24 24l-144 0c-13.3 0-24-10.7-24-24s10.7-24 24-24z",
  none: "M96 128a128 128 0 1 1 256 0A128 128 0 1 1 96 128zM0 482.3C0 383.8 79.8 304 178.3 304l91.4 0C368.2 304 448 383.8 448 482.3c0 16.4-13.3 29.7-29.7 29.7L29.7 512C13.3 512 0 498.7 0 482.3zM504 312l0-64-64 0c-13.3 0-24-10.7-24-24s10.7-24 24-24l64 0 0-64c0-13.3 10.7-24 24-24s24 10.7 24 24l0 64 64 0c13.3 0 24 10.7 24 24s-10.7 24-24 24l-64 0 0 64c0 13.3-10.7 24-24 24s-24-10.7-24-24z",
  pending:
    "M224 0a128 128 0 1 1 0 256A128 128 0 1 1 224 0zM178.3 304l91.4 0c20.6 0 40.4 3.5 58.8 9.9C323 331 320 349.1 320 368c0 59.5 29.5 112.1 74.8 144L29.7 512C13.3 512 0 498.7 0 482.3C0 383.8 79.8 304 178.3 304zM352 368a144 144 0 1 1 288 0 144 144 0 1 1 -288 0zm144-80c-8.8 0-16 7.2-16 16l0 64c0 8.8 7.2 16 16 16l48 0c8.8 0 16-7.2 16-16s-7.2-16-16-16l-32 0 0-48c0-8.8-7.2-16-16-16z",
}

const initialValueDropdown = {
  isVisible: false,
  indexVisible: undefined,
  options: [],
}

const Contacts = () => {
  const [contacts, setContacts] = useState<any[]>([
    { username: "Olinda", description: "Disponível" },
    { username: "Cloroquina", description: "Disponível" },
  ])
  const { user } = useAppSelector(useUser)
  const [users, setUsers] = useState<IUsers[]>([])
  const [apiPagination, setApiPagination] = useState<TApiPagination>(initialValuePagination)
  const [alert, setAlert] = useState<AlertProps>(initialValueAlert)
  const [isOpen, setIsOpen] = useState<boolean>(false)
  const [search, setSearch] = useState<TSearch>(initialValueSearch)
  const [dropdown, setDropdown] = useState<Dropdown>(initialValueDropdown)
  const [active, setActive] = useState<string>("Adicionados")
  const optionsMenu = ["Adicionados", "Recebidos", "Enviados"]
  const groupContacts = useRef<string>("added")

  useEffect(() => {
    handleContacts(groupContacts.current, "")
  }, [])

  const handleContacts = async (group: string, username: string) => {
    const contacts = await getContacts(1, 20, group, username)
    setContacts(contacts.users)
  }

  const getContacts = async (page: number, limit: number, group: string, username: string) => {
    const v1 = new API_V1_USER()
    const response: IResponse = await v1.getContacts(page, limit, group, username)
    if (response.statusCode !== 200) {
      setAlert({
        message: "Ocorreu um erro inesperado ao obter os usuários, tente novamente mais tarde!",
        type: "error",
        modalOpen: true,
      })
      return
    }
    return response.data
  }

  const handleModal = async () => {
    if (!isOpen) {
      const params = { page: 1, limit: 10 }
      await getUsers(params.page, params.limit)
    } else {
      setSearch(initialValueSearch)
    }
    setIsOpen(!isOpen)
  }

  const handleSearch = (e: FormEvent) => {
    e.preventDefault()
  }

  const handleChangeInput = async (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setSearch({ ...search, [name]: value })
    if (name === "users") {
      await getUsers(1, 10, value)
    } else if (name === "contacts") {
      await handleContacts(groupContacts.current, value)
    }
  }

  const getUsers = async (page?: number, limit?: number, username?: string) => {
    const v1 = new API_V1_USER()
    const response: IResponse = await v1.getUsers(page, limit, username)
    if (response.statusCode !== 200) {
      setAlert({
        message: "Ocorreu um erro inesperado ao obter os usuários, tente novamente mais tarde!",
        type: "error",
        modalOpen: true,
      })
      return
    }
    const users = response.data.users
    if (users.length === 0) {
      setUsers(users)
      return
    }
    setUsers(users)
    setAlert(initialValueAlert)
    setApiPagination({
      currentPage: Number(response.data.page),
      totalPages: Number(response.data.totalPages),
    })
  }

  const handlePageChange = async (page: number) => {
    await getUsers(page, 10, search.users)
  }

  const sendInvite = async (idInvited: string) => {
    const date = new Date().toISOString()
    const status: InviteStatus = "pending"
    const invite = { userIdInvited: idInvited, inviteStatus: status, invitedAt: date }
    const v1 = new API_V1_INVITE()
    const response: IResponse = await v1.sendInvite(invite)
    if (response.statusCode !== 200) {
      setAlert({
        message: "Ocorreu um erro inesperado ao enviar o convite, tente novamente mais tarde!",
        type: "error",
        modalOpen: true,
      })
      return
    }
  }

  const handleInvite = async (item: Item, index: number) => {
    if (!item.inviteStatus) {
      await sendInvite(item.id)
      users[index].inviteStatus = "pending"
      users[index].userIdInviter = user.id
      setUsers([...users])
    } else if (item.inviteStatus === "pending" && item.userIdInviter === user.id) {
      const options = [
        {
          text: "Cancelar Pedido",
          function: (item: Item, index: number) => updateInvite(item, index, "none"),
        },
      ]
      setDropdown({ isVisible: !dropdown.isVisible, indexVisible: index, options: options })
    } else if (item.inviteStatus === "pending") {
      const options = [
        { text: "Aceitar", function: (item: Item, index: number) => updateInvite(item, index, "accepted") },
        { text: "Recusar", function: (item: Item, index: number) => updateInvite(item, index, "none") },
      ]
      setDropdown({ isVisible: !dropdown.isVisible, indexVisible: index, options: options })
    }
  }

  const updateInvite = async (item: Item, index: number, statusInvite: InviteStatus) => {
    const invite = users[index]
    invite.inviteStatus = statusInvite
    const v1 = new API_V1_INVITE()
    const response: IResponse = await v1.updateStatusInvite(invite)
    if (response.statusCode !== 200) {
      setAlert({
        message: "Ocorreu um erro inesperado ao enviar o convite, tente novamente mais tarde!",
        type: "error",
        modalOpen: true,
      })
      return
    }
    users[index].inviteStatus = statusInvite
    setUsers([...users])
    setDropdown(initialValueDropdown)
  }

  const handleOptionMenu = (item: "Adicionados" | "Recebidos" | "Enviados") => {
    const groups = {
      Adicionados: "added",
      Recebidos: "received",
      Enviados: "sent",
    }
    setActive(item)
    groupContacts.current = groups[item]
    handleContacts(groupContacts.current, "")
  }

  return (
    <>
      <Head>
        <title>Contatos - Chat</title>
      </Head>
      <div>
        <div className="flex row items-center justify-between w-full px-20 py-10">
          <Search
            handleSearch={handleSearch}
            handleChangeInput={handleChangeInput}
            textPlaceholder="contatos"
            nameInput="contacts"
            query={search.contacts}
          />
          <button
            type="submit"
            className="bg-gray-500 text-white py-2 px-4 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600"
            onClick={handleModal}
          >
            Adicionar contato
          </button>
        </div>
        <div className="flex justify-center space-x-4 p-4 rounded-lg">
          {optionsMenu.map((item) => (
            <button
              key={item}
              className={`px-6 py-2 rounded-md transition-colors duration-300 ${
                active === item ? "bg-blue-500 text-white" : "bg-gray-200 hover:bg-gray-300"
              }`}
              onClick={() => handleOptionMenu(item as "Adicionados" | "Recebidos" | "Enviados")}
            >
              {item}
            </button>
          ))}
        </div>
        <Modal isOpen={isOpen} onClose={handleModal}>
          <div className="flex flex-col gap-5 h-full">
            <Search
              handleSearch={handleSearch}
              handleChangeInput={handleChangeInput}
              textPlaceholder="nome de usuário"
              nameInput="users"
              query={search.users}
            />
            <div className="flex-1 overflow-auto gap-5">
              <ListItem
                list={users}
                text="usuário"
                renderItem={(item: Item, index: number) => {
                  const statusKey = item.inviteStatus ? item.inviteStatus : "none"
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
                            <path
                              fill="#787878"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              d={invite[statusKey]}
                            />
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
            </div>
            <Pagination
              currentPage={apiPagination!.currentPage}
              totalPages={apiPagination!.totalPages}
              handlePageChange={handlePageChange}
            />
          </div>
        </Modal>
        {alert!.modalOpen && (
          <Alert type={alert!.type} message={alert!.message} onClose={() => setAlert(initialValueAlert)} />
        )}
        <div className="flex flex-row justify-between w-full px-20">
          <div className="flex-1 border border-gray-300 rounded-lg p-10">
            <ListItem list={contacts} text="contato" renderItem={(item) => <div>{item.username}</div>} />
          </div>
        </div>
      </div>
    </>
  )
}

export default Contacts
