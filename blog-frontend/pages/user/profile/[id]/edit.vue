<template>
  <div class="container mt-5">
    <NuxtLink :to="`/user/profile/${user_id}`" class="text-decoration-none" style="color: #1D2B53;"><i class="bi bi-arrow-left"></i>  Kembali</NuxtLink>
    <div class="card mx-auto" style="max-width: 600px;">
      <div class="card-header" style="background-color: #1D2B53;">
        <h3 class="mb-0" style="color: #F9F6E6;">Edit Profil</h3>
      </div>
      <div class="card-body">
        <form @submit.prevent="updateUserProfile">
          <!-- Form untuk Data -->
          <div class="mb-3">
            <label for="name" class="form-label fw-bold">Nama:</label>
            <input
              type="text"
              id="name"
              v-model="user.name"
              class="form-control"
              placeholder="Masukkan nama"
            />
          </div>
          <div class="mb-3">
            <label for="email" class="form-label fw-bold">Email:</label>
            <input
              type="email"
              id="email"
              v-model="user.email"
              class="form-control"
              placeholder="Masukkan email"
            />
          </div>

          <!-- Form untuk Foto -->
          <div class="mb-3">
            <label for="foto" class="form-label fw-bold">Foto Profil:</label>
            <input
              type="file"
              id="foto"
              class="form-control"
              @change="handleFileUpload"
            />
          </div>

          <!-- Tombol -->
          <div class="text-center">
            <button type="submit" class="btn btn-outline-success"><i class="bi bi-check2-circle"></i>  Simpan Perubahan</button>
            <button type="button" class="btn btn-outline-danger ms-2" @click="cancelEdit">
              <i class="bi bi-x-circle"></i>  Batal
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";
import Cookies from "js-cookie";
import { useRoute, useRouter } from "vue-router";

// Inisialisasi
const route = useRoute();
const router = useRouter();
const user = ref({
  name: '',
  email: '',
});
const fotoFile = ref(null);
const error = ref(null);
const user_id = route.params.id;

// Fungsi untuk mendapatkan token
const getTokenFromCookies = () => {
  const token = Cookies.get("token");
  return token;
};

// Mengambil data pengguna saat halaman dimuat
const fetchUserDetails = async () => {
  try {
    const token = getTokenFromCookies();
    const response = await axios.get(`http://localhost:8080/user/${user_id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    user.value = response.data;
    console.log(response)
  } catch (err) {
    error.value = err.response?.data?.message || err.message;
    console.error("Error fetching user details:", err);
  }
};

// Fungsi untuk mengunggah file foto
const handleFileUpload = (event) => {
  fotoFile.value = event.target.files[0];
};

// Fungsi untuk memperbarui data pengguna
const updateUserProfile = async () => {
  try {
    const token = getTokenFromCookies();

    // 1. Perbarui Data (Nama dan Email)
    await axios.put(
      `http://localhost:8080/user/${user_id}`,
      {
        name: user.value.name,
        email: user.value.email,
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // 2. Perbarui Foto (Jika Ada)
    if (fotoFile.value) {
      const formData = new FormData();
      formData.append("foto", fotoFile.value);

      await axios.put(
        `http://localhost:8080/user/${user_id}/foto`,
        formData,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "multipart/form-data",
          },
        }
      );
    }

    alert("Profil berhasil diperbarui!");
    router.push(`/user/profile/${route.params.id}`);
  } catch (err) {
    error.value = err.response?.data?.message || err.message;
    console.error("Error updating profile:", err);
    alert("Terjadi kesalahan saat memperbarui profil.");
  }
};

// Fungsi untuk membatalkan pengeditan
const cancelEdit = () => {
  router.push(`/user/profile/${route.params.id}`);
};

// Panggil data pengguna saat halaman dimuat
onMounted(fetchUserDetails);
</script>
