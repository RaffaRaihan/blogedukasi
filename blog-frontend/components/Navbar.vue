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
                        <NuxtLink class="nav-link" to="/">Beranda</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/about">Tentang</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/contact">Kontak</NuxtLink>
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
  background-color: #211951;
  position: sticky;
  top: 0px;
  z-index: 2;
}
.navbar .nav-link {
  color: #F0F3FF;
}
.navbar .nav-link:hover {
  color: #15F5BA;
  transition: 0.3s;
}
.profile-img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.logo{
  color: #F0F3FF;
}
.logo:hover{
  color: #15F5BA;
  transition: 0.3s;
}
.btn{
  color: #211951;
  background-color: #F0F3FF;
  border-color: #211951;
}
.btn:hover{
  color: #15F5BA;
  background-color: #211951;
  border-color: #15F5BA;
  transition: 0.3s;
}
.dropdown-menu{
  background: #211951;
}
.dropdown-item {
  color: #F0F3FF;
}
.dropdown-item:hover{
  color: #27005D;
  background-color: #15F5BA;
  transition: 0.3s;
}
</style>