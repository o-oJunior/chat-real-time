import { useAppSelector } from "@/redux/hook"
import { useUser } from "@/redux/user/slice"
import Head from "next/head"

const Home = () => {
  const { user } = useAppSelector(useUser)
  return (
    <>
      <Head>
        <title>Chat - Home</title>
      </Head>
      <main>
        <h1>Hello {user.username}!!!!</h1>
      </main>
    </>
  )
}

export default Home
