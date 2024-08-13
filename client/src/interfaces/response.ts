import { IUser } from "@/redux/user/slice"

export interface IResponse {
  data: IUser
  message: string
  statusCode: number
}
