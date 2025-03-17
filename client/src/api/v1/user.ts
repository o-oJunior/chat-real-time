type UserAuth = {
  username: string;
  password: string;
};

type CreateUser = {
  username: string;
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  confirmPassword: string;
};

const MESSAGE_ERROR = { statusCode: 500, error: "Erro na conexão com o servidor, tente novamente mais tarde!" };

export default class UserAPIService {
  private readonly BASE_URL: string;

  constructor() {
    if (!process.env.NEXT_PUBLIC_URL_API_V1) {
      throw new Error("A variável de ambiente NEXT_PUBLIC_URL_API_V1 não está definida.");
    }
    this.BASE_URL = process.env.NEXT_PUBLIC_URL_API_V1;
  }

  private async fetchAPI(endpoint: string, options: RequestInit = {}) {
    try {
      const response = await fetch(`${this.BASE_URL}${endpoint}`, {
        credentials: "include",
        ...options,
      });
      return await response.json();
    } catch (error) {
      return MESSAGE_ERROR;
    }
  }

  validateAuthentication() {
    return this.fetchAPI("/user/validate/authentication");
  }

  userAuthentication(user: UserAuth) {
    return this.fetchAPI("/user/authentication", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    });
  }

  createUser(user: CreateUser) {
    return this.fetchAPI("/user/create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    });
  }

  logout() {
    return this.fetchAPI("/user/logout");
  }

  getUsers(page: number = 1, limit: number = 10, username: string = "") {
    const queryParams = new URLSearchParams({ page: page.toString(), limit: limit.toString() });
    if (username) queryParams.append("username", username);

    return this.fetchAPI(`/user/search?${queryParams.toString()}`);
  }

  getContacts(page: number = 1, limit: number = 10, group: string = "", username: string = "") {
    const queryParams = new URLSearchParams({ page: page.toString(), limit: limit.toString() });
    if (group) queryParams.append("group", group);
    if (username) queryParams.append("username", username);

    return this.fetchAPI(`/user/contacts?${queryParams.toString()}`);
  }
}
