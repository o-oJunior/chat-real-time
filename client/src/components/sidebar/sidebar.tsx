import React, { useState } from "react"
import { useRouter } from "next/router"
import { useAppSelector } from "@/redux/hook"
import { userLogout, useUser } from "@/redux/user/slice"
import API_V1_USER from "@/api/v1/user"
import { IResponse } from "@/interfaces/response"
import { useDispatch } from "react-redux"

const Sidebar = () => {
  const [isDropdownOpen, setIsDropdownOpen] = useState(false)
  const { user } = useAppSelector(useUser)
  const router = useRouter()
  const dispatch = useDispatch()

  const handleNavigation = (path: string) => {
    router.push(path)
  }

  const handleLogout = async () => {
    const v1 = new API_V1_USER()
    const result: IResponse = await v1.logout()
    if (result.statusCode !== 200) {
      return alert("Erro ao sair da conta")
    }
    localStorage.removeItem("user")
    dispatch(userLogout())
    handleNavigation("/login")
  }

  const toggleDropdown = () => {
    setIsDropdownOpen(!isDropdownOpen)
  }

  return (
    <div className="w-20 h-screen bg-primary text-white flex flex-col justify-between items-center py-4 space-y-6">
      <h2 className="text-2xl font-bold text-white">Chat</h2>

      <div className="flex flex-col h-full w-full px-1 space-y-4">
        <div className="relative group flex items-center justify-center">
          <button
            className={`w-full p-2 flex items-center justify-center hover:bg-primary-hover rounded-lg ${
              router.pathname === "/contatos" ? "bg-primary-hover" : ""
            }`}
            onClick={() => handleNavigation("/contacts")}
            aria-label="Contatos"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              viewBox="0 0 512 512"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <path
                fill="white"
                d="M96 0C60.7 0 32 28.7 32 64l0 384c0 35.3 28.7 64 64 64l288 0c35.3 0 64-28.7 64-64l0-384c0-35.3-28.7-64-64-64L96 0zM208 288l64 0c44.2 0 80 35.8 80 80c0 8.8-7.2 16-16 16l-192 0c-8.8 0-16-7.2-16-16c0-44.2 35.8-80 80-80zm-32-96a64 64 0 1 1 128 0 64 64 0 1 1 -128 0zM512 80c0-8.8-7.2-16-16-16s-16 7.2-16 16l0 64c0 8.8 7.2 16 16 16s16-7.2 16-16l0-64zM496 192c-8.8 0-16 7.2-16 16l0 64c0 8.8 7.2 16 16 16s16-7.2 16-16l0-64c0-8.8-7.2-16-16-16zm16 144c0-8.8-7.2-16-16-16s-16 7.2-16 16l0 64c0 8.8 7.2 16 16 16s16-7.2 16-16l0-64z"
              />
            </svg>
          </button>
          <span className="absolute left-full ml-2 flex items-center hidden group-hover:flex bg-primary-hover text-white text-xs rounded px-2 py-1 whitespace-nowrap">
            Contatos
          </span>
        </div>

        <div className="relative group flex items-center justify-center">
          <button
            className={`w-full p-2 flex items-center justify-center hover:bg-primary-hover rounded-lg ${
              router.pathname === "/conversas" ? "bg-primary-hover" : ""
            }`}
            onClick={() => handleNavigation("/conversas")}
            aria-label="Conversas"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              viewBox="0 0 576 512"
            >
              <path
                fill="white"
                d="M284 224.8a34.1 34.1 0 1 0 34.3 34.1A34.2 34.2 0 0 0 284 224.8zm-110.5 0a34.1 34.1 0 1 0 34.3 34.1A34.2 34.2 0 0 0 173.6 224.8zm220.9 0a34.1 34.1 0 1 0 34.3 34.1A34.2 34.2 0 0 0 394.5 224.8zm153.8-55.3c-15.5-24.2-37.3-45.6-64.7-63.6-52.9-34.8-122.4-54-195.7-54a406 406 0 0 0 -72 6.4 238.5 238.5 0 0 0 -49.5-36.6C99.7-11.7 40.9 .7 11.1 11.4A14.3 14.3 0 0 0 5.6 34.8C26.5 56.5 61.2 99.3 52.7 138.3c-33.1 33.9-51.1 74.8-51.1 117.3 0 43.4 18 84.2 51.1 118.1 8.5 39-26.2 81.8-47.1 103.5a14.3 14.3 0 0 0 5.6 23.3c29.7 10.7 88.5 23.1 155.3-10.2a238.7 238.7 0 0 0 49.5-36.6A406 406 0 0 0 288 460.1c73.3 0 142.8-19.2 195.7-54 27.4-18 49.1-39.4 64.7-63.6 17.3-26.9 26.1-55.9 26.1-86.1C574.4 225.4 565.6 196.4 548.3 169.5zM285 409.9a345.7 345.7 0 0 1 -89.4-11.5l-20.1 19.4a184.4 184.4 0 0 1 -37.1 27.6 145.8 145.8 0 0 1 -52.5 14.9c1-1.8 1.9-3.6 2.8-5.4q30.3-55.7 16.3-100.1c-33-26-52.8-59.2-52.8-95.4 0-83.1 104.3-150.5 232.8-150.5s232.9 67.4 232.9 150.5C517.9 342.5 413.6 409.9 285 409.9z"
              />
            </svg>
          </button>
          <span className="absolute left-full ml-2 flex items-center hidden group-hover:flex bg-primary-hover text-white text-xs rounded px-2 py-1 whitespace-nowrap">
            Conversas
          </span>
        </div>

        <div className="relative group flex items-center justify-center">
          <button
            className={`w-full p-2 flex items-center justify-center hover:bg-primary-hover rounded-lg ${
              router.pathname === "/grupos" ? "bg-primary-hover" : ""
            }`}
            onClick={() => handleNavigation("/grupos")}
            aria-label="Grupos"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              viewBox="0 0 640 512"
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <path
                fill="white"
                d="M144 0a80 80 0 1 1 0 160A80 80 0 1 1 144 0zM512 0a80 80 0 1 1 0 160A80 80 0 1 1 512 0zM0 298.7C0 239.8 47.8 192 106.7 192l42.7 0c15.9 0 31 3.5 44.6 9.7c-1.3 7.2-1.9 14.7-1.9 22.3c0 38.2 16.8 72.5 43.3 96c-.2 0-.4 0-.7 0L21.3 320C9.6 320 0 310.4 0 298.7zM405.3 320c-.2 0-.4 0-.7 0c26.6-23.5 43.3-57.8 43.3-96c0-7.6-.7-15-1.9-22.3c13.6-6.3 28.7-9.7 44.6-9.7l42.7 0C592.2 192 640 239.8 640 298.7c0 11.8-9.6 21.3-21.3 21.3l-213.3 0zM224 224a96 96 0 1 1 192 0 96 96 0 1 1 -192 0zM128 485.3C128 411.7 187.7 352 261.3 352l117.3 0C452.3 352 512 411.7 512 485.3c0 14.7-11.9 26.7-26.7 26.7l-330.7 0c-14.7 0-26.7-11.9-26.7-26.7z"
              />
            </svg>
          </button>
          <span className="absolute left-full ml-2 flex items-center hidden group-hover:flex bg-primary-hover text-white text-xs rounded px-2 py-1 whitespace-nowrap">
            Grupos
          </span>
        </div>
      </div>

      <div className="flex flex-col w-full px-1 space-y-4">
        <div className="relative group flex items-center justify-center">
          <button
            className="w-full p-2 flex items-center justify-center hover:bg-primary-hover rounded-lg"
            aria-label="Usuário Logado"
            onClick={toggleDropdown}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              fill="currentColor"
              viewBox="0 0 448 512"
            >
              <path d="M224 256A128 128 0 1 0 224 0a128 128 0 1 0 0 256zm-45.7 48C79.8 304 0 383.8 0 482.3C0 498.7 13.3 512 29.7 512l388.6 0c16.4 0 29.7-13.3 29.7-29.7C448 383.8 368.2 304 269.7 304l-91.4 0z" />
            </svg>
          </button>
          {isDropdownOpen ? (
            <div className="absolute left-full ml-2 mb-10 transform mt-2 w-48 bg-white rounded-md shadow-lg z-10">
              <div className="py-2">
                <span className="block px-4 py-2 text-gray-700 text-sm border-b border-gray-200">
                  Olá, {user.username}!
                </span>
                <button
                  onClick={handleLogout}
                  className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                  Sair
                </button>
              </div>
            </div>
          ) : (
            <span className="absolute left-full ml-2 flex items-center hidden group-hover:flex bg-primary-hover text-white text-xs rounded px-2 py-1 whitespace-nowrap">
              Olá, {user.username}!
            </span>
          )}
        </div>
      </div>
    </div>
  )
}

export default Sidebar
