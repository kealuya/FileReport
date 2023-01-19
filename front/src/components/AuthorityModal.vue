<template>
  <div>
    <el-dialog
        v-model="isDialogShow"
        title="权限修改"
        width="40%"
    >

      <div>
        <div>是否发布（游客预览）</div>
        <el-radio-group v-model="item.isRelease">
          <el-radio label="true">对外发布</el-radio>
          <el-radio label="false">社内资料</el-radio>
        </el-radio-group>
      </div>

      <div>
        <div>是否不许他人编辑（仅自己更新）</div>
        <el-radio-group v-model="item.isOwnerEdit">
          <el-radio label="true">仅自己更新</el-radio>
          <el-radio label="false">其他人可更新</el-radio>
        </el-radio-group>
      </div>

      <div>
        <div>修改文档名称</div>
        <div style="display: flex;align-items: center;justify-content: start">
          <div>文档名称：</div>
          <el-input v-model="item.docName" size="small" style=";width: 450px"/>
        </div>
      </div>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="isDialogShow = false">取消</el-button>
          <el-button type="primary" @click="changeIt">确认修改</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {computed, WritableComputedRef} from "vue";
import {callAuthorityDoc} from "~/utils/doc";
import {ElMessage} from "element-plus";
import {ElLoading, ElMessageBox} from "element-plus/es";

// ==============================================================================
// 组件 v-modal 模式
const props = defineProps<{ modelValue: boolean, item: DocFile }>();
const emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void,
  (e: 'authoritySuccess'): void
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


const changeIt = () => {
  const loadingFlg = ElLoading.service({
    lock: true,
    text: '正在上传',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  let myFile: DocFile = props.item

  callAuthorityDoc(myFile).then((res: HttpResponse) => {
    if (res.success) {
      // 关闭加载
      loadingFlg.close()
      ElMessage.success('更新成功')
      // 上传成功,考虑只是通知上一个页面刷新...
      emit('authoritySuccess')
      isDialogShow.value = false
    } else {
      ElMessageBox.alert(res.msg, '提示', {
        confirmButtonText: '好的',
        callback: () => {
        }
      })
    }
  })


}

</script>

<style>


</style>
