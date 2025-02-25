<template>
  <div class="container mt-4">
    <NuxtLink to="/admin/dashboard" class="text-decoration-none" style="color: #211951;"><i class="bi bi-arrow-left"></i>  Kembali</NuxtLink>
    <div class="card mx-auto mt-3">
      <div class="card-header" style="background-color: #211951; color: #F9F6E6;">
        <h3 class="mb-0">Profil Pengguna</h3>
      </div>
      <div class="card-body">
        <div class="d-flex">
            <img
              :src="`http://localhost:8080/uploads/${users.foto}`"
              alt="Foto Profil"
              class="rounded-1"
              style="width: 500px; height: 500px;"
            />
            <div class="text-start ms-5 mt-5">
              <div class="d-flex mb-3">
              <h2 class="fw-bold" style="color: #211951;">Nama:</h2>
              <h2 class="ms-2" style="color: #1D2B53;">{{ users.name }}</h2>
            </div>
            <div class="d-flex mb-3">
              <h2 class="fw-bold" style="color: #211951;">Email:</h2>
              <h2 class="ms-2" style="color: #1D2B53;">{{ users.email }}</h2>
            </div>
            <div class="mb-3 d-flex">
              <h2 class="fw-bold">Bio:</h2>
              <h2 class="ms-3">{{ users?.bio || "N/A"}}</h2>
            </div>
            <button class="btn btn-outline-warning" @click="navigateToEditProfile"><i class="bi bi-pencil"></i>  Edit Profil</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import useUsersById from "~/composables/api/token/useUsersById";

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const { users } = useUsersById();

const router = useRouter();

const navigateToEditProfile = () => {
  router.push(`/admin/profile/${users.value.ID}/edit`);
};
</script>
  