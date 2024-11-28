import { defineStore } from 'pinia';

export type EventHandler = (data: Record<string, any>) => void;

export const useEventStore = defineStore('event', () => {
  const handlers = ref<Record<string, Record<string, EventHandler>>>({});

  const register = (event: string, name: string, fn: EventHandler) => {
    if (!handlers.value[event]) handlers.value[event] = {};
    handlers.value[event][name] = fn;
  };

  const detach = (event: string, name: string) => {
    delete handlers.value[event][name];
  };

  const get = (event: string): Record<string, EventHandler> => {
    return handlers.value?.[event];
  };

  return { handlers, register, detach, get };
});
