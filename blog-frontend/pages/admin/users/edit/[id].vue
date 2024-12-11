<template>
  <div class="container">
    <h2>Edit User</h2>
    <form @submit.prevent="updateUser">
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input type="text" class="form-control" id="name" v-model="editUser.name" required />
      </div>
      <div class="mb-3">
        <label for="email" class="form-label">Email</label>
        <input type="email" class="form-control" id="email" v-model="editUser.email" required />
      </div>
      <div class="mb-3">
        <label for="role" class="form-label">Role</label>
        <select class="form-select" id="role" v-model="editUser.role" required>
          <option value="user">User</option>
          <option value="admin">Admin</option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary">Save Changes</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import axios from 'axios';

const router = useRouter();
const route = useRoute();

const editUser = ref({
  id: null,
  name: '',
  email: '',
  role: '',
});

const fetchUser = async () => {
  try {
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    const response = await axios.get(`http://localhost:8080/admin/users/${route.params.id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    editUser.value = response.data;
  } catch (error) {
    console.error('Error fetching user:', error);
  }
};

const updateUser = async () => {
  try {
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    await axios.put(`http://localhost:8080/admin/users/${editUser.value.id}`, editUser.value, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    router.push('/'); // Kembali ke halaman utama setelah berhasil
  } catch (error) {
    console.error('Error updating user:', error);
  }
};

onMounted(() => {
  fetchUser();
});
</script>