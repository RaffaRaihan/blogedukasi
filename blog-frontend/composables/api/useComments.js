import { ref } from 'vue';
import axios from 'axios';

export default function useComments() {
  const comments = ref([]);

  const fetchComments = async (articleId) => {
    try {
      const response = await axios.get(`http://localhost:8080/articles/${articleId}/comments`);
      comments.value = response.data;
    } catch (err) {
      console.error('Error fetching comments:', err);
    }
  };

  return {
    comments,
    fetchComments,
  };
}