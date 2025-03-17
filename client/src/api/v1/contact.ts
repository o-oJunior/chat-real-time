import { InviteStatus } from "@/components/list/listUser";

const MESSAGE_ERROR = { statusCode: 500, error: "Erro na conexão com o servidor, tente novamente mais tarde!" };

export interface Contact {
  userIdInviter?: string;
  inviteStatus?: InviteStatus;
}

export default class ContactAPIService {
  private readonly BASE_URL: string;

  constructor() {
    if (!process.env.NEXT_PUBLIC_URL_API_V1) {
      throw new Error("A variável de ambiente NEXT_PUBLIC_URL_API_V1 não está definida.");
    }
    this.BASE_URL = process.env.NEXT_PUBLIC_URL_API_V1;
  }

  async updateStatusContact(contact: Contact) {
    try {
      const response = await fetch(`${this.BASE_URL}/contact/update/${contact.inviteStatus}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(contact),
      });

      if (!response.ok) {
        throw new Error(`Erro ao atualizar status: ${response.status} - ${response.statusText}`);
      }

      return await response.json();
    } catch (error) {
      return MESSAGE_ERROR;
    }
  }
}
