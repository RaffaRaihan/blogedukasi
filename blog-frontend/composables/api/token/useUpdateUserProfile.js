import { ref } from 'vue';
import axios from 'axios';
import useAuth from '@/composables/api/token/useAuth';
import useUsersById from '@/composables/api/token/useUsersById';

export default function useUpdateUserProfile() {
  const { getTokenFromCookies } = useAuth();
  const { users } = useUsersById();
  const fotoFile = ref(null);
  const error = ref(null);
  const alertMessage = ref('');
  const alertClass = ref('');

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
        `http://localhost:8080/user/${users.value.ID}`,
        {
          name: users.value.name,
          email: users.value.email,
          bio: users.value.bio,
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
          `http://localhost:8080/user/${users.value.ID}/foto`,
          formData,
          {
            headers: {
              Authorization: `Bearer ${token}`,
              "Content-Type": "multipart/form-data",
            },
          }
        );
      }

      alertMessage.value = `Profil berhasil diperbaharui!`;
      alertClass.value = 'alert alert-success';
      setTimeout(() => {
        window.location.href =`/user/profile/${users.value.ID}`;
      }, 3000);
    } catch (err) {
      error.value = err.response?.data?.message || err.message;
      console.error("Error updating profile:", err);
      alert("Terjadi kesalahan saat memperbarui profil.");
    }
  };

  return {
    fotoFile,
    users,
    error,
    alertMessage,
    alertClass,
    handleFileUpload,
    updateUserProfile,
  };
}
