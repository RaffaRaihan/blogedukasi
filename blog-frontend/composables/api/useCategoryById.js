import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute } from 'vue-router';

export default function useFetchCategory() {
  const category = ref({});
  const error = ref('');
  const route = useRoute();

  const fetchCategory = async (id) => {
    try {
      const response = await axios.get(`http://localhost:8080/category/${id}`);
      console.log('Respons API:', response.data);
      
      if (response.data.data) {
        category.value = {
          ...response.data.data,
          articles: response.data.data.articles?.filter(article => article.status === "Sesuai") || []
        };
      }
    } catch (err) {
      error.value = err.response?.data?.message || err.message;
      console.error('Error fetching category:', err);
    }
  };

  onMounted(() => {
    const categoryId = route.params.id;
    fetchCategory(categoryId);
  });

  return {
    category,
    error,
    fetchCategory,
  };
}
