import { ref, onMounted } from 'vue';
import axios from 'axios';

export default function useArticles() {
  const articles = ref([]);
  const loadingArticles = ref(true);
  const errorArticles = ref(null);

  const fetchArticles = async () => {
    try {
      const response = await axios.get('http://localhost:8080/articles');
      articles.value = response.data
        .filter(article => article.status === 'Sesuai')
        .sort((a, b) => new Date(b.UpdatedAt) - new Date(a.UpdatedAt));
    } catch (error) {
      console.error('Error fetching articles:', error);
      errorArticles.value = error;
    } finally {
      loadingArticles.value = false;
    }
  };

  onMounted(fetchArticles);

  return {
    articles,
    loadingArticles,
    errorArticles,
    fetchArticles
  };
}
