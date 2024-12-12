<template>
  <div>
    <NuxtPage />
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Ubuntu:ital,wght@0,300;0,400;0,500;0,700;1,300;1,400;1,500;1,700&display=swap');
body{
  background-color: #F2F9FF;
  font-family: "Ubuntu";
}
</style>

<script setup>
import { useHead } from '@vueuse/head';

useHead({
  titleTemplate: (title) => (title ? `${title} - My Website` : 'My Website')
});
</script>
