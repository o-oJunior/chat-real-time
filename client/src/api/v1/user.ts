type UserAuth = {
  username: string
  password: string
}

type CreateUser = {
  username: string
  firstName: string
  lastName: string
  email: string
  password: string
  confirmPassword: string
}

export default class API_V1_USER {
  private BASE_URL_API_V1: string

  constructor() {
    this.BASE_URL_API_V1 = process.env.NEXT_PUBLIC_URL_API_V1!
  }

  async validateAuthentication() {
    const result = await fetch(`${this.BASE_URL_API_V1}/user/validate/authentication`, {
      credentials: "include",
    })
    return result.json()
  }

  async userAuthentication(user: UserAuth) {
    const result = await fetch(`${this.BASE_URL_API_V1}/user/authentication`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(user),
    })
    return result.json()
  }

  async createUser(user: CreateUser) {
    const result = await fetch(`${this.BASE_URL_API_V1}/user/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(user),
    })
    return result.json()
  }

  async logout(){
    const result = await fetch(`${this.BASE_URL_API_V1}/user/logout`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include"
    })
    return result.json()
  }
}
