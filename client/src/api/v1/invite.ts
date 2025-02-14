const MESSAGE_ERROR = {statusCode: 500, error: "Erro na conexão com o servidor, tente novamente mais tarde!"}

type InviteStatus = "none" | "pending" | "accepted"

export interface Invite {
  userIdInviter?: string
  inviteStatus: InviteStatus
}
export default class API_V1_INVITE {
    private BASE_URL_API_V1: string
  
    constructor() {
      this.BASE_URL_API_V1 = process.env.NEXT_PUBLIC_URL_API_V1!
    }

    async sendInvite(invite: Invite){
        try {
            const result = await fetch(`${this.BASE_URL_API_V1}/invite/send`, {
                method: "POST",
                headers: {
                  "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify(invite),
              })
              return result.json()
        } catch (error) {
            return MESSAGE_ERROR
        }
    }

    async updateStatusInvite(invite: Invite) {
      try {
        const result = await fetch(`${this.BASE_URL_API_V1}/invite/update/${invite.inviteStatus}`, {
            method: "PUT",
            headers: {
              "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify(invite),
          })
          if(result.status !== 500){
            return result.json()
          }
    } catch (error) {
        return MESSAGE_ERROR
    }
    }
}