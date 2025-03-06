import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute } from 'vue-router';
import useAuth from '@/composables/api/token/useAuth';
import useComments from '@/composables/api/useComments';
import useArticle from '~/composables/api/useArticlesById';
import { jwtDecode } from 'jwt-decode';

export default function useCommentForm() {
  const { getTokenFromCookies } = useAuth();
  const { comments, fetchComments } = useComments();
  const { fetchArticle } = useArticle();

  const user_id = ref('');
  const newComment = ref('');
  const replyComment = ref('');
  const alertMessage = ref('');
  const alertClass = ref('');
  const route = useRoute();

  const decodeToken = () => {
    try {
      const token = getTokenFromCookies();
      if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
      const decoded = jwtDecode(token);
      user_id.value = decoded.user_id;
    } catch (error) {
      console.error('Error decoding token:', error);
    }
  };

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

  const ReplyComment = async (CommentsId) => {
    try {
      const token = getTokenFromCookies();
      if (!token) throw new Error("Token tidak ditemukan. Harap login terlebih dahulu.");

      const response = await axios.post(
        `http://localhost:8080/user/reply`,
        {
          content: replyComment.value,
          comment_id: CommentsId,
          user_id: user_id.value,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const replyCommentData = response.data;
      comments.value.push(replyCommentData);
      replyComment.value = "";

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
  }

  const editComment = ref('');

  const EditComment = async (CommentsId) => {
    try {
      const token = getTokenFromCookies();
      if (!token) throw new Error("Token tidak ditemukan. Harap login terlebih dahulu.");

      const response = await axios.put(
        `http://localhost:8080/user/comments/${CommentsId}`,
        {
          content: editComment.value, // Menggunakan editComment.value yang sesuai
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      // Perbarui komentar di dalam array tanpa menambahkan data baru
      const updatedComment = response.data;
      const index = comments.value.findIndex(comment => comment.ID === CommentsId);
      if (index !== -1) {
        comments.value[index] = updatedComment;
      }

      alertMessage.value = "Komentar berhasil diperbarui!";
      alertClass.value = "alert alert-success";
      setTimeout(() => {
        window.location.reload();
      }, 0);

    } catch (err) {
      console.error("Error mengedit komentar:", err);

      if (err.response && err.response.status === 401) {
        alertMessage.value = "Harap login terlebih dahulu!";
        alertClass.value = "alert alert-danger";
        setTimeout(() => {
          window.location.href = "/login";
        }, 3000);
      } else {
        alertMessage.value = "Gagal memperbarui komentar!";
        alertClass.value = "alert alert-danger";
      }
    }
  };

  const deleteComment = ref(null);

  const DeleteComment = async (id) => {
    try {
      const token = getTokenFromCookies();
      if (!token) {
        throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
      }
  
      await axios.delete(`http://localhost:8080/user/comments/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      alertMessage.value = "Komentar berhasil dihapus!";
      alertClass.value = "alert alert-success";
      setTimeout(() => {
        window.location.reload();
      }, 0);
      console.log(`Komentar dengan ID ${id} berhasil dihapus.`);
    } catch (error) {
      console.error('Error removing comment:', error);
      alertMessage.value = "Gagal menghapus komentar!";
      alertClass.value = "alert alert-danger";
    }
  };

  onMounted(() => {
    const articleId = route.params.id;
    decodeToken();
    fetchArticle(articleId);
    fetchComments(articleId);
  });

  return {
    user_id,
    newComment,
    replyComment,
    editComment,
    deleteComment,
    alertMessage,
    alertClass,
    submitComment,
    ReplyComment,
    EditComment,
    DeleteComment,
  };
}