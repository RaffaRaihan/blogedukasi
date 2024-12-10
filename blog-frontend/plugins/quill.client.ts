// plugins/quill.client.ts
import { defineNuxtPlugin } from '#app';
import Quill from 'quill';
import 'quill/dist/quill.snow.css'; // Atau gunakan 'quill.bubble.css' jika ingin tema bubble.

export default defineNuxtPlugin(() => {
  return {
    provide: {
      quill: Quill,
    },
  };
});
