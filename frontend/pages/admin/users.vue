<script setup lang="ts">
import { definePageMeta, useLoaderStore } from '#imports';
import SpaceLayout from '~/layouts/space.vue';
import type { User } from '~/types/user';
import { deleteUser, fetchUsers } from '~/scripts/service';
import { useToast } from 'vue-toastification';
import { useRouter } from '#app';
import ButtonBluish from '~/components/buttons/button-bluish.vue';
import ShieldIcon from '~/components/icons/shield-icon.vue';
import moment from 'moment';
import ButtonDanger from '~/components/buttons/button-danger.vue';

definePageMeta({
  middleware: ['auth', 'admin'],
});

const props = defineProps<{
  isAdmin?: boolean;
}>();

const env = useRuntimeConfig();
const router = useRouter();
const route = useRoute();
const users = ref<Array<User>>([]);
const cursors = ref({
  next_cursor: '',
  prev_cursor: '',
});

const fetchData = async (cursor: string = '') => {
  route.query['cursor'] = cursor;

  useLoaderStore().loading = true;
  const data = await fetchUsers(cursor, props?.isAdmin);
  useLoaderStore().loading = false;

  if (data instanceof Error) {
    if (process.client) useToast().error(data.message);
    await router.push('/admin');
    return;
  }

  users.value = data.data;
  cursors.value = {
    next_cursor: data.next_cursor,
    prev_cursor: data.prev_cursor,
  };
};

onMounted(async () => {
  await fetchData((route.query?.cursor as string | undefined) || '');
});

const formatDate = (s: string): string => {
  return moment(new Date(s)).format('YYYY-MM-DD HH:mm:ss');
};

const deleting = ref(false);
const delUser = async (userID: string) => {
  deleting.value = true;
  const data = await deleteUser(userID);
  deleting.value = false;

  if (data instanceof Error) {
    if (process.client) useToast().warning(data.message);
    return;
  }

  users.value.filter((u) => u.id != userID);
};
</script>

<template>
  <SpaceLayout>
    <h1 class="fredoka text-3xl mb-5">
      Manage {{ isAdmin ? 'admins' : 'users' }}
    </h1>
    <section class="mt-3">
      <ul class="mt-3 max-w-sm mx-auto">
        <li
          v-for="user in users"
          :key="user.id"
          class="bg-widget px-3 py-1.5 rounded-lg drop-shadow-sm"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2">
              <img
                class="w-11 h-11 rounded-xl"
                :src="env.public.api + '/profileimage/' + user.image_url"
                alt=""
              />
              <div>
                <h3 class="text-lg fredoka">{{ user.name }}</h3>
                <p class="text-xs text-gray-400">
                  {{ formatDate(user.created_at) }}
                </p>
              </div>
            </div>
            <ButtonBluish
              :to="'/admin/users/' + user.id"
              class="hidden md:block !w-min"
              >Edit</ButtonBluish
            >
          </div>
          <div class="mt-2">
            <p class="mb-1">
              <span class="fredoka">Email:</span>
              <span
                class="text-gray-200 bg-main-from px-1.5 py-0.5 rounded-lg m-1 text-sm"
                >{{ user.email }}</span
              >
            </p>
            <p class="mb-1">
              <span class="fredoka">Role:</span>
              <span
                class="text-gray-200 bg-main-from px-1.5 py-0.5 rounded-lg m-1 text-sm"
                >{{ user.role }}</span
              >
              <ShieldIcon
                v-if="user.role == 'admin'"
                class="text-yellow-400 -mt-0.5"
              />
            </p>
            <hr class="border-0 h-1 bg-main-from rounded-lg mb-1" />
            <p class="mb-1">
              <span class="fredoka">ID:</span>
              <span
                class="text-gray-200 bg-main-from px-1.5 py-0.5 rounded-lg m-1 text-sm"
                >{{ user.id }}</span
              >
            </p>
            <p class="mb-1">
              <span class="fredoka">Google ID:</span>
              <span
                class="text-gray-200 bg-main-from px-1.5 py-0.5 rounded-lg m-1 text-sm"
                >{{ user.gid }}</span
              >
            </p>
            <hr class="border-0 h-1 bg-main-from rounded-lg mb-3 mt-2" />
            <ButtonBluish
              :to="'/admin/users/' + user.id"
              class="mb-2 md:hidden"
            >
              Edit
            </ButtonBluish>
            <ButtonDanger
              @click="delUser(user.id)"
              :is-loading="deleting"
              class="mb-2"
            >
              Delete
            </ButtonDanger>
          </div>
        </li>
      </ul>

      <div class="max-w-sm mx-auto grid grid-cols-2 gap-2 mb-2 mt-5">
        <ButtonBluish
          :disabled="cursors.prev_cursor === ''"
          @click="fetchData(cursors.prev_cursor)"
        >
          Previous
        </ButtonBluish>
        <ButtonBluish
          :disabled="cursors.next_cursor === ''"
          @click="fetchData(cursors.next_cursor)"
        >
          Next
        </ButtonBluish>
      </div>
    </section>
  </SpaceLayout>
</template>
