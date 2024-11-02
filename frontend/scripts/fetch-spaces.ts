import {newRequest} from "~/scripts/request";
import type {Space, SpaceFile} from "~/types/space";

export const fetchSpaces = async (): Promise<Array<Space> | null> => {
    try {
        const res = await newRequest('/spaces')
        if(res.status != 200) return null
        const data = await res.json()
        return data as Array<Space>
    } catch (e: unknown) {
        console.error(e)
        return null
    }
}

export const fetchSpace = async (spaceID: string): Promise<Space | null> => {
    try {
        const res = await newRequest(`/spaces/${spaceID}`)
        if(res.status != 200) return null
        const data = await res.json()
        return data as Space
    } catch (e: unknown) {
        console.error(e)
        return null
    }
}

export const fetchDirs = async (spaceID: string, path: string): Promise<Array<string> | null> => {
    try {
        const res = await newRequest(`/spaces/${spaceID}/fs?path=${path}`)
        if(res.status != 200) return null
        return await res.json() as Array<string>
    } catch (e: unknown) {
        console.error('Failed to fetch dirs', e)
        return null
    }
}

export const fetchFiles = async (spaceID: string, path: string): Promise<Array<SpaceFile> | null> => {
    try {
        const res = await newRequest(`/spaces/${spaceID}/files?path=${path}`)
        if(res.status != 200) return null
        return await res.json() as Array<SpaceFile>
    } catch (e: unknown) {
        console.error('Failed to fetch files', e)
        return null
    }
}

export const dirRoute = (p: string, d: string): string => {
    const path = p.split('?')[0]
    const searchParams = new URLSearchParams(p.split('?')[1] || "")
    const currentDir = searchParams.get('directory')
    if(!currentDir) searchParams.set('directory', '/' + d)
    else searchParams.set('directory', currentDir + '/' + d)
    return `${path}?${searchParams.toString()}`
}