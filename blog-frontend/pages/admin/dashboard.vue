<template>
    <div class="container-fluid">
      <div class="row">
        <!-- Sidebar -->
        <aside class="col-md-3 col-lg-2 d-flex flex-column bg-light vh-100 p-3">
          <h4 class="mb-4">Company</h4>
          <nav class="nav flex-column"> 
            <NuxtLink class="nav-link text-black active text-primary" to="/">Dashboard</NuxtLink>
            <NuxtLink class="nav-link text-black" to="/admin/users">Users</NuxtLink>
            <NuxtLink class="nav-link text-black" to="/">Article</NuxtLink>
            <NuxtLink class="nav-link text-black" to="/">Category</NuxtLink>
            <NuxtLink v-model="logout" class="nav-link text-black mt-auto" to="/">Sign out</NuxtLink>
          </nav>
        </aside>
  
        <!-- Main Content -->
        <main class="col-md-9 col-lg-10 p-4">
          <!-- Admin Cards -->
          <div class="mb-4">
            <h2>Workspace admins</h2>
            <div class="d-flex">
              <div class="card me-2" style="width: 150px;">
                <div class="card-body text-center">
                  <h6 class="card-title">John Johnson</h6>
                  <p class="text-muted">CEO</p>
                  <button class="btn btn-sm btn-danger">Remove</button>
                </div>
              </div>
              <div class="card d-flex align-items-center justify-content-center" style="width: 150px; border: dashed 2px #ddd;">
                <button class="btn btn-outline-primary">+ Add Admin</button>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  </template>
  
<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const logout = () => {
  // Hapus token dari localStorage
  localStorage.removeItem('token');

  // Arahkan pengguna ke halaman login (misalnya /login)
  router.push('/login');
};
</script>
  