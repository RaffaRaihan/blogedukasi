<template>
  <div class="min-vh-100 d-flex align-items-center justify-content-center bg-light">
    <div class="card shadow border-0" style="width: 100%; max-width: 400px;">
      <div class="card-body">
        <h2 class="text-center fw-bold mb-4" style="color: #1D2B53;">Create your account</h2>
        <form @submit.prevent="handleRegister">
          <!-- Input Full Name -->
          <div class="mb-3">
            <input
              v-model="name"
              type="text"
              class="form-control"
              placeholder="Full name"
              required
            />
          </div>
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
          <!-- Input Confirm Password -->
          <div class="mb-3">
            <input
              v-model="confirmPassword"
              type="password"
              class="form-control"
              placeholder="Confirm password"
              required
            />
          </div>
          <!-- Button Register -->
          <button
            type="submit"
            class="btn w-100"
          >
            SIGN UP
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router'; // Import Vue Router
import axios from 'axios';

// State untuk form
const name = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');

// Inisialisasi router
const router = useRouter();

const handleRegister = async () => {
  // Validasi password
  if (password.value !== confirmPassword.value) {
    alert('Passwords do not match!');
    return;
  }

  try {
    // Kirim data ke backend
    const response = await axios.post('http://localhost:8080/register', {
      name: name.value,
      email: email.value,
      password: password.value,
    });

    // Cek status respons
    if (response.status === 201) {
      alert('Registration Successful!');
      // Reset form
      name.value = '';
      email.value = '';
      password.value = '';
      confirmPassword.value = '';
      // Redirect ke halaman login
      router.push('/login');
    } else {
      alert('Registration failed. Please try again.');
    }
  } catch (error) {
    console.error('Error during registration:', error);
    alert('An error occurred during registration. Please check your input or try again later.');
  }
};
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
</style>