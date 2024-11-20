import {type EventHandler, useEventStore} from "~/stores/event";

export const handleOnMessage = (ws: WebSocket, event: MessageEvent<string>) => {
    const data = JSON.parse(event.data) as { type: string } & Record<string, any>
    const events = useEventStore()

    console.log(`Handling event ${data.type}`)

    const handlers = events.get(data.type)
    for(const [key, fn] of Object.entries(handlers)) {
        console.log(key, fn)
        console.info(`Executing event handler ${key} for event: ${data.type}`)
        fn?.(data)
    }
}
