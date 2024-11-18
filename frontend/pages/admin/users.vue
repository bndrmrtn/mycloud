<script setup lang="ts">
import {definePageMeta, useLoaderStore} from "#imports";
import SpaceLayout from "~/layouts/space.vue";
import type {User} from "~/types/user";
import {fetchUsers} from "~/scripts/service";
import {useToast} from "vue-toastification";
import {useRouter} from "#app";
import ButtonBluish from "~/components/buttons/button-bluish.vue";

definePageMeta({
  middleware: ['auth', 'admin']
})

const env = useRuntimeConfig()
const router = useRouter()
const users = ref<Array<User>>([])
const cursors = ref({
  next_cursor: "",
  prev_cursor: ""
})

onMounted(async () => {
  const data = await fetchUsers()
  useLoaderStore().loading = false

  if(data instanceof Error) {
    if(process.client) useToast().error(data.message)
    await router.push('/admin')
    return
  }

  users.value = data.data
  cursors.value = {
    next_cursor: data.next_cursor,
    prev_cursor: data.prev_cursor
  }
})
</script>

<template>
  <SpaceLayout>
    <h1 class="fredoka text-3xl mb-5">Manage users</h1>
    <section class="mt-3 md:grid md:grid-cols-2 md:gap-3">
      <div>
        <h2 class="text-2xl fredoka mb-3">Create/Edit</h2>
        <form action=""></form>
      </div>
      <ul class="mt-3 md:mt-0">
        <li>
          <h2 class="text-2xl fredoka mb-3">Users</h2>
        </li>
        <li v-for="user in users" :key="user.id"
          class="bg-widget px-3 py-1.5 rounded-lg flex items-center justify-between"
        >
          <div class="flex items-center space-x-2">
            <img class="w-11 h-11 rounded-lg" alt="User image" :src="`${env.public.api}/profileimage/${user.image_url}`" />
            <div>
              <p>{{ user.name }}</p>
              <p class="text-sm">{{ user.email }}</p>
            </div>
          </div>
          <div>
            <ButtonBluish>Edit</ButtonBluish>
          </div>
        </li>
      </ul>
    </section>
  </SpaceLayout>
</template>
