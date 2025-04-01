import "@/styles/globals.css"
import type { AppProps } from "next/app"
import Head from "next/head"
import Providers from "../redux/providers"
import Authentication from "./authentication"
import Notification from "./notification"

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Providers>
        <Head>
          <link rel="icon" href="/gopher.png" />
        </Head>
        <Authentication>
          <Notification>
            <Component {...pageProps} />
          </Notification>
        </Authentication>
      </Providers>
    </>
  )
}
