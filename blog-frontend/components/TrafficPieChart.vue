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

const colors = [
  '#9fa8da', // Monday
  '#a5d6d9', // Tuesday
  '#90caf9', // Wednesday
  '#ce93d8', // Thursday
  '#f48fb1', // Friday
  '#b39ddb', // Saturday
  '#e3f2fd'  // Sunday
];

onMounted(async () => {
  if (!process.client || !chartRef.value) return;

  try {
    // Ganti URL ini dengan API backend kamu
    const token = getTokenFromCookies();
    const response = await axios.get('http://localhost:8080/admin/stats-views-category', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const labels = response.data.labels;
    const data = response.data.data;

    if (chartInstance) chartInstance.destroy();

    chartInstance = new Chart(chartRef.value, {
      type: 'pie',
      data: {
        labels,
        datasets: [{
          data,
          backgroundColor: colors,
          borderColor: '#ffffff',
          borderWidth: 2
        }]
      },
      options: {
        responsive: true,
        plugins: {
          legend: {
            position: 'top',
            labels: {
              color: '#444'
            }
          }
        }
      }
    });
  } catch (error) {
    console.error('Gagal ambil data dari API:', error);
  }
});
</script>

<template>
  <div class="container">
    <canvas ref="chartRef"></canvas>
  </div>
</template>
