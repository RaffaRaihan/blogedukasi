<script setup lang="ts">
const description = ref('My amazing site.')
</script>

<template>
  <div>
    <Meta name="description" :content="description" />
  </div>
</template>
