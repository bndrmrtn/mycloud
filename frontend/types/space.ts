export interface Space {
    id: string
    name: string
    size: number
}

export interface File {
    id: string
    directory: string
    file_name: string
    fileinfo: {
        info: FileInfo
    }
}

export interface FileInfo {
    id: string
    extension: string
    type: string
    size: number
    hash: string
}