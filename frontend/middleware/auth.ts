import { newRequest } from "~/scripts/request"
import type { User } from "~/types/user"

export default defineNuxtRouteMiddleware(async to => {
     if(import.meta.server) return

     let ok = false

     try {
          const res = await newRequest('/me')
          if(res.status === 200) {
               const data = await res.json()
               useAuthStore().user = data as User
               ok = true
          }
     } catch (err) {
          console.error('Failed to authenticate', err)
     }

     if(!ok) return navigateTo('/')
})