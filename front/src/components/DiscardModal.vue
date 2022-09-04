<template>
  <div>
    <el-dialog
        v-model="isDialogShow"
        title="废弃文档"
        width="40%"
    >
      <div>
        <span>你确定要废弃 </span><span style="font-weight: bold;color: #ec5c5c">{{ item.name }}({{ item.version }})</span><span> 吗?</span>
      </div>
      <div>
        <span style="font-size: 12px;color: #706e6e">废弃文档后，文件不得发布、更新，他人不可见</span>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="isDialogShow = false">我怂了</el-button>
          <el-button type="primary" @click="discardIt">我就是要废了它</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {computed, WritableComputedRef} from "vue";

// ==============================================================================
// 组件 v-modal 模式
const props = defineProps<{ modelValue: boolean, item: DocFile }>();
const emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void,
  (e: 'discardSuccess'): void
}>()

// 这种方式对组件内的组件更友好，且可操作性比较强
const isDialogShow: WritableComputedRef<boolean> = computed({
  get(): boolean {
    return props.modelValue
  },
  set(value: boolean) {
    emit('update:modelValue', value)
  }
})
// ==============================================================================
const discardIt = () => {

}

</script>

<style>


</style>
