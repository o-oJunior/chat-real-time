import { InviteStatus } from "@/components/list/listUser";

const MESSAGE_ERROR = { statusCode: 500, error: "Erro na conexão com o servidor, tente novamente mais tarde!" };

export interface Invite {
  userIdInviter?: string;
  inviteStatus?: InviteStatus;
}

export default class InviteAPIService {
  private readonly BASE_URL: string;

  constructor() {
    if (!process.env.NEXT_PUBLIC_URL_API_V1) {
      throw new Error("A variável de ambiente NEXT_PUBLIC_URL_API_V1 não está definida.");
    }
    this.BASE_URL = process.env.NEXT_PUBLIC_URL_API_V1;
  }

  async sendInvite(invite: Invite) {
    try {
      const response = await fetch(`${this.BASE_URL}/invite/send`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(invite),
      });
      return await response.json();
    } catch (error) {
      return MESSAGE_ERROR;
    }
  }

  async updateStatusInvite(invite: Invite) {
    try {
      const response = await fetch(`${this.BASE_URL}/invite/update/${invite.inviteStatus}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(invite),
      });

      if (!response.ok) {
        throw new Error(`Erro ao atualizar status do convite: ${response.status} - ${response.statusText}`);
      }

      return await response.json();
    } catch (error) {
      console.error("Erro em updateStatusInvite:", error);
      return MESSAGE_ERROR;
    }
  }
}
