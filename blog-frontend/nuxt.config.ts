// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ['@bootstrap-vue-next/nuxt'],
  css: ['bootstrap/dist/css/bootstrap.min.css'],
  plugins: [
    '~/plugins/bootstrap.client.ts',
    '~/plugins/quill.client.ts',
  ],
  build: {
    transpile: ['@heroicons/vue']
  },
  app: {
    head: {
      titleTemplate: '%s - My Website',
      title: 'My Website',
      meta: [
        { name: 'description', content: 'Deskripsi default website' }
      ]
    }
  }
})