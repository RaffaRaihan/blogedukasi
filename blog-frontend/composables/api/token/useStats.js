import { ref } from 'vue';
import axios from 'axios';
import useAuth from '@/composables/api/token/useAuth';

export default function useStatsArticles() {

  const { getTokenFromCookies } = useAuth();
  // State untuk menyimpan data
  const stats = ref([]);
  const error = ref(null);
  
  // Fetch data dari API
  const fetchStats = async () => {
    try {
      const token = getTokenFromCookies();
      const response = await axios.get('http://localhost:8080/api/article-stats', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
      stats.value = response.data;
    } catch (err) {
      console.error('Gagal fetch data statistik:', err);
      error.value = err.message;
    }
  };

  return {
    stats,
    fetchStats,
  };
}