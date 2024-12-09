<script setup lang="ts">
import type {FileDiff} from "~/types/analytics";
import { Bar } from 'vue-chartjs'

const props = defineProps<{
  diff: FileDiff
}>()

const chartData = ref({
  labels: ['Total', 'Unique'],
  datasets: [
    {
      label: 'Bytes',
      backgroundColor: '#847ef7',
      data: [props.diff.total_file_size, props.diff.unique_file_size],
    },
  ],
})
const chartOptions = ref({
  responsive: true,
  maintainAspectRatio: false,
})
</script>

<template>
  <Bar
    :data="chartData"
    :options="chartOptions"
  />
</template>