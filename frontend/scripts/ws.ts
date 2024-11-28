import { type EventHandler, useEventStore } from '~/stores/event';

export const handleOnMessage = (ws: WebSocket, event: MessageEvent<string>) => {
  const data = JSON.parse(event.data) as {
    event: string | undefined;
    data: Record<string, any>;
  };

  if (data?.event == undefined || !data.event) return;

  const events = useEventStore();

  console.log(`Handling event ${data.event}`);

  const handlers = events.get(data.event);
  for (const [key, fn] of Object.entries(handlers)) {
    console.log(key, fn);
    console.info(`Executing event handler ${key} for event: ${data.event}`);
    fn?.(data.data);
  }
};
