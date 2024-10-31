import { useRouter } from "next/router"
import { useEffect } from "react"
import { addUserData, useUser } from "../redux/user/slice"
import { useDispatch } from "react-redux"
import { useAppSelector } from "@/redux/hook"
import Sidebar from "@/components/sidebar/sidebar"
import API_V1_USER from "@/api/v1/user"

const Authentication = ({ children }: { children: React.ReactNode }) => {
  const { user } = useAppSelector(useUser)
  const dispatch = useDispatch()
  const router = useRouter()

  useEffect(() => {
    const nowMS = new Date().getTime()
    const expiresAt = user.expiresAt * 1000
    if (user.username == "" || nowMS >= expiresAt) {
      validateAuthentication()
    }
    validatePageNotFound()
  }, [user, router.pathname])

  const validateAuthentication = async () => {
    const v1 = new API_V1_USER()
    const result = await v1.validateAuthentication()
    if (!result.data || result.statusCode !== 200) {
      return router.push("/login")
    }
    dispatch(addUserData(result.data))
  }

  const validatePageNotFound = () => {
    if (router.pathname == "/404") {
      return router.push("/")
    }
  }
  return (
    <>
      {user.username !== "" ? (
        <div className="flex row w-full">
          <div className="mr-4">{router.pathname !== "/404" && <Sidebar />}</div>
          <div className="py-5 w-full pr-5">{children}</div>
        </div>
      ) : (
        <div>{children}</div>
      )}
    </>
  )
}

export default Authentication
