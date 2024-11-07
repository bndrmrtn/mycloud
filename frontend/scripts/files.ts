import {newRequest} from "~/scripts/request";

export const fetchCodeContent = async (fileID: string): Promise<string | null> => {
    try {
        const res = await newRequest(`/files/${fileID}`)
        if(res.status != 200) return null
        return await res.text()
    } catch (e: unknown) {
        console.error('Failed to fetch files', e)
        return null
    }
}
