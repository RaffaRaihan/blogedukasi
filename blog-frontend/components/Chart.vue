<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { Chart, registerables } from 'chart.js';
import useAuth from '@/composables/api/token/useAuth';

const { getTokenFromCookies } = useAuth();

const chartRef = ref(null);
let chartInstance = null;

if (process.client) {
  Chart.register(...registerables);
}

onMounted(async () => {
  if (!process.client || !chartRef.value) return;

  try {
    // Ganti dengan endpoint API kamu
    const token = getTokenFromCookies();
    const response = await axios.get('http://localhost:8080/admin/stats-weekly', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const labels = response.data.labels;
    const data = response.data.data;

    chartInstance = new Chart(chartRef.value, {
      type: 'bar',
      data: {
        labels,
        datasets: [{
          label: 'Traffic',
          data,
          backgroundColor: 'blue',
          borderColor: 'blue',
          borderWidth: 1
        }]
      },
      options: {
        responsive: true,
        plugins: {
          legend: {
            display: true,
            labels: { color: 'black' }
          }
        },
        scales: {
          y: { beginAtZero: true }
        }
      }
    });
  } catch (err) {
    console.error('Gagal mengambil data:', err);
  }
});
</script>

<template>
  <div class="container">
    <canvas ref="chartRef"></canvas>
  </div>
</template>
