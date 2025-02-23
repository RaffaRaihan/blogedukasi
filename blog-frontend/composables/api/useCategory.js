import { ref } from 'vue';
import axios from 'axios';

export default function useCategory() {
  const category = ref([]);
  const loadingCategory = ref(true);

  // Fungsi untuk mendapatkan data kategori
  const fetchCategory = async () => {
    try {
      const response = await axios.get('http://localhost:8080/category'); // Ganti dengan URL backend Anda
      category.value = response.data;
    } catch (error) {
      console.error('Error fetching category:', error);
    } finally {
      loadingCategory.value = false;
    }
  };

  return {
    category,
    loadingCategory,
    fetchCategory,
  };
}