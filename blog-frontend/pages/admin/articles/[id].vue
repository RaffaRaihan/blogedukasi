<template>
  <LoginNavbar v-if="isLoggedIn"/>
  <Navbar v-else />
  <div class="container mt-4">
    <div class="row">
      <div class="col-lg-8">
        <!-- Breadcrumb -->
        <nav aria-label="breadcrumb" v-if="articles?.category?.name">
          <ol class="breadcrumb">
            <li class="breadcrumb-item">
              <NuxtLink to="/admin/articles" v-if="isLoggedIn">Kembali</NuxtLink>
              <NuxtLink to="/" v-else>Beranda</NuxtLink>
            </li>
            <li class="breadcrumb-item">
              <NuxtLink :to="`/category/${articles.category.ID}`">{{ articles.category.name }}</NuxtLink>
            </li>
            <li class="breadcrumb-item active" aria-current="page">
              {{ articles.title }}
            </li>
          </ol>
        </nav>
        <!-- Article -->
        <div v-if="articles">
          <h2>{{ articles.title }}</h2>
          <p class="text-muted">By {{ articles?.author?.name || "Belum di input" }} - {{ formatDate(articles.CreatedAt) }}</p>        
          <h4>Apa itu {{ articles.title }}?</h4>
          <div v-html="formattedContent"></div><br>     
          <hr>
          <!-- Alert Message -->
          <div v-if="alertMessage" class="alert" :class="alertClass" role="alert">{{ alertMessage }}</div>
          <!-- Comments Section -->
          <div class="mt-5" v-if="articles.status === 'Sesuai'">
            <h4>Komentar</h4>
            <ul class="list-group" v-if="comments.length > 0">
              <li class="list-group-item" v-for="comment in comments" :key="comment.id">
                <div class="d-flex mt-2">
                  <img 
                    :src="comment.user.foto ? `http://localhost:8080/uploads/${comment.user.foto}` : '../../../assets/img/windah senyum Roblox.jpg'" 
                    class="img-fluid" 
                    alt="Foto Profile" 
                    style="width: 5rem; height: 5rem; border-radius: 50%;"
                  >
                  <div class="d-block ms-3">
                    <h5 class="mb-0 fw-bold">{{ comment.user.name || 'user tidak ada' }}</h5>
                    <p class="mb-1">{{ comment.content }}</p>
                    <div class="d-flex">
                      <NuxtLink class="text-muted mt-1">Balas</NuxtLink>
                    </div>
                  </div>
                </div>
              </li>
            </ul>
            <p v-else>Belum ada komentar.</p>

            <!-- Add Comment Form -->
            <div class=" mt-3 mb-3">
              <h5>Tambah Komentar</h5>
              <form @submit.prevent="submitComment(articles.ID)">
                <div>
                  <textarea v-model="newComment" class="form-control" placeholder="Tulis komentar..."></textarea>
                  <!-- Pastikan articleId dikirim dengan benar -->
                  <button type="submit" class="btn btn-outline-primary mt-3">Comment</button>
                </div>
              </form>
            </div>
          </div>
          <form @submit.prevent="updateArticle" class="container mb-4 g-0" v-else>
            <p class="text-muted mt-2">*Catatan untuk author</p>
            <QuillEditor v-model="editArticle.note"/>
            <button class="btn mt-2">Kirim</button>
          </form>
          <!-- Admin Konfirmasi Status Artikel -->
          <div v-if="isLoggedIn && articles.status !== 'Sesuai'" class="mt-3 mb-2">
            <h5>Konfirmasi Status Artikel</h5>
            <p>Status saat ini: <strong>{{ articles.status }}</strong></p>
            <button
              class="btn me-2"
              @click="confirmStatus('Sesuai')"
              :disabled="articles.status === 'Sesuai'"
            >
              Tandai Sesuai
            </button>
            <button
              class="btn btn-outline-danger"
              @click="confirmStatus('Belum Sesuai')"
              :disabled="articles.status === 'Belum Sesuai'"
            >
              Tandai Belum Sesuai
            </button>
          </div>
        </div>
      </div>
      <!-- Sidebar -->
      <Sidebar />
    </div>
  </div>
  <Footer />
