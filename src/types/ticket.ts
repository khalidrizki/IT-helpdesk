export interface Ticket {
  id: number
  title: string
  requester: string
  requesterEmail: string
  assignee: string | null
  category: string
  status: string
  priority: string
  description: string
  affectedAsset?: string
}
