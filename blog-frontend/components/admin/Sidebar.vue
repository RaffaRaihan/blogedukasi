<template>
    <!-- Sidebar -->
    <aside class="col-md-3 col-lg-2 d-flex flex-column bg-light vh-100 p-3">
          <h4 class="mb-4">Company</h4>
          <nav class="nav flex-column"> 
            <NuxtLink class="nav-link text-black active text-primary" to="/admin/dashboard">Dashboard</NuxtLink>
            <NuxtLink class="nav-link text-black" to="/admin/users">Users</NuxtLink>
            <NuxtLink class="nav-link text-black" to="/admin/articles">Article</NuxtLink>
            <NuxtLink class="nav-link text-black" to="/admin/category">Categories</NuxtLink>
            <Button @onclick="handleLogout" class="btn btn-primary text-black mt-auto">Sign out</Button>
          </nav>
        </aside>
</template>

<script setup>
import Cookies from 'js-cookie' // Pastikan library js-cookie sudah diinstal

const handleLogout = () => {
  // Hapus token dari cookies
  Cookies.remove('token')

  // Redirect ke halaman login
  window.location.href = '/login'
}
</script>
