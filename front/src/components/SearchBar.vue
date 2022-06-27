<template>
  <div class="mt-4">
    <el-input
        v-model="input"
        placeholder="检索文档"
        class="input-with-select"
        clearable
        style="height: 40px"
    >
      <template #prepend>
        <el-select v-model="select" placeholder="全部" style="width: 100px ; ">
          <el-option style=";height: 40px" label="PPT" value="PPT"/>
          <el-option label="WORD" value="WORD"/>
          <el-option label="全部" value="ALL"/>
        </el-select>
      </template>
      <template #append>
        <el-button :icon="Search"/>
      </template>
    </el-input>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {Search} from "@element-plus/icons-vue"

const input = ref('')
const select = ref('ALL')

</script>

<style>
.input-with-select .el-input-group__prepend {
  background-color: var(--el-fill-color-blank);
}
</style>
