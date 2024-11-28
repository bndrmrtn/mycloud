import { defineStore } from 'pinia';
import type { User } from '~/types/user';
import { computed } from 'vue';
import { newRequest } from '~/scripts/request';
import { ca } from 'cronstrue/dist/i18n/locales/ca';

export const useAuthStore = defineStore('auth', () => {
  const userData = ref<User | null>(null);

  const isLoggedIn = computed(() => userData.value != null);
  const user = computed(() => userData.value as User);

  const logout = async () => {
    try {
      await newRequest('/logout');
    } catch (e: unknown) {
      console.error('Failed to logout', e);
    }
    userData.value = null;
  };

  const handle = async () => {
    try {
      const res = await newRequest(`/me`);
      if (res.status != 200) return;
      const data = await res.json();
      userData.value = data as User;
    } catch (e: unknown) {
      console.error(e);
    }
  };

  const set = (data: User | null) => {
    userData.value = data;
  };

  return { set, user, logout, isLoggedIn, handle };
});
