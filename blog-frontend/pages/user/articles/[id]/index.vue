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
              <NuxtLink to="/user/dashboard" v-if="isLoggedIn">Beranda</NuxtLink>
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

          <!-- Like Button -->
          <div class="d-flex align-items-center mt-3">
            <button @click="likeArticle" class="btn btn-sm" :class="likedByUser ? 'btn-danger' : 'btn-outline-danger'">
              <i class="bi" :class="likedByUser ? 'bi-heart-fill' : 'bi-heart'"></i> Suka
            </button>
            <span class="ms-2">{{ likeCount }} suka</span>
          </div>    
          <!-- Comments Section -->
          <div class="mt-5">
            <h4>Komentar</h4>
            <div class="accordion" id="accordionExample" v-if="comments.length > 0">
              <div
                class="accordion-item"
                v-for="comment in comments.filter(c => !c.parent_id)"
                :key="comment.ID"
                style="border-color:#211951;"
              >
                <h2 class="accordion-header">
                  <button
                    class="accordion-button"
                    type="button"
                    data-bs-toggle="collapse"
                    :data-bs-target="'#collapse' + comment.ID"
                    aria-expanded="false"
                    :aria-controls="'collapse' + comment.ID"
                  >
                    <div class="d-flex mt-2">
                      <img 
                        :src="comment.user?.foto ? `http://localhost:8080/uploads/${comment.user.foto}` : '../../../assets/img/windah senyum Roblox.jpg'" 
                        class="img-fluid mt-2" 
                        alt="Foto Profile" 
                        style="width: 5rem; height: 5rem; border-radius: 50%;"
                      >
                      <div class="d-block ms-3">
                        <div class="d-block">
                          <h5 class="mb-0 fw-bold">
                            {{ comment.user?.name || 'user tidak ada' }}
                            <span
                              v-if="comment.user && comment.user.ID === articles?.author?.ID"
                              class="badge bg-primary ms-2"
                            >Author</span>
                          </h5>
                          <p class="mb-2 text-muted">@{{ comment.user?.email || '-' }}</p>
                        </div>
                        <p class="mb-1">{{ comment.content }}</p>
                        <div class="d-flex">
                          <NuxtLink class="text-muted mt-2" data-bs-toggle="modal" data-bs-target="#exampleModal" @click="setReplyComment(comment.ID)">Balas</NuxtLink>
                        
                          <!-- Hanya pemilik komentar yang bisa edit/hapus -->
                          <template v-if="isUserDecoded && comment.user && Number(comment.user.ID) === Number(userData?.ID)">
                            <a class="text-decoration-none text-dark ms-2 mt-2" href="#" role="button" id="dropdown" data-bs-toggle="dropdown" aria-expanded="false"><i class="bi bi-three-dots-vertical"></i></a>
                            <ul class="dropdown-menu" aria-labelledby="dropdown">
                              <li><NuxtLink class="dropdown-item" data-bs-toggle="modal" data-bs-target="#editModal" @click="setReplyComment(comment.ID)">Edit</NuxtLink></li>
                              <li><NuxtLink class="dropdown-item" data-bs-toggle="modal" data-bs-target="#hapusModal" @click="setReplyComment(comment.ID)">Hapus</NuxtLink></li>
                            </ul>
                          </template>
                        </div>
                      </div>
                    </div>
                  </button>
                </h2>
                <div
                  :id="'collapse' + comment.ID"
                  class="accordion-collapse collapse"
                  data-bs-parent="#accordionExample"
                >
                  <div class="accordion-body">
                    <strong>Balasan:</strong>
                    <div v-if="comments.some(c => c.parent_id === comment.ID)">
                      <div
                        v-for="reply in comments.filter(c => c.parent_id === comment.ID)"
                        :key="reply.ID"
                        class="mt-2 p-2 border rounded border-dark"
                      >
                        <div class="d-flex">
                          <img 
                            :src="reply.user?.foto ? `http://localhost:8080/uploads/${reply.user.foto}` : '../../../assets/img/windah senyum Roblox.jpg'" 
                            class="img-fluid mt-2" 
                            alt="Foto Profile" 
                            style="width: 3rem; height: 3rem; border-radius: 50%;"
                          >
                          <div class="d-block ms-3">
                            <div class="d-block">
                              <h6 class="mb-0 fw-bold">
                                {{ reply.user?.name || 'user tidak ada' }}
                                <span
                                  v-if="reply.user && reply.user.ID === articles?.author?.ID"
                                  class="badge bg-primary ms-2"
                                >Author</span>
                              </h6>
                              <p class="mb-1 text-muted">@{{ reply.user?.email || '-' }}</p>
                            </div>
                            <p class="mb-1">{{ reply.content }}</p>
                            <NuxtLink class="text-muted mt-2" data-bs-toggle="modal" data-bs-target="#exampleModal" @click="setReplyComment(reply.ID)">Balas</NuxtLink>
                          
                            <!-- Hanya pemilik reply yang bisa edit/hapus -->
                            <template v-if="isUserDecoded && reply.user && Number(reply.user.ID) === Number(userData?.ID)">
                              <a class="text-decoration-none text-dark ms-2 mt-2" href="#" role="button" id="dropdown" data-bs-toggle="dropdown" aria-expanded="false"><i class="bi bi-three-dots-vertical"></i></a>
                              <ul class="dropdown-menu" aria-labelledby="dropdown">
                                <li><NuxtLink class="dropdown-item" data-bs-toggle="modal" data-bs-target="#editModal" @click="setReplyComment(reply.ID)">Edit</NuxtLink></li>
                                <li><NuxtLink class="dropdown-item" data-bs-toggle="modal" data-bs-target="#hapusModal" @click="setReplyComment(reply.ID)">Hapus</NuxtLink></li>
                              </ul>
                            </template>
                          </div>
                        </div>
                      </div>
                    </div>
                    <p v-else>Tidak ada balasan.</p>
                  </div>
                </div>
              </div>
            </div>
            <p v-else>Belum ada komentar.</p>

            <!-- Add Comment Form -->
            <div class=" mt-3 mb-3">
              <h5>Tambah Komentar</h5>
              <form @submit.prevent="submitComment(articles.ID)">
                <div>
                  <textarea v-model="newComment" class="form-control" placeholder="Tulis komentar..." style="border-color:#211951;"></textarea>
                  <!-- Pastikan articleId dikirim dengan benar -->
                  <button type="submit" class="btn-2 btn-outline-primary mt-3">Komentar</button>
                </div>
              </form>
            </div>
            <div class="my-3">
              <h5>Laporkan Artikel</h5>
              <NuxtLink :to="`/user/articles/${articles.ID}/report`" class="text-muted">Laporkan Artikel</NuxtLink>
            </div>
          </div>
        </div>
      </div>
      <!-- Sidebar -->
      <Sidebar />
    </div>
  </div>

  <!-- Modal Reply-->
  <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h1 class="modal-title fs-5" id="exampleModalLabel">Balas Komentar</h1>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <textarea v-model="replyComment" class="form-control" placeholder="Tulis komentar..."></textarea>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">Batal</button>
          <button type="submit" class="btn btn-outline-primary" @click="sendReplyComment" data-bs-dismiss="modal">Kirim</button>
        </div>
      </div>
    </div>
  </div>
  <!-- Modal Edit-->
  <div class="modal fade" id="editModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h1 class="modal-title fs-5" id="exampleModalLabel">Edit Komentar</h1>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <textarea v-model="editComment" class="form-control" placeholder="Tulis komentar..."></textarea>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">Batal</button>
          <button type="submit" class="btn btn-outline-primary" @click="sendEditComment" data-bs-dismiss="modal">Simpan Perubahan</button>
        </div>
      </div>
    </div>
  </div>
  <!-- Modal hapus -->
  <div class="modal fade" id="hapusModal" tabindex="-1" aria-labelledby="hapusModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="logoutModalLabel">Konfirmasi Hapus</h5>
        </div>
        <div class="modal-body">
          Apakah Anda yakin ingin menghapus komentar ini?
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Batal</button>
          <button type="button" class="btn btn-danger" @click="sendHapusComment" data-bs-dismiss="modal">Hapus</button>
        </div>
      </div>
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
import { computed, ref, onMounted } from 'vue';
import { format } from 'date-fns';
import id from 'date-fns/locale/id/index.js'
import { useRoute } from 'vue-router';
import axios from 'axios';

