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

const MESSAGE_ERROR = {statusCode: 500, error: "Erro na conexão com o servidor, tente novamente mais tarde!"}

export default class API_V1_USER {
  private BASE_URL_API_V1: string

  constructor() {
    this.BASE_URL_API_V1 = process.env.NEXT_PUBLIC_URL_API_V1!
  }

  async validateAuthentication() {
    try {
      const result = await fetch(`${this.BASE_URL_API_V1}/user/validate/authentication`, {
        credentials: "include",
      })
      return result.json()
    } catch (error) {
      return MESSAGE_ERROR
    }
  }

  async userAuthentication(user: UserAuth) {
    try {
      const result = await fetch(`${this.BASE_URL_API_V1}/user/authentication`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(user),
      })
      return result.json()
    } catch (error) {
      return MESSAGE_ERROR
    }
  }

  async createUser(user: CreateUser) {
    try {
      const result = await fetch(`${this.BASE_URL_API_V1}/user/create`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(user),
      })
      return result.json()
    } catch (error) {
      return MESSAGE_ERROR
    }
  }

  async logout(){
    try {
      const result = await fetch(`${this.BASE_URL_API_V1}/user/logout`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include"
      })
      return result.json()
    } catch (error) {
      return MESSAGE_ERROR
    }
  }

  async getUsers(page: number = 1, limit: number = 10, username: string = ""){
      try {
        const includeUsername = username !== "" ? `&username=${username}` : ""
        const result = await fetch(`${this.BASE_URL_API_V1}/user/search?page=${page}&limit=${limit}${includeUsername}`, 
          {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        })
        return result.json()
      } catch (error) {
        return MESSAGE_ERROR
      }
  }
}
