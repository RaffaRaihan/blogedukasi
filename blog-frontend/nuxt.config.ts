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
      title: 'My Website',
      meta: [
        { name: 'description', content: 'Deskripsi default website' }
      ],
      script: [
        { src: 'https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js' }
      ]
    }
  }
})