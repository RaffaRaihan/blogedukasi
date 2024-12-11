<template>
  <!-- Navbar -->
  <nav class="navbar navbar-expand-lg">
        <div class="container">
            <NuxtLink class="logo navbar-brand" to="/">Raffa Mr.</NuxtLink>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav ms-auto">
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/">Home</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/about">About</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/contact">Contact</NuxtLink>
                    </li>
                </ul>
              <div class="d-flex" v-if="isLoggedIn">
                  <form class="d-flex">
                    <input class="input form-control me-2" type="search" placeholder="Search" aria-label="Search">
                </form>
                <div class="dropdown ms-3 mt-1">
                    <a class="d-flex align-items-center text-decoration-none text-white" href="#" role="button" id="profileDropdown" data-bs-toggle="dropdown" aria-expanded="false">
                        <img src="https://picsum.photos/200" alt="Profile" class="profile-img">
                        <span class="ms-2" style="color: #FF004D;">Magda</span>
                    </a>
                    <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="profileDropdown">
                        <li><NuxtLink class="dropdown-item" to="login">Profile</NuxtLink></li>
                        <li><a class="dropdown-item" href="#">Settings</a></li>
                        <li><hr class="dropdown-divider"></li>
                        <li><a class="dropdown-item" href="#">Logout</a></li>
                    </ul>
                </div>
              </div>
              <li class="nav-link" v-else><NuxtLink class="btn" to="/login">Login</NuxtLink></li>
            </div>
        </div>
    </nav>
</template>

<script setup>
import { ref } from 'vue';

const isLoggedIn = ref(false);
</script>

<style scoped>
.navbar {
  background-color: #1D2B53; /* Warna gelap */
}
.navbar .nav-link {
  color: #FF004D; /* Warna teks */
}
.navbar .nav-link:hover {
  color: #FFFFFF; /* Warna teks saat hover */
}
.profile-img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.logo{
  color: #FF004D;
}
.logo:hover{
  color: #FFFFFF;
}
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
.dropdown-menu{
  background: #1D2B53;
}
.dropdown-item {
  color: #FF004D;
}
.dropdown-item:hover{
  color: #1D2B53;
  background-color: #FF004D;
}
.input{
  color: #FF004D;
  border-color: #FF004D;
}
</style>