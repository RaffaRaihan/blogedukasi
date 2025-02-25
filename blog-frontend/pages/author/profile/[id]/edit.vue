<template>
    <div class="container mt-5">
      <div class="card mx-auto mt-2">
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
                v-model="users.name"
                class="form-control"
                placeholder="Masukkan nama"
              />
            </div>
            <div class="mb-3">
              <label for="email" class="form-label fw-bold">Email:</label>
              <input
                type="email"
                id="email"
                v-model="users.email"
                class="form-control"
                placeholder="Masukkan email"
              />
            </div>
            <div class="mb-3">
              <label for="email" class="form-label fw-bold">Bio:</label>
              <textarea
               class="form-control" 
               id="exampleFormControlTextarea1" 
               v-model="users.bio"
               rows="3"
               placeholder="Masukan bio">
              </textarea>
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
              <button type="submit" class="btn btn-success">Simpan Perubahan</button>
              <button type="button" class="btn btn-danger ms-2" @click="cancelEdit">Batal</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import useUpdateUserProfile from "~/composables/api/token/useUpdateUserProfile";
  import { useRouter } from "vue-router";
  
  // Inisialisasi
  const router = useRouter();
  
  // Fungsi untuk memperbarui data pengguna
  const { users, updateUserProfile, handleFileUpload } = useUpdateUserProfile();
  
  // Fungsi untuk membatalkan pengeditan
  const cancelEdit = () => {
    router.push(`/author/profile/${users.value.ID}`);
  };
  </script>
  