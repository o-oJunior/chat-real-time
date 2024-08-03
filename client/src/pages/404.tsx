import Head from "next/head"
import React from "react"

type Props = {}

const NotFound = (props: Props) => {
  return (
    <>
      <Head>
        <title>404 Not Found</title>
      </Head>
      <div>Not Found</div>
    </>
  )
}

export default NotFound
