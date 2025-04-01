import UserAPIService from "@/api/v1/user"
import WebSocketService from "@/api/v1/websocket"
import Input from "@/components/input/input"
import { IResponse } from "@/interfaces/response"
import { useAppSelector } from "@/redux/hook"
import { addUserData, useUser } from "@/redux/user/slice"
import Head from "next/head"
import { useRouter } from "next/router"
import React, { ChangeEvent, FormEvent, useEffect, useState } from "react"
import { useDispatch } from "react-redux"

type UserAuth = {
  username: string
  password: string
}

type Props = {
  toggleAuthentication: (event: boolean) => void
}

const initialValue: UserAuth = { username: "", password: "" }

const Auth = ({ toggleAuthentication }: Props) => {
  const [userAuth, setUser] = useState<UserAuth>(initialValue)
  const [messageError, setMessageError] = useState<string>("")
  const { user } = useAppSelector(useUser)
  const router = useRouter()
  const dispatch = useDispatch()

  const handleChangeInput = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target
    setUser({ ...userAuth, [name]: value })
  }

  const updateInputValidationState = (
    input: HTMLInputElement,
    isValid: boolean,
    validationMessage: string = "",
  ): void => {
    if (isValid) {
      input.classList.remove("border-red-400", "focus:border-red-600")
      input.classList.add("border-gray-300", "focus:border-gray-400")
      input.setCustomValidity("")
    } else {
      input.classList.remove("border-gray-300", "focus:border-gray-400")
      input.classList.add("border-red-400", "focus:border-red-600")
      input.setCustomValidity(validationMessage)
    }
  }

  const validateInput = (input: HTMLInputElement): void => {
    const messageField: string = "Campo obrigatório!*"
    const { value } = input
    let validationMessage = ""
    if (!value.trim()) {
      validationMessage = messageField
    }
    const isValid = validationMessage === ""
    updateInputValidationState(input, isValid, validationMessage)
  }

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault()
    const form: HTMLFormElement = document.querySelector("form")!
    const inputs: NodeListOf<HTMLInputElement> | undefined = form?.querySelectorAll("input")
    const inputsInvalid: string[] = []
    const storeInputInvalid = ({ name }: HTMLInputElement, isValidity: boolean) => {
      isValidity ? null : inputsInvalid.push(name)
    }
    const validationCall = (input: HTMLInputElement) => {
      validateInput(input)
      input.reportValidity()
    }
    inputs?.forEach((input: HTMLInputElement): void => {
      input.setCustomValidity("")
      validationCall(input)
      const isValidity = input.checkValidity()
      storeInputInvalid(input, isValidity)
      input.addEventListener("input", () => {
        validationCall(input)
        storeInputInvalid(input, isValidity)
      })
    })
    if (inputsInvalid.length != 0) {
      return setMessageError("Preencha os campos corretamente!")
    }
    if (userAuth.username.trim() == "" || userAuth.password == "") {
      return setMessageError("Preencha todos os campos!")
    }
    const userService = new UserAPIService()
    const result: IResponse = await userService.userAuthentication(userAuth)
    if (result.statusCode !== 200) {
      return setMessageError(result.message)
    }
    localStorage.setItem("user", JSON.stringify(result.data))
    dispatch(addUserData(result.data))
    const webSocketService = new WebSocketService()
    webSocketService.connectWebSocket(result.data.id, null)
    setMessageError("")
    router.push("/")
  }

  useEffect(() => {
    if (user.username !== "") {
      router.push("/")
    }
  }, [user])

  return (
    <>
      <Head>
        <title>Chat - Entrar</title>
      </Head>
      <div className="min-h-screen flex flex-wrap-reverse items-center justify-around bg-gray-200">
        <div className="bg-white p-8 rounded-lg shadow-md max-w-sm w-full">
          <h2 className="text-xl font-bold mb-6 text-center">Entrar</h2>
          <form className="space-y-3 flex flex-col" onSubmit={handleSubmit} method="POST">
            <Input
              type="text"
              name="username"
              placeholder="Usuário"
              maxLength={15}
              onChange={handleChangeInput}
              value={userAuth.username}
            />
            <Input
              type="password"
              name="password"
              placeholder="Senha"
              maxLength={20}
              onChange={handleChangeInput}
              value={userAuth.password}
            />
            {messageError && (
              <div className="flex justify-center py-1 bg-red-200">
                <span className="text-red-600 text-xs">{messageError}</span>
              </div>
            )}
            <button
              type="submit"
              className="w-full bg-primary text-white py-2 px-4 rounded-md hover:bg-primary-hover focus:outline-none focus:bg-primary-hover"
            >
              Entrar
            </button>
            <div className="flex justify-center">
              <span
                className="text-xs text-purple-900 cursor-pointer underline"
                onClick={() => toggleAuthentication(false)}
              >
                Não tem conta? Criar conta.
              </span>
            </div>
          </form>
        </div>
      </div>
    </>
  )
}

export default Auth
