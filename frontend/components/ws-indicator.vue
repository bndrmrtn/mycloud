<script setup lang="ts">
import ModalUtil from '~/components/utils/modal-util.vue';
import PlugIcon from '~/components/icons/plug-icon.vue';

const props = defineProps<{
  active: boolean;
}>();

const show = ref(false);

const msg = computed(
  () => `Realtime connection is ${props.active ? 'active' : 'inactive'}`
);
</script>

<template>
  <button
    @click="show = true"
    v-tooltip="msg"
    class="w-4 h-4 m-5 rounded-full fixed bottom-0 left-0 z-[999] cursor-pointer"
    :class="{ 'bg-green-300': active, 'bg-red-300': !active }"
  ></button>

  <ModalUtil v-if="show" :show="true" @close="show = false">
    <h2 class="text-2xl fredoka">
      Connection is
      <span :class="{ 'text-green-300': active, 'text-red-300': !active }">
        {{ active ? 'active' : 'inactive' }}
      </span>
      <PlugIcon class="-mt-1 ml-1" />
    </h2>

    <p class="text-left mt-2 text-gray-200">
      Our file-sharing service uses WebSocket connections to provide instant
      notifications. This means users receive real-time alerts for uploads,
      shares, and changes, ensuring they stay updated immediately without
      needing to refresh or reload. 🚀
    </p>
  </ModalUtil>
</template>
