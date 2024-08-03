export default async function API_VALIDATE_AUTH() {
  const BASE_URL_API_V1 = process.env.NEXT_PUBLIC_URL_API_V1
  const result = await fetch(`${BASE_URL_API_V1}/user/validate/authentication`, {
    credentials: "include",
  })
  return result.json()
}
