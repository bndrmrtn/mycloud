
export interface FileDiff {
    total_file_size: number
    unique_file_size: number
}

export type Containers = Array<{
    container: string
    size: number
}>

export interface Analytics {
    file_difference: FileDiff
    os_file_container: Containers
}