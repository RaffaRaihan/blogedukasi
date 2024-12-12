<template>
    <LoginNavbar v-if="isLoggedIn"/>
      <Navbar v-else />
      <div class="container mt-5">
        <h1 class="mb-4">About Us</h1>
        <p>Selamat datang di website kami! Kami menyediakan artikel-artikel berkualitas yang dirancang untuk membantu Anda mendapatkan informasi yang akurat dan bermanfaat.</p>
        <p>Kami berdedikasi untuk menghadirkan konten terbaik bagi pembaca kami. Jika Anda memiliki pertanyaan atau masukan, jangan ragu untuk menghubungi kami.</p>
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