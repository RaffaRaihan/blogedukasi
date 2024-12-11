<template>
  <div ref="editorContainer" class="quill-editor"></div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
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
    // Inisialisasi editor Quill
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

    // Gunakan props untuk mengatur nilai awal editor
    if (props.modelValue) {
      editor.clipboard.dangerouslyPasteHTML(props.modelValue);
      console.log('Model Value:', props.modelValue);
    }

    // Emit perubahan konten ke parent
    editor.on('text-change', () => {
      const content = editor.root.innerHTML;
      emit('update:modelValue', content);
    });
  }
});
</script>


<style scoped>
.quill-editor {
  height: 300px;
}
</style>
  