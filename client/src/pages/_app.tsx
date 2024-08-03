import "@/styles/globals.css"
import type { AppProps } from "next/app"
import Layout from "./layout"
import Head from "next/head"
import Providers from "../redux/providers"
import Authentication from "./authentication"

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Providers>
        <Authentication>
          <Head>
            <link rel="icon" href="/gopher.png" />
          </Head>
          <Layout>
            <Component {...pageProps} />
          </Layout>
        </Authentication>
      </Providers>
    </>
  )
}
