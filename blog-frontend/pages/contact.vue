<template>
    <LoginNavbar v-if="isLoggedIn"/>
    <Navbar v-else />
    <div class="container mt-5 mb-5">
      <h1 class="mb-4" style="color: #1D2B53;">Kontak</h1>
      <form>
        <div class="mb-3">
          <label for="name" class="form-label">Nama</label>
          <input type="text" class="form-control" id="name" placeholder="Nama Kamu" />
        </div>
        <div class="mb-3">
          <label for="email" class="form-label">Email</label>
          <input type="email" class="form-control" id="email" placeholder="Email" />
        </div>
        <div class="mb-3">
          <label for="message" class="form-label">Pesan</label>
          <textarea class="form-control" id="message" rows="4" placeholder="Pesan Kamu"></textarea>
        </div>
        <button type="submit" class="btn btn-primary">Kirim</button>
      </form>
    </div>
    <Footer />
</template>

<script setup>
import { computed } from 'vue';

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Return true jika token ada, false jika tidak ada
});

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};
</script>

<style scoped>
label{
  color: #1D2B53;;
}
</style>