import ApiPostUser from "@/api/v1/post/user"
import Head from "next/head"
import { ChangeEvent, FormEvent, useEffect, useState } from "react"
import Input from "../../components/input/input"
import { IUser, useUser } from "../../redux/user/slice"
import { useRouter } from "next/router"
import { useAppSelector } from "../../redux/hook"

type UserAuth = {
  username: string
  password: string
}

type Response = {
  data: IUser
  message: string
  statusCode: number
}

const initialValue: UserAuth = { username: "", password: "" }

const Login = () => {
  const [userAuth, setUser] = useState<UserAuth>(initialValue)
  const [messageError, setMessageError] = useState<string>("")
  const { user } = useAppSelector(useUser)
  const router = useRouter()

  const handleChangeInput = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target
    setUser({ ...userAuth, [name]: value })
  }

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault()
    if (userAuth.username.trim() == "" || userAuth.password == "") {
      return setMessageError("Preencha todos os campos!")
    }
    const result: Response = await ApiPostUser(userAuth)
    if (result.statusCode !== 200) {
      return setMessageError(result.message)
    }
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
      <div className="min-h-screen flex flex-wrap-reverse items-center justify-around bg-gray-100">
        <div className="bg-white p-8 rounded-lg shadow-md max-w-sm w-full">
          <h2 className="text-xl font-bold mb-6 text-center">Entrar</h2>
          <form onSubmit={handleSubmit} method="POST">
            <div className="mb-4">
              <Input
                type="text"
                name="username"
                placeholder="Usuário"
                maxLength={15}
                onChange={handleChangeInput}
                value={userAuth.username}
              />
            </div>
            <div className="mb-4">
              <Input
                type="password"
                name="password"
                placeholder="Senha"
                maxLength={20}
                onChange={handleChangeInput}
                value={userAuth.password}
              />
            </div>
            {messageError && (
              <div className="flex justify-center py-1 my-4">
                <span className="text-red-600 text-sm">{messageError}</span>
              </div>
            )}
            <button
              type="submit"
              className="w-full bg-primary text-white py-2 px-4 rounded-md hover:bg-primary-hover focus:outline-none focus:bg-primary-hover"
            >
              Entrar
            </button>
          </form>
        </div>
      </div>
    </>
  )
}

export default Login
