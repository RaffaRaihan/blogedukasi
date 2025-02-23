export default function useAuth() {
    const getTokenFromCookies = () => {
      const token = document.cookie
        .split('; ')
        .find(row => row.startsWith('token='))
        ?.split('=')[1];
      return token;
    };
  
    return {
      getTokenFromCookies,
    };
  }
  