<template>
  <div ref="editorContainer" class="quill-editor"></div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from 'vue';
import { useNuxtApp } from '#app';

// Deklarasikan props
const props = defineProps({
  modelValue: String, // Properti untuk menerima nilai awal
});

// Emit untuk dua arah binding
const emit = defineEmits(['update:modelValue']);

const editorContainer = ref<HTMLDivElement | null>(null);
let editor: any = null; // Deklarasi editor di luar blok onMounted

onMounted(() => {
  const { $quill } = useNuxtApp();

  if (editorContainer.value) {
    editor = new $quill(editorContainer.value, {
      placeholder: 'Masukan Content',
      theme: 'snow',
      modules: {
        toolbar: [
          ['bold', 'italic', 'underline'],
          [{ header: [1, 2, false] }],
          [{ list: 'ordered' }, { list: 'bullet' }],
          ['link', 'image'],
        ],
      },
    });

    if (props.modelValue) {
      editor.clipboard.dangerouslyPasteHTML(props.modelValue);
    }

    // Tambahkan class 'img-fluid' ke semua gambar dalam Quill
    editor.on('text-change', () => {
      const content = editor.root.innerHTML;
      editor.root.querySelectorAll('img').forEach((img: { classList: { add: (arg0: string) => void; }; }) => {
        img.classList.add('img-fluid');
      });
      emit('update:modelValue', content);
    });
  }
});

// Pantau perubahan pada modelValue dan perbarui konten editor jika perlu
watch(
  () => props.modelValue,
  (newValue) => {
    if (editor && newValue !== editor.root.innerHTML) {
      editor.clipboard.dangerouslyPasteHTML(newValue || '');
    }
  }
);
</script>

<style scoped>
.quill-editor {
  height: 300px;
}
.img-fluid {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 0 auto;
}
.ql-editor p:has(img) {
  margin: 0;
  padding: 0;
}
</style>
