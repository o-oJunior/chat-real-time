import "@/styles/globals.css"
import type { AppProps } from "next/app"
import Head from "next/head"
import Providers from "../redux/providers"
import Authentication from "./authentication"

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Providers>
        <Head>
          <link rel="icon" href="/gopher.png" />
        </Head>
        <Authentication>
          <Component {...pageProps} />
        </Authentication>
      </Providers>
    </>
  )
}
