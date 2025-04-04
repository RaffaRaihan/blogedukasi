import { ref } from 'vue';
import axios from 'axios';

export default function useFetchUsers() {
  const users = ref([]);
  const loadingUsers = ref(true);

  const fetchUsers = async () => {
    try {
      const token = document.cookie
        .split('; ')
        .find(row => row.startsWith('token='))
        ?.split('=')[1];
      
      const response = await axios.get('http://localhost:8080/admin/users', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      
      users.value = response.data;
    } catch (error) {
      console.error('Error fetching users:', error);
    } finally {
      loadingUsers.value = false;
    }
  };

  return {
    users,
    loadingUsers,
    fetchUsers,
  };
}