export const apiFetch = (path: string, init: RequestInit|undefined = undefined) => {
    if(!path.startsWith('/')) path = '/' + path
    if(!init) init = { credentials: 'include' }
    else init.credentials = 'include'
    return fetch(`${useRuntimeConfig().public.api}${path}`, init)
}