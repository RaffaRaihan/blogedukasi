// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
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
        { src: "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js",
          integrity: "sha384-kenU1KFdBIe4zVF0s0G1M5b4hcpxyD9F7jL+ojB+vyF0/tzoPpFltm9LXgo+4p7N",
          crossorigin: "anonymous" 
        },
        { src: 'https://cdn.jsdelivr.net/npm/quill@2.0.3/dist/quill.js',
          integrity: "sha384-kenU1KFdBIe4zVF0s0G1M5b4hcpxyD9F7jL+ojB+vyF0/tzoPpFltm9LXgo+4p7N",
          crossorigin: "anonymous"
        }
      ]
    }
  }
})