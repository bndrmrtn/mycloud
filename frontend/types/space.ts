export interface Space {
    id: string
    name: string
    size: number
}

export interface SpaceFile {
    id: string
    directory: string
    file_name: string
    fileinfo: {
        info: SpaceFileInfo
    }
}

export interface SpaceFileInfo {
    id: string
    extension: string
    type: string
    size: number
    hash: string
}