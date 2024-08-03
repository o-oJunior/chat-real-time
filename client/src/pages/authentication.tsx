import { useRouter } from "next/router"
import { useEffect } from "react"
import { addUserData, useUser } from "../redux/user/slice"
import { useDispatch } from "react-redux"
import API_VALIDATE_AUTH from "@/api/v1/get/user"
import { useAppSelector } from "@/redux/hook"
import Sidebar from "@/components/sidebar/siderbar"

const Authentication = ({ children }: { children: React.ReactNode }) => {
  const { user } = useAppSelector(useUser)
  const dispatch = useDispatch()
  const router = useRouter()

  useEffect(() => {
    validateAuthentication()
    validatePageNotFound()
  }, [router.pathname])

  const validateAuthentication = async () => {
    const result = await API_VALIDATE_AUTH()
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
      {user.username !== "" && router.pathname !== "/404" && <Sidebar />}
      {children}
    </>
  )
}

export default Authentication
