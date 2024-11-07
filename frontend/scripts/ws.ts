import {useRoute} from "#app";
import {useTimeAgo} from "@vueuse/core";
import {useToast} from "vue-toastification";

export const handleOnMessage = (ws: WebSocket, event: MessageEvent<string>) => {
    const data = JSON.parse(event.data) as { type: string }
    const fn = useToast().info

    switch (data.type) {
        case 'space_file_upload_succeed':
            handleSpaceUploadNotification(fn, data as { file_space_id: string })
            break
        case 'space_file_delete_succeed':
            handleSpaceDeleteNotification(fn, data as { file_space_id: string })
            break
        case 'download_request_finished':
            handleDownloadRequestFinishedNotification(fn, data as { download_id: string, request_id: string, download_expiry: string })
            break
    }
}

const handleSpaceUploadNotification = (fn: (s: string) => void, {file_space_id}) => {
    const route = useRoute()
    if(route.path != `/spaces/${file_space_id}`) return
    fn('New file just uploaded.\nReload to see the changes.')
}

const handleSpaceDeleteNotification = (fn: (s: string) => void, {file_space_id}) => {
    const route = useRoute()
    if(route.path != `/spaces/${file_space_id}`) return
    fn('A file just deleted.\nReload to see the changes.')
}

const handleDownloadRequestFinishedNotification = (fn: (s: string) => void, { download_expiry }) => {
    const ta = useTimeAgo(new Date(download_expiry))
    fn(`Your file download request is ready.\nExpires ${ta.value}.`)
}
