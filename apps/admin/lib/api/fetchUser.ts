import { useUser } from '../store/useUser';

export async function fetchUser(): Promise<void> {
  try {
    const response = await fetch('/api/v1/users');
    if (!response.ok) {
      // Redirect to login page if the response is not ok
      window.location.href = '/login';
      return;
    }
    const userData = await response.json();
    
    // Set the user data in the User Store
    const { setUser } = useUser.getState();
    setUser(userData);
  } catch (error) {
    console.error('Error fetching user:', error);
    // Redirect to login page if there's an error
    window.location.href = '/login';
  }
}

