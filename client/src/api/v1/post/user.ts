export default async function ApiPostUser(user: any) {
  const BASE_URL_API_V1 = process.env.NEXT_PUBLIC_URL_API_V1
  const result = await fetch(`${BASE_URL_API_V1}/user/authentication`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(user),
  })
  return result.json()
}
