import type {Space} from "~/types/space";
import {newRequest} from "~/scripts/request";
import {useToast} from "vue-toastification";

export const createSpace = async (spaceName: string): Promise<Space | Error> => {
    try {
        const res = await newRequest('/spaces', {
            method: 'POST',
            body: JSON.stringify({name: spaceName})
        })
        if(res.status != 200) {
            const data = await res.json()
            if(data?.error) return Error(data.error)
            return Error('An unknown error occurred')
        }
        const data = await res.json()
        return data as Space
    } catch (e: unknown) {
        console.error('Failed to create space', e)
        return Error('An unknown error occurred')
    }
}

export const createFile = async (id: string, file: File, dir: string, name: string): Promise<Error | null> => {
    try {
        const data = new FormData()
        data.append('file', file)
        data.append('directory', dir)
        if(name) data.append('filename', name)

        const res = await newRequest(`/spaces/${id}/upload`, {
            method: 'POST',
            body: data,
        })

        if(res.status == 200) return null

        const err = await res.json()
        if(err?.error) return Error(err.error)
        return Error('An unknown error occurred')
    } catch (e: unknown) {
        console.error('Failed to upload file', e)
        return Error('An unknown error occurred')
    }
}