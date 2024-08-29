import API_V1_USER from "@/api/v1/user"
import Input from "@/components/input/input"
import { IResponse } from "@/interfaces/response"
import Head from "next/head"
import React, { ChangeEvent, FormEvent, useState } from "react"

type Props = {
  toggleAuthentication: (event: boolean) => void
}

type TValidationRule = (value: string) => string

interface IValidationRule {
  [key: string]: TValidationRule
}

type CreateUser = {
  username: string
  firstName: string
  lastName: string
  email: string
  password: string
  confirmPassword: string
  description: string
}

const initialValue = {
  username: "",
  firstName: "",
  lastName: "",
  email: "",
  password: "",
  confirmPassword: "",
  description: "Disponível",
}

const CreateAccount = ({ toggleAuthentication }: Props) => {
  const [createUser, setCreateUser] = useState<CreateUser>(initialValue)
  const [messageError, setMessageError] = useState<string>("")
  const [messageSuccess, setMessageSuccess] = useState<string>("")

  const handleChangeInput = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target
    const regex: RegExp = /[^A-Za-záàâãéèêíïóôõöúçñÁÀÂÃÉÈÊÍÏÓÔÕÖÚÇÑ '-]/g
    const validationRules: IValidationRule = {
      username: (value: string) => value.trim(),
      firstName: (value: string) => value.replace(regex, "").trim(),
      lastName: (value: string) => value.replace(regex, ""),
    }
    if (validationRules[name]) {
      setCreateUser({ ...createUser, [name]: validationRules[name](value) })
    } else {
      setCreateUser({ ...createUser, [name]: value })
    }
  }

  const isEmailValid = (email: string) => {
    const re =
      /^(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/

    return re.test(email)
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

  const messageField: string = "Campo obrigatório!*"
  const validationRules: IValidationRule = {
    username: (value: string) =>
      value.length >= 3 ? "" : "Nome de usuário deve conter no mínimo 3 caracteres",
    firstName: (value: string) => (value.length >= 3 ? "" : "Nome deve conter no mínimo 3 caracteres"),
    lastName: (value: string) => (value.length >= 3 ? "" : "Sobrenome deve conter no mínimo 3 caracteres"),
    email: (value: string) => (isEmailValid(value) ? "" : "Email Inválido!"),
  }

  const validateInput = (input: HTMLInputElement): void => {
    const { name, value } = input
    let validationMessage = ""
    if (!value.trim()) {
      validationMessage = messageField
    } else if (validationRules[name]) {
      validationMessage = validationRules[name](value)
    }

    if (name === "username") {
      const re = /^[a-zA-Z0-9_\^~`´@]+$/
      const isValid = re.test(value)
      if (!isValid) {
        validationMessage = "Nome de usuário contém caracteres especiais ou espaços em branco"
      }
    }
    const isValid = validationMessage === ""
    updateInputValidationState(input, isValid, validationMessage)
  }

  const checkedPassword = (form: HTMLFormElement): void => {
    const inputPassword: HTMLInputElement = form?.querySelector('input[name="password"]')!
    const inputConfirmPassword: HTMLInputElement = form?.querySelector('input[name="confirmPassword"]')!
    const updateInputState = (isValid: boolean, validationMessage: string) => {
      updateInputValidationState(inputPassword!, isValid, validationMessage)
      updateInputValidationState(inputConfirmPassword!, isValid, validationMessage)
    }
    const fieldBlank = !inputPassword?.value.trim() && !inputConfirmPassword?.value.trim()
    const isValueDifferent = inputPassword?.value !== inputConfirmPassword?.value
    if (fieldBlank) {
      updateInputState(false, messageField)
    } else if (inputPassword!.value.length < 5) {
      updateInputState(false, "A senha deve conter no minimo 5 caracteres.")
    } else if (isValueDifferent) {
      updateInputState(false, "As senhas não coincidem!")
    } else {
      updateInputState(true, "")
    }
  }

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault()
    const form: HTMLFormElement = document.querySelector("form")!
    const inputs: NodeListOf<HTMLInputElement> | undefined = form?.querySelectorAll("input")
    const inputsInvalid: string[] = []
    const storeInputInvalid = ({ name }: HTMLInputElement, isValidity: boolean) => {
      isValidity ? null : inputsInvalid.push(name)
    }
    const validationCall = (input: HTMLInputElement, form: HTMLFormElement) => {
      validateInput(input)
      checkedPassword(form)
      input.reportValidity()
    }
    inputs?.forEach((input: HTMLInputElement): void => {
      input.setCustomValidity("")
      validationCall(input, form)
      const isValidity = input.checkValidity()
      storeInputInvalid(input, isValidity)
      input.addEventListener("input", () => {
        validationCall(input, form)
        storeInputInvalid(input, isValidity)
      })
    })
    if (inputsInvalid.length != 0) {
      return setMessageError("Preencha os campos corretamente!")
    }
    const v1 = new API_V1_USER()
    const result: IResponse = await v1.createUser(createUser)
    if (result.statusCode !== 201) {
      return setMessageError(result.message)
    }
    setMessageError("")
    setMessageSuccess("Conta criada com sucesso!")
    setTimeout(() => {
      setMessageSuccess("")
      toggleAuthentication(true)
    }, 1000)
  }

  return (
    <>
      <Head>
        <title>Chat - Criar conta</title>
      </Head>
      <div className="flex min-h-screen items-center justify-center bg-gray-200">
        <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
          <h2 className="text-xl font-bold text-center text-gray-800">Criar conta</h2>
          <form className="space-y-3 flex flex-col" onSubmit={handleSubmit}>
            <Input
              type="text"
              name="username"
              placeholder="Nome de usuário*"
              maxLength={20}
              onChange={handleChangeInput}
              value={createUser.username}
            />

            <div className="flex flex-wrap -mx-2">
              <div className="w-full md:w-1/2 px-2">
                <Input
                  type="text"
                  name="firstName"
                  placeholder="Nome*"
                  maxLength={20}
                  onChange={handleChangeInput}
                  value={createUser.firstName}
                />
              </div>
              <div className="w-full md:w-1/2 px-2 mt-4 md:mt-0">
                <Input
                  type="text"
                  name="lastName"
                  placeholder="Sobrenome*"
                  maxLength={50}
                  onChange={handleChangeInput}
                  value={createUser.lastName}
                />
              </div>
            </div>

            <Input
              type="email"
              name="email"
              placeholder="E-mail*"
              maxLength={40}
              onChange={handleChangeInput}
              value={createUser.email}
            />

            <div className="flex flex-wrap -mx-2">
              <div className="w-full md:w-1/2 px-2">
                <Input
                  type="password"
                  name="password"
                  placeholder="Senha*"
                  maxLength={20}
                  onChange={handleChangeInput}
                  value={createUser.password}
                />
              </div>
              <div className="w-full md:w-1/2 px-2 mt-4 md:mt-0">
                <Input
                  type="password"
                  name="confirmPassword"
                  placeholder="Confirmar senha*"
                  maxLength={20}
                  onChange={handleChangeInput}
                  value={createUser.confirmPassword}
                />
              </div>
            </div>
            {messageError && (
              <div className="flex justify-center py-1 bg-red-200">
                <span className="text-red-600 text-xs">{messageError}</span>
              </div>
            )}
            {messageSuccess && (
              <div className="flex justify-center py-1 bg-green-200">
                <span className="text-green-600 text-xs">{messageSuccess}</span>
              </div>
            )}
            <button
              type="submit"
              className="w-full px-4 py-2 text-white bg-primary rounded-md hover:bg-primary-hover focus:outline-none focus:ring-2 focus:ring-primary-hover"
            >
              Criar Conta
            </button>
            <div className="flex justify-center">
              <span
                className="text-xs text-purple-900 cursor-pointer underline"
                onClick={() => toggleAuthentication(true)}
              >
                Já tem conta? Entrar.
              </span>
            </div>
          </form>
        </div>
      </div>
    </>
  )
}

export default CreateAccount
