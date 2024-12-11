<template>
  <div>
    <NuxtPage />
  </div>
</template>

<style>
body{
  background-color: #F2F9FF;
}
</style>

<script setup>
import { useHead } from '@vueuse/head';

useHead({
  titleTemplate: (title) => (title ? `${title} - My Website` : 'My Website')
});
</script>
