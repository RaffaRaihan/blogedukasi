<script setup>
import { ref, onMounted } from 'vue';
import { Chart, registerables } from 'chart.js';

// Pastikan Chart.js hanya berjalan di klien
if (process.client) {
  Chart.register(...registerables);
}

const chartRef = ref(null);

onMounted(() => {
  if (chartRef.value && process.client) {
    new Chart(chartRef.value, {
      type: 'line',
      data: {
        labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
        datasets: [{
          label: 'Traffic',
          data: [2000, 2200, 2500, 3500, 2700, 2300, 1500],
          borderColor: 'blue',
          backgroundColor: 'rgba(0, 0, 255, 0.1)',
          pointBackgroundColor: 'blue',
          pointBorderColor: 'white',
          pointRadius: 5,
          fill: true
        }]
      },
      options: {
        responsive: true,
        plugins: {
          legend: { display: true, labels: { color: 'black' } }
        },
        scales: {
          y: { beginAtZero: false }
        }
      }
    });
  }
});
</script>

<template>
  <div class="container">
    <canvas ref="chartRef"></canvas>
  </div>
</template>
