import { ref, onMounted } from 'vue';
import axios from 'axios';

export default function useArticles() {

  // State untuk artikel populer
  const popularArticles = ref([]);
  const errorPopular = ref(null);

  // State untuk artikel biasa
  const articles = ref([]);
  const errorArticles = ref(null);

  // Fetch semua artikel
  const fetchArticles = async () => {
    try {
      const response = await axios.get('http://localhost:8080/articles');
      articles.value = response.data
        .filter(article => article.status === 'Sesuai')
        .sort((a, b) => new Date(b.UpdatedAt) - new Date(a.UpdatedAt));
    } catch (error) {
      console.error('Error fetching articles:', error);
      errorArticles.value = error;
    }
  };

  const fetchPopularArticles = async () => {
    try {
      const response = await axios.get('http://localhost:8080/popular-articles');
      console.log("Popular Articles Response:", response.data);
      popularArticles.value = response.data || []; // Jika response.data undefined, set array kosong
    } catch (error) {
      console.error("Error fetching popular articles:", error);
      errorArticles.value = error;
    } 
  };

  // Track view artikel
  const trackArticleView = async (articleId) => {
    try {
      await axios.post('http://localhost:8080/track-view', { article_id: articleId });
    } catch (error) {
      console.error('Gagal mencatat view artikel:', error);
    }
  };

  onMounted(() => {
    fetchArticles();
    fetchPopularArticles(); // Ambil data artikel populer saat komponen dipasang
  });

  return {
    articles,
    errorArticles,
    fetchArticles,
    trackArticleView,
    popularArticles,
    errorPopular,
    fetchPopularArticles
  };
}
