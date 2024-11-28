import { newRequest } from '~/scripts/request';
import type { User } from '~/types/user';

export default defineNuxtRouteMiddleware(async (to) => {
  // @ts-ignore
  if (import.meta.server) return;

  if (useAuthStore().user.role != 'admin') return navigateTo('/');
});
