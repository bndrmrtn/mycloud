<script setup lang="ts">
import InputUtil from "~/components/utils/input-util.vue";
import type {User} from "~/types/user";

const user = ref<User>({} as User)
</script>

<template>
  <div class="px-3 py-4 rounded-lg bg-widget">
    <h2 class="text-2xl fredoka mb-3">Edit</h2>
    <form @submit.prevent>
      <InputUtil class="mb-2" placeholder="Name (e.g. John Doe)" :value="user.name" />
      <InputUtil class="mb-2" placeholder="Name (e.g. John Doe)" :value="user.name" />
    </form>
  </div>
</template>