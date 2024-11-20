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

export const fetchUsers = async (cursor: string, admins: boolean): Promise<Pagination<User> | Error> => {
    try {
        const res = await newRequest(`/admin/${admins ? 'admins' : 'users'}?cursor=${cursor}`)
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

export const deleteUser = async (id: string): Promise<Error | null> => {
    try {
        const res = await newRequest(`/admin/users/${id}`, {
            method: 'DELETE'
        })
        if(res.status == 200) {
            return null
        }

        const err = await res.json()
        if(err?.error) return Error(err.error)
        return Error('An unknown error occurred')
    } catch (e: unknown) {
        console.error(e)
        return Error('An unknown error occurred')
    }
}