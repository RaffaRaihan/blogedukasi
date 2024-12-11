<template>
    <LoginNavbar v-if="isLoggedIn"/>
    <Navbar v-else />
    <div>Content about</div>
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