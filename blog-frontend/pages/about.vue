<template>
    <LoginNavbar v-if="isLoggedIn"/>
      <Navbar v-else />
      <div class="container mt-lg-4 mb-4">
        <div class="row">
          <div class="col-lg-6 col-md-8 col-sm-12">
            <h1>About Us</h1>
            <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Sit odio dolores magnam molestias reprehenderit in tempore harum doloribus dolorem, mollitia similique! Aliquam modi, dicta omnis cumque totam voluptatem autem odio?</p>
          </div>
          <div class="col-lg-6">
            <img src="../assets/img/freelancer_web_developer_vector.jpg" class="img-fluid" alt="img-about">
          </div>
        </div>
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