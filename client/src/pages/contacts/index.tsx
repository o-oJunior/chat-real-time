import API_V1_INVITE from "@/api/v1/invite"
import API_V1_USER from "@/api/v1/user"
import ListUser, { InviteStatus, Item, IUsers } from "@/components/list/listUser"
import Alert, { AlertProps, initialValueAlert } from "@/components/modal/alert"
import Modal from "@/components/modal/modal"
import Pagination, { initialValuePagination, TApiPagination } from "@/components/pagination/pagination"
import Search from "@/components/search/search"
import { IResponse } from "@/interfaces/response"
import { useAppSelector } from "@/redux/hook"
import { useUser } from "@/redux/user/slice"
import Head from "next/head"
import React, { ChangeEvent, FormEvent, useEffect, useRef, useState } from "react"

type Group = "added" | "received" | "sent"

type TSearch = {
  contacts: string
  users: string
}

const initialValueSearch = {
  contacts: "",
  users: "",
}

const Contacts = () => {
  const [contacts, setContacts] = useState<any[]>([
    { username: "Olinda", description: "Disponível" },
    { username: "Cloroquina", description: "Disponível" },
  ])
  const { user } = useAppSelector(useUser)
  const [users, setUsers] = useState<IUsers[]>([])
  const [invitePagination, setInvitePagination] = useState<TApiPagination>(initialValuePagination)
  const [contactPagination, setContactPagination] = useState<TApiPagination>(initialValuePagination)
  const [alert, setAlert] = useState<AlertProps>(initialValueAlert)
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false)
  const [search, setSearch] = useState<TSearch>(initialValueSearch)
  const [active, setActive] = useState<string>("Adicionados")
  const optionsMenu = ["Adicionados", "Recebidos", "Enviados"]
  const groupContacts = useRef<Group>("added")

  useEffect(() => {
    getContactsByUseEffect(groupContacts.current, "")
  }, [])

  const getContactsByUseEffect = async (group: string, username: string) => {
    await getContacts(1, 20, group, username)
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
    setContacts(response.data.users)
    setContactPagination({
      currentPage: Number(response.data.page),
      totalPages: Number(response.data.totalPages),
    })
    return response.data
  }

  const handleModal = async () => {
    if (!isModalOpen) {
      const params = { page: 1, limit: 10 }
      await getUsers(params.page, params.limit)
    } else {
      setSearch(initialValueSearch)
    }
    setIsModalOpen(!isModalOpen)
  }

  const handleChangeInput = async (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setSearch({ ...search, [name]: value })
    if (name === "users") {
      await getUsers(1, 10, value)
    } else if (name === "contacts") {
      await getContacts(1, 20, groupContacts.current, value)
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
    setInvitePagination({
      currentPage: Number(response.data.page),
      totalPages: Number(response.data.totalPages),
    })
  }

  const handlePageChange = async (page: number, list: "users" | "contacts") => {
    const update = {
      contacts: async () => await getContacts(page, 20, groupContacts.current, search.contacts),
      users: async () => await getUsers(page, 10, search.users),
    }
    if (update[list]) {
      await update[list]()
    }
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

  const handleInvite = async (item: Item, index: number, inviteStatus: InviteStatus) => {
    if (!item.inviteStatus && isModalOpen) {
      await sendInvite(item.id)
      users[index].inviteStatus = "pending"
      users[index].userIdInviter = user.id
      setUsers([...users])
      return
    }
    updateInvite(item, index, inviteStatus)
  }

  const updateInvite = async (item: Item, index: number, statusInvite: InviteStatus) => {
    item.inviteStatus = statusInvite
    const v1 = new API_V1_INVITE()
    const response: IResponse = await v1.updateStatusInvite(item)
    if (response.statusCode !== 200) {
      setAlert({
        message: "Ocorreu um erro inesperado ao enviar o convite, tente novamente mais tarde!",
        type: "error",
        modalOpen: true,
      })
      return
    }
    if (isModalOpen) {
      users[index].inviteStatus = statusInvite
      setUsers([...users])
    } else {
      const listContacts = contacts.filter((contact) => contact.id !== item.id)
      setContacts(listContacts)
    }
  }

  const handleOptionMenu = async (item: "Adicionados" | "Recebidos" | "Enviados") => {
    const groups = {
      Adicionados: "added",
      Recebidos: "received",
      Enviados: "sent",
    }
    setActive(item)
    groupContacts.current = groups[item] as Group
    await getContacts(1, 20, groupContacts.current, "")
  }

  return (
    <>
      <Head>
        <title>Contatos - Chat</title>
      </Head>
      <div>
        <div className="flex row items-center justify-between w-full px-20 py-10">
          <Search
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
        <Modal isOpen={isModalOpen} onClose={handleModal}>
          <div className="flex flex-col gap-5 h-full">
            <Search
              handleChangeInput={handleChangeInput}
              textPlaceholder="nome de usuário"
              nameInput="users"
              query={search.users}
            />
            <div className="flex-1 overflow-auto gap-5">
              <ListUser users={users} userIdLogged={user.id} text="usuário" handleInvite={handleInvite} />
            </div>
            <Pagination
              currentPage={invitePagination!.currentPage}
              totalPages={invitePagination!.totalPages}
              handlePageChange={(currentPage) => handlePageChange(currentPage, "users")}
            />
          </div>
        </Modal>
        {alert!.modalOpen && (
          <Alert type={alert!.type} message={alert!.message} onClose={() => setAlert(initialValueAlert)} />
        )}
        <div className="flex flex-col justify-between w-full px-20">
          <div className="flex-1 rounded-lg p-10">
            <ListUser users={contacts} userIdLogged={user.id} text="contato" handleInvite={handleInvite} />
          </div>
          {contacts.length > 0 && (
            <div>
              <Pagination
                currentPage={contactPagination!.currentPage}
                totalPages={contactPagination!.totalPages}
                handlePageChange={(currentPage) => handlePageChange(currentPage, "contacts")}
              />
            </div>
          )}
        </div>
      </div>
    </>
  )
}

export default Contacts
