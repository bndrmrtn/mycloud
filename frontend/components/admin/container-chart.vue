<script setup lang="ts">
import type {Containers} from "~/types/analytics";
import { Pie } from 'vue-chartjs'

const props = defineProps<{
  containers: Containers
}>()

const getLightColor = () => {
  let letters = 'BCDEF'.split('');
  let color = '#';
  for (let i = 0; i < 6; i++ ) {
    color += letters[Math.floor(Math.random() * letters.length)];
  }
  return color;
}

const getColors = (n: number): Array<string> => {
  let colors = []
  for(let i = 0; i < n; i++) {
    colors.push(getLightColor())
  }
  return colors
}

const chartData = ref({
  labels: [...props.containers.map(c => c.container)],
  datasets: [
    {
      label: 'Bytes',
      backgroundColor: [...getColors(props.containers.length)],
      data: [...props.containers.map(c => c.size)],
    },
  ],
})
const chartOptions = ref({
  responsive: true,
  maintainAspectRatio: false,
})
</script>

<template>
  <Pie
      :data="chartData"
      :options="chartOptions"
  />
</template>
