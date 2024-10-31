import API_V1_USER from "@/api/v1/user"
import Card from "@/components/card/card"
import ListItem from "@/components/listItem/listItem"
import Alert from "@/components/modal/alert"
import Modal from "@/components/modal/modal"
import Pagination from "@/components/pagination/pagination"
import Search from "@/components/search/search"
import { IResponse } from "@/interfaces/response"
import Head from "next/head"
import React, { ChangeEvent, FormEvent, useState } from "react"

type TSearch = {
  contacts: string
  users: string
}

const initialValueSearch = {
  contacts: "",
  users: "",
}

type TUsers = {
  username: string
  description: string
}

type TApiPagination = {
  currentPage: number
  totalPages: number
}

const initialValuePagination: TApiPagination = {
  currentPage: 0,
  totalPages: 0,
}

type AlertProps = {
  type: "success" | "error" | "warning"
  message: string
  modalOpen: boolean
}

const initialValueAlert: AlertProps = {
  type: "error",
  message: "",
  modalOpen: false,
}
const Contacts = () => {
  const [contacts, setContacts] = useState<any[]>([
    { name: "Olinda", description: "Disponível" },
    { name: "Cloroquina", description: "Disponível" },
  ])
  const [users, setUsers] = useState<TUsers[]>([])
  const [apiPagination, setApiPagination] = useState<TApiPagination>(initialValuePagination)
  const [alert, setAlert] = useState<AlertProps>(initialValueAlert)
  const [isOpen, setIsOpen] = useState<boolean>(false)
  const [search, setSearch] = useState<TSearch>(initialValueSearch)
  const handleModal = async () => {
    if (!isOpen) {
      const params = { page: 1, limit: 10 }
      await getUsers(params.page, params.limit)
    }
    setIsOpen(!isOpen)
  }

  const handleSearch = (e: FormEvent) => {
    e.preventDefault()
  }

  const handleChangeInput = async (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setSearch({ ...search, [name]: value })
    await getUsers(1, 10, value)
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
                  renderItem={(item: any) => (
                    <Card>
                      <div className="flex flex-row justify-between">
                        <div className="flex flex-col gap-1">
                          <span>{item.username}</span>
                          <span className="text-sm text-gray-400">{item.description}</span>
                        </div>
                      </div>
                    </Card>
                  )}
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
        </div>
        <div className="flex flex-row justify-between w-full px-20">
          <div className="flex-1 border border-gray-300 rounded-lg p-10">
            <ListItem list={contacts} text="contato" renderItem={(item) => <div>{item.name}</div>} />
          </div>
        </div>
      </div>
    </>
  )
}

export default Contacts
