import {useToast} from "vue-toastification";
// @ts-ignore
import DownloadRequestFinished from "~/components/toast/download-request-finished.vue";
// @ts-ignore
import DownloadIcon from "~/components/icons/download-icon.vue";

export const handleSocketDownloadRequestFinished = (d: any) => {
    if(!process.client) return

    const toast = useToast()
    const data = d as { type: string; request_id: string; download_id: string; download_expiry: string }

    // Not working
    toast.info({
        component: DownloadRequestFinished,
        listeners: {
            move: () => useRouter().push('/downloads#' + data.download_id)
        },
        props: {
            request_id: data.request_id
        },
        closeOnClick: false,
        timeout: false,
        hideProgressBar: true,
        icon: DownloadIcon,
        position: "top-right",
        pauseOnFocusLoss: true,
        pauseOnHover: false,
        draggable: true,
        draggablePercent: 0.6,
        showCloseButtonOnHover: false,
        closeButton: "button",
        rtl: false
    })
}