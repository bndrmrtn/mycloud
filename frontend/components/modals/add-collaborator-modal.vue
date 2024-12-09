<template>
  <ModalUtil :show="true" @close="emit('close')" :disable-close="processing">
    <h2 class="text-2xl fredoka">Add new collaborator</h2>

    <form class="mt-5 mb-2" @submit.prevent>
      <InputUtil
          :disabled="processing"
          v-model="email"
          placeholder="Enter collaborator's email"
      />

      <div class="flex items-center space-x-2 mt-3 mx-auto w-min">
            <button @click="crud.create = !crud.create" v-tooltip="'Create'" :class="{'!bg-green-400': crud.create}" class="fredoka tile">C</button>
            <button @click="crud.read = !crud.read" v-tooltip="'Read'" :class="{'!bg-green-400': crud.read}" class="fredoka tile">R</button>
            <button @click="crud.update = !crud.update" v-tooltip="'Update'" :class="{'!bg-green-400': crud.update}" class="fredoka tile">U</button>
            <button @click="crud.delete = !crud.delete" v-tooltip="'Delete'" :class="{'!bg-green-400': crud.delete}" class="fredoka tile">D</button>
          </div>

      <ButtonGreenish
          @click="submit"
          :isLoading="processing"
          class="mt-3 py-2.5"
      >Add</ButtonGreenish
      >
    </form>
  </ModalUtil>
</template>

<script setup lang="ts">
import ModalUtil from '~/components/utils/modal-util.vue';
import InputUtil from '~/components/utils/input-util.vue';
import ButtonGreenish from '~/components/buttons/button-greenish.vue';
import { useToast } from 'vue-toastification';
import {putCollaborator} from "~/scripts/collaborators";

const route = useRoute()
const email = ref('');
const processing = ref(false);
const crud = ref({create: false, read: false, update: false, delete: false})

const emit = defineEmits(['close', 'updated']);

const submit = async () => {
  processing.value = true;
  const sp = await putCollaborator(route.params.id as string, email.value, crud.value);
  processing.value = false;

  if (sp instanceof Error)
    return process.client ? useToast().error(sp.message) : null;

  email.value = '';
  emit('updated');
  emit('close');
};
</script>

<style scoped>
.tile {
  @apply transition cursor-pointer w-10 h-10 text-lg rounded-lg bg-gray-400 flex items-center justify-center hover:opacity-80
}
</style>
