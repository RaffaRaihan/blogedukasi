<template>
  <LoginNavbar v-if="isLoggedIn"/>
  <Navbar v-else/>
  <div class="container mt-4">
    <h1 class="mb-4">Pesan yang Diterima</h1>
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
        <!-- Loop untuk menampilkan kategori -->
        <tr v-for="(messages, index) in messages" :key="messages.ID">
          <th scope="row">{{ index + 1 }}</th>
          <td>{{ messages.content }}</td>
          <td>{{ messages.reply }}</td>
          <td>
            <i 
              v-if="messages.reply" 
              class="bi bi-check-circle-fill text-success"
              title="Sudah dibalas">
            </i>
            <p 
              v-else 
              class="text-muted">
              Belum Dibalas
            </p>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  <Footer />
</template>

<script setup>
import useAuth from '~/composables/api/token/useAuth';
import useMessages from '~/composables/api/token/useMessages';

definePageMeta({
  middleware: 'auth',
});

const { getTokenFromCookies } = useAuth();

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Return true jika token ada, false jika tidak ada
});

const { messages, fetchMessage } = useMessages();
fetchMessage();

</script>

