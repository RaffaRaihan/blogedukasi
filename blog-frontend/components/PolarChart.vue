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
  'rgba(159, 168, 218, 0.5)', // Monday
  'rgba(165, 214, 217, 0.5)', // Tuesday
  'rgba(144, 202, 249, 0.5)', // Wednesday
  'rgba(206, 147, 216, 0.5)', // Thursday
  'rgba(244, 143, 177, 0.5)', // Friday
  'rgba(179, 157, 219, 0.5)', // Saturday
  'rgba(227, 242, 253, 0.5)'  // Sunday
];

onMounted(async () => {
  if (!process.client || !chartRef.value) return;

  try {
    const token = getTokenFromCookies();
    const { data } = await axios.get('http://localhost:8080/admin/stats-top-author', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const labels = data.labels;
    const values = data.data;

    if (chartInstance) chartInstance.destroy();

    chartInstance = new Chart(chartRef.value, {
      type: 'polarArea',
      data: {
        labels,
        datasets: [{
          label: 'Traffic',
          data: values,
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
          },
          tooltip: {
            callbacks: {
              label: (context) => {
                const label = context.label || '';
                const value = context.formattedValue || '';
                return `${label}: ${value}`;
              }
            }
          }
        },
        scales: {
          r: {
            ticks: {
              color: '#444'
            },
            grid: {
              color: '#ccc'
            }
          }
        }
      }
    });

  } catch (error) {
    console.error('Gagal mengambil data dari API:', error);
  }
});
</script>

<template>
  <div class="container">
    <canvas ref="chartRef"></canvas>
  </div>
</template>
