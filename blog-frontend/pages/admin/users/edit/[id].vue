<template>
  <div class="container mt-5">
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
      <div class="mb-3">
        <label for="editImage" class="form-label">Foto Profile</label>
        <input
          type="file"
          id="editImage"
          class="form-control"
          @change="handleImageChange"
        />
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
const userId = route.params.id;

const editUser = ref({
  id: null,
  name: '',
  email: '',
  role: '',
});

const profileImage = ref(null);

const fetchUser = async () => {
  try {
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    const response = await axios.get(`http://localhost:8080/admin/users/${userId}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    editUser.value = response.data.data;
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

    await axios.put(`http://localhost:8080/admin/users/${userId}`, {
      name: editUser.value.name,
      email: editUser.value.email,
      role: editUser.value.role,
    }, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (profileImage.value){
      await uploadImage();
    }

    alert('Berhasil Update user');
    router.push('/admin/users');
  } catch (error) {
    console.error('Error updating user:', error);
  }
};

const handleImageChange = (event) => {
  const file = event.target.files[0];
  if (file) {
    profileImage.value = file;
  }
};

const uploadImage = async () => {
  if (!profileImage.value) {
    alert('Silakan pilih gambar terlebih dahulu.');
    return;
  }

  const formData = new FormData();
  formData.append('foto', profileImage.value);

  try {
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    const response = await axios.put(`http://localhost:8080/admin/users/${userId}/foto`, formData, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'multipart/form-data',
      },
    });

    console.log(response.data)
    alert('Gambar berhasil diunggah.');
  } catch (error) {
    console.error('Error uploading image:', error);
    alert('Gagal mengunggah gambar.');
  }
};

onMounted(() => {
  fetchUser();
});
</script>
