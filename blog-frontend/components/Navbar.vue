<template>
  <!-- Navbar -->
  <nav class="navbar navbar-expand-lg">
    <div class="container">
      <NuxtLink class="navbar-brand" to="/">Raffa Mr.</NuxtLink>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav ms-auto">
          <li class="nav-item"><NuxtLink class="nav-link text-white" to="/">Beranda</NuxtLink></li>
          <li class="nav-item"><NuxtLink class="nav-link text-white" to="/about">About</NuxtLink></li>
          <li class="nav-item"><NuxtLink class="nav-link text-white" to="/contact">Contact</NuxtLink></li>
          <!-- Logika untuk menampilkan input atau tombol login -->
          <li class="nav-item" v-if="isLoggedIn">
            <input class="form-control" type="text" placeholder="search..." aria-label="default input example">
          </li>
          <li class="nav-item" v-else>
            <NuxtLink class="btn btn-outline-light" to="/login">Login</NuxtLink>
          </li>
        </ul>
      </div> 
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue';

const isLoggedIn = ref(false);
</script>

<style scoped>
body {
  font-family: Arial, sans-serif;
}
.card-title {
  font-size: 1.25rem;
  font-weight: bold;
}
.card-text {
  font-size: 0.9rem;
}
.navbar {
  background-color: #00A885;
}
.navbar-brand {
  font-weight: bold;
  color: white;
}
</style>