</template>

<script setup>
import useAuth from '@/composables/api/token/useAuth';
import useDecodeToken from '@/composables/api/token/useDecodeToken';
import useArticlesById from '~/composables/api/useArticlesById';
import useComments from '@/composables/api/useComments';
import useCommentForm from '~/composables/api/token/useCommentsForm';
import axios from 'axios';
import { computed } from 'vue';
import { format } from 'date-fns';
import { id } from 'date-fns/locale';
import { useRoute } from 'vue-router';

const editArticle = ref({
  note: '',
});

const { getTokenFromCookies } = useAuth();

const { decodeToken } = useDecodeToken();
decodeToken();

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Return true jika token ada, false jika tidak ada
});

const formatDate = (date) => {
  const formattedDate = new Date(date);
  if (isNaN(formattedDate)) {
    return 'Tanggal tidak valid';
  }
  return format(formattedDate, 'dd MMMM yyyy', { locale: id });
};
const route = useRoute();
const articleId = route.params.id;

const { articles, fetchArticle } = useArticlesById();
fetchArticle(articleId);

const { comments, fetchComments } = useComments();
fetchComments(articleId);

const { newComment, alertMessage, alertClass, submitComment } = useCommentForm();

const formattedContent = computed(() => {
  if (!articles.value?.content) return '';
  return articles.value.content
    .replace(/<img /g, '<img class="img-fluid" ')
    .replace(/style="[^"]*"/g, ''); // Hapus inline style
});

const updateArticle = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      alertMessage.value = `Token tidak ditemukan. Pastikan Anda sudah login!!`;
      alertClass.value = 'alert alert-danger';
      setTimeout(() => {
        window.location.href = "/login";
      }, 3000);
      return;
    }

    console.log('Artikel yang akan di-update:', editArticle.value);

    // Pertama, update artikel
    const updatedArticle = {
      note: editArticle.value.note,
    };

    await axios.put(
      `http://localhost:8080/admin/articles/${articleId}`,
      updatedArticle,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    alertMessage.value = `Catatan berhasil di kirim`;
    alertClass.value = 'alert alert-success';
    setTimeout(() => {
      window.location.href = "/admin/articles";
    }, 3000);
  } catch (error) {
    console.error('Error updating article:', error);
    alertMessage.value = `Gagal mengirim catatan`;
    alertClass.value = 'alert alert-danger';
  }
};

const confirmStatus = async (status) => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      alertMessage.value = `Token tidak ditemukan. Pastikan Anda sudah login!!`;
      alertClass.value = 'alert alert-danger';
      return;
    }

    await axios.put(
      `http://localhost:8080/admin/articles/${articleId}`,
      { status },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    alertMessage.value = `Status artikel berhasil diubah menjadi "${status}".`;
    alertClass.value = 'alert alert-success';
    setTimeout(() => {
      window.location.reload();
    }, 1500);

    // Update tampilan status artikel
    fetchArticle(articleId);
  } catch (error) {
    console.error('Gagal mengubah status artikel:', error);
    alertMessage.value = `Gagal mengubah status artikel.`;
    alertClass.value = 'alert alert-danger';
  }
};

watchEffect(() => {
  if (articles.value && articles.value.note) {
    editArticle.value.note = articles.value.note;
  }
});
</script>

<style scoped>
.breadcrumb {
  background-color: transparent;
  padding: 0;
  margin-bottom: 1rem;
}
.breadcrumb-item a {
  text-decoration: none;
  color: #15F5BA;
}
.breadcrumb-item.active {
  color: #211951;
}
.btn{
  color: #F0F3FF;
  background-color: #211951;
  border-color: #211951;
}
.btn:hover{
  color: #211951;
  background-color: #F0F3FF;
  transition: 0.3s;
}
.alert {
  position: absolute;
  z-index: 1;
  width: 855px;
}
.ql-editor p:has(img) {
  display: flex;
  justify-content: center;
  margin: 0;
  padding: 0;
}
.ql-editor p{
  background-color: transparent;
}
</style>
