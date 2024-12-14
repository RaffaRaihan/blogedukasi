<template>
    <LoginNavbar v-if="isLoggedIn"/>
    <Navbar v-else />
    <div class="container mt-5 mb-5">
      <h1 class="mb-4">Contact Us</h1>
      <form>
        <div class="mb-3">
          <label for="name" class="form-label">Name</label>
          <input type="text" class="form-control" id="name" placeholder="Your Name" />
        </div>
        <div class="mb-3">
          <label for="email" class="form-label">Email</label>
          <input type="email" class="form-control" id="email" placeholder="Your Email" />
        </div>
        <div class="mb-3">
          <label for="message" class="form-label">Message</label>
          <textarea class="form-control" id="message" rows="4" placeholder="Your Message"></textarea>
        </div>
        <button type="submit" class="btn">Send</button>
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
.btn{
  color: #1D2B53;
  background-color: #FF004D;
  border-color: #1D2B53;
}
.btn:hover{
  color: #FF004D;
  background-color: #1D2B53;
  border-color: #FF004D;
}
</style>