// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  css: [
      '~/assets/css/main.css',
      'vue-toastification/dist/index.css'
  ],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  runtimeConfig: {
    public: {
      api: process.env.API_URL,
    },
  },

  modules: ['@pinia/nuxt'],
})