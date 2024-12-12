<template>
  <div class="min-vh-100 d-flex align-items-center justify-content-center bg-light">
    <div class="card shadow border-0" style="width: 100%; max-width: 400px;">
      <div class="card-body">
        <div class="mt-1"><i class="bi bi-arrow-left" style="color: #FF004D;"></i><NuxtLink to="/" class="text-decoration-none" style="color: #FF004D;">  Back</NuxtLink></div>
        <h2 class="text-center fw-bold mb-4" style="color: #1D2B53;">Sign in</h2>
        <form @submit.prevent="handleLogin">
          <!-- Input Email -->
          <div class="mb-3">
            <input
              v-model="email"
              type="email"
              class="form-control"
              placeholder="Enter email"
              required
            />
          </div>
          <!-- Input Password -->
          <div class="mb-3">
            <input
              v-model="password"
              type="password"
              class="form-control"
              placeholder="Enter password"
              required
            />
          </div>
          <!-- Button Login -->
          <button
            type="submit"
            class="btn w-100"
          >
            SIGN IN
          </button>
        </form>
        <!-- Register Link -->
        <p class="text-center mt-3" style="color: #1D2B53;">
          No account? <a href="/register" class="link fw-bold">Sign up</a>
        </p>
      </div>
    </div>
  </div>
</template>
  
<script setup>
import { ref } from 'vue'
import axios from 'axios'
import Cookies from 'js-cookie' // Pastikan library ini sudah diinstal dengan `npm install js-cookie`

const email = ref('')
const password = ref('')

const handleLogin = async () => {
  try {
    const response = await axios.post('http://localhost:8080/login', {
      email: email.value,
      password: password.value,
    })

    // Simpan token ke dalam cookies
    Cookies.set('token', response.data.token, { expires: 7 }) // Token disimpan selama 7 hari (opsional)

    alert(`Login Berhasil! Role: ${response.data.role}`)

    // Redirect ke halaman sesuai role
    if (response.data.role === 'admin') {
      window.location.href = '/admin/dashboard'
    } else {
      window.location.href = '/user/dashboard'
    }
  } catch (error) {
    alert(error.response?.data?.error || 'Login Gagal. Periksa kembali data Anda.')
  }
}
</script>

<style scoped>
.btn{
  color: #1D2B53;
  background-color: #FF004D;
  border-color: #1D2B53;
}
.btn:hover{
  color: #FF004D;
  background-color: #1D2B53;
  border-color: #FF004D;
}
.link{
  color: #FF004D;
  text-decoration: none;
}
.link:hover{
  color: #1D2B53;
  text-decoration: underline;
}
input::placeholder{
  color: #1D2B53;
}
</style>