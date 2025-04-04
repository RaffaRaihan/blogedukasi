<template>
  <div class="container-fluid">
  <div class="row">
    <!-- Sidebar -->
    <AuthorSidebar />
    <div class="col-md-9 col-lg-10 p-4">
      <div class="d-flex justify-content-between align-items-center mb-4">
        <h1 class="mb-4">Halaman Author - Pesan</h1>
        <NuxtLink to="/author/message/send-message" class="btn btn-outline-primary">Kirim Pesan Ke Admin</NuxtLink>
      </div>
    <hr>
    <table class="table table-hover align-middle mt-3">
      <thead>
        <tr>
          <th scope="col">No</th>
          <th scope="col">Pesan Anda</th>
          <th scope="col">Pesan Admin</th>
          <th scope="col">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(message, index) in messages" :key="message.ID">
          <th scope="row">{{ index + 1 }}</th>
          <td v-html="message.content"></td>
          <td v-html="message.reply"></td>
          <td>
            <i 
              v-if="message.reply" 
              class="bi bi-check-circle-fill text-success"
              title="Sudah dibalas">
            </i>
            <p v-else class="text-muted">Belum Dibalas</p>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  </div>
  </div>
</template>

<script setup>
import useMessages from '~/composables/api/token/useMessages';

definePageMeta({
  middleware: 'auth',
});

const { messages, fetchMessage } = useMessages();
fetchMessage();
</script>
