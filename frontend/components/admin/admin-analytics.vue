<template>
  <template v-if="d != null">
    <h2 class="fredoka text-2xl">Analytics</h2>
    <section class="md:grid md:grid-cols-2 md:gap-2">
      <div class="p-2 mt-2 rounded-lg bg-widget">
        <diff-chart :diff="d?.file_difference" />
      </div>
      <div class="p-2 mt-2 rounded-lg bg-widget">
        <container-chart :containers="d?.os_file_container" />
      </div>
    </section>
  </template>
</template>

<script setup lang="ts">
import DiffChart from "~/components/admin/diff-chart.vue";
import type {Analytics} from "~/types/analytics";
import {fetchAnalytics} from "~/scripts/analytics";
import {useLoaderStore} from "#imports";
import ContainerChart from "~/components/admin/container-chart.vue";

const loader = useLoaderStore()
const d = ref<Analytics|null>(null)

onMounted(async () => {
  loader.start()
  const res = await fetchAnalytics()
  if(!res) return
  d.value = res
  loader.finish()
})
</script>