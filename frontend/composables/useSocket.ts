import {useToast} from "vue-toastification";
import type {Ref} from "vue";

const socket = ref<WebSocket|null>(null)
const loopExists = ref(false)

export type HandlerFunc = (content: any) => void

const handlers = ref({}) as Ref<{[key: string]: HandlerFunc}>

export const useSocket = () => {
    const connect = () => {
        socket.value = new WebSocket(`ws://localhost:3002/ws`)
        socket.value.onerror = error
        socket.value.onopen = open
        socket.value.onmessage = handler
    }

    const error = () => {
        if(process.client) useToast().warning('Real-time connection lost')
        connect()
    }

    const open = () => {
        if(process.client) useToast().info('Real time connection established')
        if(!loopExists.value) {
            loopExists.value = true
            loop()
        }
    }

    const loop = () => {
        setTimeout(() => {
            send('echo', 'echo')
            loop()
        }, 2500)
    }

    const handler = (e: MessageEvent<string>) => {
        const msg = JSON.parse(e.data) as { type: string }

        if(handlers.value?.[msg.type]) handlers.value[msg.type](msg)
    }

    const on = (type: string, handler: HandlerFunc) => {
        handlers.value[type] = handler
    }

    const send = (type: string, data: any) => {
        socket.value?.send(JSON.stringify({type, data}))
    }

    const finish = () => {
        socket.value?.close()
    }

    return {
        connect,
        send,
        on,
        finish
    }
}