const { getTokenFromCookies } = useAuth();
const { decodeToken, userData } = useDecodeToken();

const route = useRoute();
const articleId = route.params.id;

// Decode token on mount
onMounted(() => {
  decodeToken();
  fetchArticle(articleId);
  fetchComments(articleId);
});

const likeCount = ref(0);
const likedByUser = ref(false);

const fetchLikes = async () => {
  try {
    const { data } = await axios.get(`http://localhost:8080/articles/like/${articleId}`);
    likeCount.value = data.total_likes;
    likedByUser.value = data.liked_by_user;
  } catch (error) {
    console.error('Gagal mengambil data like:', error);
  }
};

const likeArticle = async () => {
  try {
    const token = getTokenFromCookies();
    const config = {
      headers: {
        Authorization: `Bearer ${token}`
      }
    };

    if (likedByUser.value) {
      // Sudah like, jadi kita unlike (DELETE)
      await axios.delete(`http://localhost:8080/user/${articleId}/like`, config);
    } else {
      // Belum like, jadi kita like (POST)
      await axios.post(`http://localhost:8080/user/${articleId}/like`, {}, config);
    }

    await fetchLikes(); // update tampilan
  } catch (error) {
    console.error('Gagal melakukan toggle like:', error);
  }
};


onMounted(() => {
  fetchLikes();
});

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token;
});

const isUserDecoded = computed(() => !!userData.value?.ID);

const formatDate = (date) => {
  const formattedDate = new Date(date);
  return isNaN(formattedDate) ? 'Tanggal tidak valid' : format(formattedDate, 'dd MMMM yyyy', { locale: id });
};

const { articles, fetchArticle } = useArticlesById();
const { comments, fetchComments } = useComments();
const {
  newComment,
  replyComment,
  editComment,
  alertMessage,
  alertClass,
  submitComment,
  ReplyComment,
  EditComment,
  DeleteComment
} = useCommentForm();

const formattedContent = computed(() => {
  if (!articles.value?.content) return '';
  return articles.value.content
    .replace(/<img /g, '<img class="img-fluid" ')
    .replace(/style="[^"]*"/g, '');
});

const selectedCommentId = ref(null);

const setReplyComment = (commentId, content = '') => {
  selectedCommentId.value = commentId;
  editComment.value = content;
};

const sendReplyComment = () => {
  if (selectedCommentId.value) ReplyComment(selectedCommentId.value);
};

const sendEditComment = () => {
  if (selectedCommentId.value) EditComment(selectedCommentId.value, editComment.value);
};

const sendHapusComment = () => {
  if (selectedCommentId.value) DeleteComment(selectedCommentId.value);
};
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
.btn-2{
  padding: 0.5rem;
  border: solid 2px;
  border-radius: 8px;
  color: #F0F3FF;
  background-color: #211951;
  border-color: #211951;
}
.btn-2:hover{
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
.badge.bg-primary {
  background-color: #7F27FF !important;
}
</style>
