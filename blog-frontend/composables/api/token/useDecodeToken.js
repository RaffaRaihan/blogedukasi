import { ref } from 'vue';
import { jwtDecode } from 'jwt-decode';
import useAuth from '@/composables/api/token/useAuth';

export default function useDecodeToken() {
  const { getTokenFromCookies } = useAuth();
  const userId = ref('');

  // Fungsi untuk mendekode token JWT
  const decodeToken = () => {
    try {
      const token = getTokenFromCookies();
      if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

      const decoded = jwtDecode(token); // Dekode JWT
      userId.value = decoded.user_id; // Ambil user_id dari payload
      console.log(decoded);
    } catch (error) {
      console.error('Error decoding token:', error);
    }
  };

  return {
    userId,
    decodeToken,
  };
}
