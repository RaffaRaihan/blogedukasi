<template>
    <!-- Sidebar -->
    <div class="col-lg-4">
      <div class="mb-4">
        <h5>Paling Populer</h5>
        <ul class="list-group">
          <li class="list-group-item">Ini Adalah Judul Artikel <br><span class="text-muted">22 Agustus 2022</span></li>
          <li class="list-group-item">Ini Adalah Judul Artikel <br><span class="text-muted">22 Agustus 2022</span></li>
          <li class="list-group-item">Ini Adalah Judul Artikel <br><span class="text-muted">22 Agustus 2022</span></li>
        </ul>
      </div>
      <div>
        <h5>Kategori</h5>
        <ul v-for="category in category" :key="category.id" class="list-grup">
          <li class="list-grup-item"><NuxtLink class="text-muted" to="/">{{ category.name }}</NuxtLink></li>
        </ul>
      </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const category = ref([]);

const fetchCategory = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category'); // Ganti dengan URL backend Anda
    category.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Panggil fetchCategory saat komponen dimuat
onMounted(() => {
  fetchCategory();
});
</script>