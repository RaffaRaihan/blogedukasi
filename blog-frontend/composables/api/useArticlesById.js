import { ref } from 'vue';
import axios from 'axios';

export default function useArticlesById() {
  const articles = ref([]);
  const error = ref(null);

  const fetchArticle = async (id) => {
    try {
      const response = await axios.get(`http://localhost:8080/articles/${id}`);
      console.log(response);
      articles.value = response.data.data;
    } catch (err) {
      error.value = err.response?.data?.message || err.message;
    }
  };

  return {
    articles,
    error,
    fetchArticle,
  };
}
