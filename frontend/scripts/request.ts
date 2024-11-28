export const newRequest = (path: string, init: RequestInit = {}) => {
  if (!path.startsWith('/')) path = '/' + path;
  path = useRuntimeConfig().public.api + path;
  if (!init.credentials) init.credentials = 'include';
  return fetch(path, init);
};
