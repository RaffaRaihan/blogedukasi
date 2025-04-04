<script setup>
import { ref, onMounted } from 'vue';
import { Chart, registerables } from 'chart.js';

if (process.client) {
  Chart.register(...registerables);
}

const chartRef = ref(null);

onMounted(() => {
  if (chartRef.value && process.client) {
    new Chart(chartRef.value, {
      type: 'bar',
      data: {
        labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
        datasets: [{
          label: 'Traffic',
          data: [2000, 2200, 2500, 3500, 2700, 2300, 1000],
          backgroundColor: 'blue',
          borderColor: 'blue',
          borderWidth: 1
        }]
      },
      options: {
        responsive: true,
        plugins: {
          legend: { display: true, labels: { color: 'black' } }
        },
        scales: {
          y: {
            beginAtZero: true
          }
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
