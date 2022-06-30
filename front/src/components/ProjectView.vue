<template>

</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {CountTo} from "~/components/CountTo";
//已建项目
const startVal1 = ref(0)
const endVal1 = ref(5)
const countToRef1 = ref<InstanceType<typeof CountTo>>()
//文档总数
const startVal2 = ref(0)
const endVal2 = ref(214)
const countToRef2 = ref<InstanceType<typeof CountTo>>()
//维护人员
const startVal3 = ref(0)
const endVal3 = ref(15)
const countToRef3 = ref<InstanceType<typeof CountTo>>()
//发布次数
const startVal4 = ref(0)
const endVal4 = ref(1435)
const countToRef4 = ref<InstanceType<typeof CountTo>>()
onMounted(() => {
  setTimeout(() => {
    countToRef1.value.start()
    countToRef2.value.start()
    countToRef3.value.start()
    countToRef4.value.start()
  }, 600)
})

</script>

<style scoped>




</style>
