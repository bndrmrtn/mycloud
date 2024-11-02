import {defineStore} from "pinia";

export const useLoaderStore = defineStore('loader', () => {
    const loading = ref(true)

    const start = () => {
        loading.value = true
    }

    const finish = () => {
        loading.value = false
    }

    return { loading, start, finish }
})
