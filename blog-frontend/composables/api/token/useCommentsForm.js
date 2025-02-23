import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute } from 'vue-router';
import useAuth from '@/composables/api/token/useAuth';
import useComments from '@/composables/api/useComments';
import useArticle from '~/composables/api/useArticlesById';

export default function useCommentForm() {
  const { getTokenFromCookies } = useAuth();
  const { comments, fetchComments } = useComments();
  const { fetchArticle } = useArticle();

  const user_id = ref('');
  const newComment = ref('');
  const alertMessage = ref('');
  const alertClass = ref('');
  const route = useRoute();

  const submitComment = async (articleId) => {
    try {
      const token = getTokenFromCookies();
      if (!token) throw new Error("Token tidak ditemukan. Harap login terlebih dahulu.");

      const response = await axios.post(
        `http://localhost:8080/user/articles/${articleId}/comments`,
        {
          content: newComment.value,
          article_id: articleId,
          user_id: user_id.value,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const newCommentData = response.data;
      comments.value.push(newCommentData);
      newComment.value = "";

      alertMessage.value = `Komentar berhasil ditambahkan!`;
      alertClass.value = 'alert alert-success';
      setTimeout(() => {
        window.location.reload();
      }, 0);
    } catch (err) {
      console.error("Error submitting comment:", err);
      const token = getTokenFromCookies();
      
      if (!token) {
        alertMessage.value = 'Harap login dulu!';
        alertClass.value = 'alert alert-danger';
        setTimeout(() => {
          window.location.href = "/login";
        }, 3000);
      } else {
        alertMessage.value = 'Gagal menambahkan komentar!';
        alertClass.value = 'alert alert-danger';
      }
    }
  };

  onMounted(() => {
    const articleId = route.params.id;
    fetchArticle(articleId);
    fetchComments(articleId);
  });

  return {
    user_id,
    newComment,
    alertMessage,
    alertClass,
    submitComment,
  };
}
