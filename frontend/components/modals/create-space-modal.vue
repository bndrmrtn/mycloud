<template>
  <ModalUtil :show="true" @close="emit('close')" :disable-close="processing">
    <h2 class="text-2xl fredoka">Create a new space <PlanetIcon/></h2>

    <form class="mt-5 mb-2" @submit.prevent>
      <InputUtil :disabled="processing" v-model="spaceName" placeholder="Enter space name" />
      <ButtonGreenish @click="submit" :isLoading="processing" class="mt-3 py-2.5">Create</ButtonGreenish>
    </form>
  </ModalUtil>
</template>

<script setup lang="ts">
import ModalUtil from "~/components/utils/modal-util.vue";
import PlanetIcon from "~/components/icons/planet-icon.vue";
import InputUtil from "~/components/utils/input-util.vue";
import ButtonGreenish from "~/components/buttons/button-greenish.vue";
import {space} from "~/scripts/space";
import {useToast} from "vue-toastification";

const spaceName = ref('')
const processing = ref(false)

const emit = defineEmits(['close', 'finish'])

const submit = async () => {
  processing.value = true
  const space = await space(spaceName.value)
  processing.value = false

  if(space instanceof Error) return process.client ? useToast().error(space.message) : null

  spaceName.value = ''
  emit('finish', space)
  emit('close')
}
</script>