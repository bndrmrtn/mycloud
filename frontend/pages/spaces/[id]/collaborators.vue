<script setup lang="ts">
import SpaceLayout from '~/layouts/space.vue';
import {onMounted, definePageMeta, useLoaderStore, ref} from "#imports";
import type {Collaborator} from "~/types/user";
import {fetchCollaborators, putCollaborator} from "~/scripts/collaborators";
import AddCollaboratorModal from "~/components/modals/add-collaborator-modal.vue";

definePageMeta({
  middleware: ['space', 'auth'],
})

const route = useRoute()
const loader = useLoaderStore()
const addModal = ref(false)

const collaborators = ref<Array<Collaborator>>([])
const timeoutStore = ref<Record<string, NodeJS.Timeout|undefined>>({})

const load = async ()  => {
  loader.start()
  const data = await fetchCollaborators(route.params.id as string)
  if (data) {
    collaborators.value = data
  }
  loader.finish()
}

onMounted(load)

const updateCollab = (collab: Collaborator, key: 'upload_file' | 'read_file' | 'update_file' | 'delete_file', value: boolean) => {
  collab.permission = collab.permission || {}
  collab.permission[key] = value

  const cid = timeoutStore.value[collab.id]
  if(cid) clearTimeout(cid)

  timeoutStore.value[collab.id] = setTimeout( () => {
    timeoutStore.value[collab.id] = undefined

    putCollaborator(route.params.id.toString(), collab.user.email, {
      create: collab.permission.upload_file || false,
      read: collab.permission.read_file || false,
      update: collab.permission.update_file || false,
      delete: collab.permission.delete_file || false,
    })
  }, 700)
}
</script>

<template>
  <AddCollaboratorModal :show="true" v-if="addModal" @close="addModal = false" @updated="load" />

  <SpaceLayout>
    <div class="flex items-center justify-between">
      <h1 class="fredoka text-3xl mb-5">Collaborators</h1>
      <div class="flex items-center space-x-2">
        <buttons-button-pinkle @click="addModal = true" class="!w-min">
          Add
        </buttons-button-pinkle>
        <buttons-button-pinkle :to="`/spaces/${route.params.id as string}`" class="!w-min">
          Files
        </buttons-button-pinkle>
      </div>
    </div>
    <section class="mt-4">
      <ul>
        <li class="text-center fredoka my-2" v-if="collaborators.length < 1">
          You have 0 collaborators
        </li>
        <li
            v-for="collab in collaborators" :key="collab.id"
            class="md:flex md:items-center md:justify-between bg-widget p-3 rounded-lg shadow-md mb-2"
        >
          <div class="flex items-center space-x-2">
            <ProfileImage :user="collab.user" />
            <div>
              <p class="fredoka">{{ collab.user.name }}</p>
              <p class="text-sm text-gray-400 text-ellipsis">{{ collab.user.email }}</p>
            </div>
          </div>
          <div class="flex items-center space-x-2 mt-2 md:mt-0">
            <button @click="updateCollab(collab, 'upload_file', !collab.permission?.upload_file)" v-tooltip="'Create'" :class="{'!bg-green-400': collab.permission?.upload_file}" class="fredoka tile">C</button>
            <button @click="updateCollab(collab, 'read_file', !collab.permission?.read_file)" v-tooltip="'Read'" :class="{'!bg-green-400': collab.permission?.read_file}" class="fredoka tile">R</button>
            <button @click="updateCollab(collab, 'update_file', !collab.permission?.update_file)" v-tooltip="'Update'" :class="{'!bg-green-400': collab.permission?.update_file}" class="fredoka tile">U</button>
            <button @click="updateCollab(collab, 'delete_file', !collab.permission?.delete_file)" v-tooltip="'Delete'" :class="{'!bg-green-400': collab.permission?.delete_file}" class="fredoka tile">D</button>
          </div>
        </li>
      </ul>
    </section>
  </SpaceLayout>
</template>

<style scoped>
.tile {
  @apply transition cursor-pointer w-7 h-7 rounded-lg bg-gray-400 flex items-center justify-center hover:opacity-80
}
</style>