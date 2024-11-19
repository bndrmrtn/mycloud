import {newRequest} from "~/scripts/request";
import type {Pagination, Service} from "~/types/service";
import type {User} from "~/types/user";

export const fetchServiceInfo = async (): Promise<Service | Error> => {
    try {
        const res = await newRequest('/')
        if(res.status == 200) {
            const data = await res.json()
            return data as Service
        }

        const err = await res.json()
        if(err?.error) return Error(err.error)
        return Error('An unknown error occurred')
    } catch (e: unknown) {
        console.error(e)
        return Error('An unknown error occurred')
    }
}

export const fetchUsers = async (): Promise<Pagination<User> | Error> => {
    try {
        const res = await newRequest('/admin/users')
        if(res.status == 200) {
            const data = await res.json()
            return data as Pagination<User>
        }

        const err = await res.json()
        if(err?.error) return Error(err.error)
        return Error('An unknown error occurred')
    } catch (e: unknown) {
        console.error(e)
        return Error('An unknown error occurred')
    }
}