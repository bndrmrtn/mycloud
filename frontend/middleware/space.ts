import { newRequest } from '~/scripts/request';
import type { User } from '~/types/user';
import { validate, version } from 'uuid';
import { abortNavigation } from '#app';

export default defineNuxtRouteMiddleware(async (to) => {
  const id = (to.params?.id as string) || null;

  if (!id) return;

  if (validate(id) && version(id) == 4) return;

  return abortNavigation({
    statusCode: 404,
    statusMessage: 'Invalid space ID',
  });
});
