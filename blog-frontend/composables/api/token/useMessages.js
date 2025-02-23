import { ref } from 'vue';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';
import useAuth from '@/composables/api/token/useAuth';

export default function useMessages() {
  const { getTokenFromCookies } = useAuth();
  const messages = ref([]);

  const fetchMessage = async () => {
    try {
      const token = getTokenFromCookies();
      if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

      // Decode token untuk mendapatkan userId
      const decoded = jwtDecode(token);
      const userId = decoded.user_id; // Pastikan nama field sesuai dengan isi payload token Anda

      const response = await axios.get(`http://localhost:8080/user/messages/${userId}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      console.log(response.data);
      messages.value = response.data.data;
    } catch (error) {
      console.error('Error fetching messages:', error);
    }
  };

  return {
    messages,
    fetchMessage,
  };
}
