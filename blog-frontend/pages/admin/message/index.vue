<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Report</h1>
        <!-- Tabel Kategori -->
        <table class="table table-hover align-middle mt-3">
          <thead>
            <tr>
              <th scope="col">No</th>
              <th scope="col">Sender_Id</th>
              <th scope="col">Nama</th>
              <th scope="col">Pesan</th>
              <th scope="col">Reply</th>
              <th scope="col">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(msg, index) in message" :key="msg.ID">
              <th scope="row">{{ index + 1 }}</th>
              <td>{{ msg.user_id }}</td>
              <td>{{ msg.name }}</td>
              <td>{{ msg.content }}</td>
              <td>
                {{ msg.reply }}
              </td>
              <td>
                <button
                  v-if="!msg.reply"
                  class="btn btn-outline-primary"
                  data-bs-toggle="modal"
                  data-bs-target="#exampleModal"
                  @click="openReplyModal(msg)"
                >
                  Balas
                </button>
                <i
                  v-else
                  class="bi bi-check-circle-fill text-success"
                  title="Sudah dibalas"
                ></i>
              </td>
            </tr>
          </tbody>
        </table>
        <!-- Modal -->
        <div
          class="modal fade"
          id="exampleModal"
          tabindex="-1"
          aria-labelledby="exampleModalLabel"
          aria-hidden="true"
        >
          <div class="modal-dialog">
            <div class="modal-content">
              <div class="modal-header">
                <h1 class="modal-title fs-5" id="exampleModalLabel">Balas Pesan</h1>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
              </div>
              <div class="modal-body">
                <textarea
                  v-model="replyMessage.reply"
                  class="form-control"
                  rows="4"
                  placeholder="Tulis balasan Anda di sini"
                ></textarea>
              </div>
              <div class="modal-footer">
                <button
                  type="button"
                  class="btn btn-outline-danger"
                  data-bs-dismiss="modal"
                >
                  Batal
                </button>
                <button
                  type="button"
                  class="btn btn-outline-primary"
                  @click="replyMessages"
                  data-bs-dismiss="modal"
                >
                  Kirim
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const message = ref([]); // Data pesan
const replyMessage = ref({ reply: '' }); // Data balasan
const selectedMessageId = ref(null); // ID pesan yang dipilih untuk balasan

// Ambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

// Fetch pesan dari API
const fetchMessage = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const response = await axios.get('http://localhost:8080/admin/messages', {
      headers: { Authorization: `Bearer ${token}` },
    });
    message.value = response.data.data;
  } catch (error) {
    console.error('Error fetching messages:', error);
  }
};

// Buka modal balas
const openReplyModal = (msg) => {
  selectedMessageId.value = msg.ID; // Simpan ID pesan yang dipilih
  replyMessage.value.reply = ''; // Reset input balasan
};

// Kirim balasan
const replyMessages = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    if (!replyMessage.value.reply.trim()) {
      alert('Balasan tidak boleh kosong!');
      return;
    }

    await axios.post(
      `http://localhost:8080/admin/${selectedMessageId.value}/reply-message`,
      replyMessage.value,
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    alert('Balasan berhasil dikirim!');
    await fetchMessage(); // Refresh daftar pesan
  } catch (error) {
    console.error('Error sending reply:', error);
    alert('Gagal mengirim balasan.');
  }
};

// Ambil data pesan saat komponen dimuat
onMounted(() => {
  fetchMessage();
});
</script>
