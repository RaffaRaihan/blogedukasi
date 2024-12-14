<template>
  <!-- Footer -->
  <footer class="text-center text-lg-start">
    <div class="container p-4">
      <div class="row">
        <div class="col-lg-6 col-md-12 mb-4 mb-md-0">
          <h5 class="text-uppercase">Contact</h5>
          <ul class="list-unstyled mb-0">
            <li class="link">
              <p><i class="bi bi-envelope"></i> raffamraihan580@gmail.com</p>
            </li>
            <li class="link">
              <p><i class="bi bi-telephone"></i> +62 812-2159-4441</p>
            </li>
            <li class="link">
              <p>
                <i class="bi bi-house"></i> Kp.Barupadang Desa.Cipada Kec.Cisarua Kab.Bandung Barat Prov.Jawa Barat
              </p>
            </li>
          </ul>
        </div>
        <div class="col-lg-3 col-md-6 mb-4 mb-md-0">
          <h5 class="text-uppercase">Navigation</h5>
          <ul class="list-unstyled mb-0">
            <li>
              <NuxtLink to="/" class="link text-decoration-none">Home</NuxtLink>
            </li>
            <li>
              <NuxtLink to="/about" class="link text-decoration-none">About Us</NuxtLink>
            </li>
            <li>
              <NuxtLink to="/contact" class="link text-decoration-none">Contact</NuxtLink>
            </li>
            <!-- Hanya tampilkan tombol Login jika pengguna belum login -->
            <li v-if="!isLoggedIn">
              <NuxtLink to="/login" class="link text-decoration-none">Login</NuxtLink>
            </li>
          </ul>
        </div>
        <div class="col-lg-3 col-md-6 mb-4 mb-md-0">
          <h5 class="text-uppercase">My Sosmed</h5>
          <ul class="list-unstyled mb-0">
            <li>
              <NuxtLink to="/" target="_black" class="link text-decoration-none"><i class="bi bi-facebook"></i> Raffa Raihan</NuxtLink>
            </li>
            <li>
              <NuxtLink to="/" target="_black" class="link text-decoration-none"><i class="bi bi-instagram"></i> raffa.mr_</NuxtLink>
            </li>
            <li>
              <NuxtLink to="/" target="_black" class="link text-decoration-none"><i class="bi bi-linkedin"></i> Raffa Raihan</NuxtLink>
            </li>
            <li>
              <NuxtLink to="/" target="_black" class="link text-decoration-none"><i class="bi bi-github"></i> RaffaRaihan</NuxtLink>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <div class="text-center p-3" style="background-color: rgba(0, 0, 0, 0.2);">
      &copy; 2024 raffamr, Inc. All rights reserved
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue';
import 'bootstrap-icons/font/bootstrap-icons.css';

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Mengembalikan true jika token ada, false jika tidak
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
footer {
  width: 100%;
  background-color: #211951;
  color: #FF004D;
  text-align: center;
}

.link {
  color: #FF004D;
}
.link:hover {
  color: #FFFFFF;
}
</style>
