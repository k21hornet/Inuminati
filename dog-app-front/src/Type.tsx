export type User = {
  id: number
  username: string
  email: string
  icon: string
  createdAt: Date
  updatedAt: Date
}

export type Dog = {
  id: number
  createdAt: Date
  updatedAt: Date
  img: string
  caption: string | null
  user_id: number
}