export interface User {
    id: string
    g_id: string
    name: string
    email: string
    image_url: string
    role: 'user' | 'admin'
}