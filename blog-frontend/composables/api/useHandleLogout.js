import { ref } from 'vue';
import Cookies from 'js-cookie';

export default function useHandleLogout() {
  const alertMessage = ref('');
  const alertClass = ref('');

  const handleLogout = () => {
    try {
      Cookies.remove('token');
      alertMessage.value = 'Logout Berhasil!';
      alertClass.value = 'alert alert-success';
      
      setTimeout(() => {
        window.location.href = '/login';
      }, 1500);
    } catch (error) {
      console.error('Terjadi kesalahan saat logout:', error);
      alertMessage.value = 'Logout gagal. Silakan coba lagi.';
      alertClass.value = 'alert alert-danger';
    }
  };

  return {
    alertMessage,
    alertClass,
    handleLogout,
  };
}
