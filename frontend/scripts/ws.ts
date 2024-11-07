import {useToast} from "vue-toastification";
import {useRoute} from "#app";
import {useTimeAgo} from "@vueuse/core";

export const handleOnMessage = (ws: WebSocket, event: MessageEvent<string>) => {
    const data = JSON.parse(event.data) as { type: string }

    switch (data.type) {
        case 'space_file_upload_succeed':
            handleSpaceUploadNotification(data as { file_space_id: string })
            break
        case 'space_file_delete_succeed':
            handleSpaceDeleteNotification(data as { file_space_id: string })
            break
        case 'download_request_finished':
            handleDownloadRequestFinishedNotification(data as { download_id: string, request_id: string, download_expiry: string })
            break
    }
}

const handleSpaceUploadNotification = ({file_space_id}) => {
    const route = useRoute()
    if(route.path != `/spaces/${file_space_id}`) return
    useToast().info('New file just uploaded.\nReload to see the changes.')
}

const handleSpaceDeleteNotification = ({file_space_id}) => {
    const route = useRoute()
    if(route.path != `/spaces/${file_space_id}`) return
    useToast().info('A file just deleted.\nReload to see the changes.')
}

const handleDownloadRequestFinishedNotification = ({ download_expiry }) => {
    const ta = useTimeAgo(new Date(download_expiry))
    useToast().info(`Your file download request is ready.\nExpires ${ta.value}.`)
}
