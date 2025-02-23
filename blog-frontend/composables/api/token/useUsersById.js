import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute } from 'vue-router';
import useAuth from '@/composables/api/token/useAuth';

export default function useUsersById() {
  const { getTokenFromCookies } = useAuth();
  const route = useRoute();
  const users = ref('');
  const error = ref(null);

  const fetchUsers = async (id) => {
    try {
      const token = getTokenFromCookies();
      if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

      const response = await axios.get(`http://localhost:8080/user/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      console.log('Respons API:', response.data.data);
      users.value = response.data.data;
    } catch (err) {
      error.value = err.response?.data?.message || err.message;
      console.error('Error fetching users:', err);
    }
  };

  onMounted(() => {
    const userId = route.params.id;
    console.log("User ID dari URL:", userId);
    fetchUsers(userId);
  });

  return {
    users,
    error,
    fetchUsers,
  };
}
