<template>
  <div>
    <el-dialog
        v-model="isDialogShow"
        title="上传文件"
        width="40%"
        :before-close="handleClose"
    >
      <el-upload
          ref="uploadRef"
          class="upload-demo"
          action=""
          :auto-upload="false"
          :on-exceed="handleExceed"
          limit="1"
          :http-request="httpRequest"
      >
        <template #trigger>
          <el-button type="primary">选择文件</el-button>
        </template>

        <el-button class="ml-3" type="success" @click="submitUpload">
          上传
        </el-button>

        <template #tip>
          <div class="el-upload__tip">
            上传的文件会自动版本+0.01，并维持既存文件名称
          </div>
        </template>
      </el-upload>
      <div style="display: flex;align-items: center;justify-content: start">
        <div  >手动版本调整：</div>
        <el-input v-model="motoVersion" size="small" style=";width: 50px" maxlength="5"/>
      </div>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">关闭</el-button>
      </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref, toRef, toRefs, watchEffect} from "vue";
import {ElMessage, genFileId, UploadInstance, UploadProps, UploadRawFile, UploadRequestOptions} from "element-plus";


const ak = 'SNZGBWTDEF0IRJKXJGJF';
const sk = 'W3H3nbgxHU3zDAblqwvTjO18V6X9ZeIexyn7Ter1';
const server = 'obs.cn-north-4.myhuaweicloud.com';
const bucket = 'file-report-store';

const props = defineProps<{ isShow: boolean }>();
const {isShow} = toRefs(props)
// 组件内部状态控制
const isDialogShow = ref(false)

const emit = defineEmits<{
  (e: 'updatedCallBack'): void,
  (e: 'closeCallBack'): void
}>()

watchEffect(() => {
  // 响应参数变化，只有入参变化，才会触发
  isDialogShow.value = isShow.value
})
const handleClose = () => {
  // 清空自己
  uploadRef.value!.clearFiles()
  // 关闭自己
  isDialogShow.value = false
  // 告诉父组件，变更v-modal
  // 如果组件依赖父组件传参且改变状态，入参不可修改，需要通过事件通知父组件修改入参状态
  emit('closeCallBack')
}

const uploadRef = ref<UploadInstance>()

// 替换掉上一个文件
const handleExceed: UploadProps['onExceed'] = (files) => {
  uploadRef.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  uploadRef.value!.handleStart(file)
}
const submitUpload = () => {
  // submit方法出发httpRequest方法执行
  uploadRef.value!.submit()
}

// 实际的上传工作
const httpRequest = (options: UploadRequestOptions) => {
  // @ts-ignore : 确保有该对象
  let obsClient: any = window.document.getObsClient(ak, sk, server);
  console.log(options.file.name)

  // todo
  // 1.获取扩展名，与元扩展名进行比对，不同报错
  // 2.后台获取文件当前版本号等内容
  // 3.将选择的文件名称改成【已经上传文件名称_版本号_扩展名形式】上传，用于版本控制
  // 4.文件上传
  // 5.成功后，将文件信息进行后台同步
  let uploadFileName = options.file.name

  obsClient.putObject({
    Bucket: bucket,
    Key: uploadFileName,
    SourceFile: options.file
  }).then(function (result: any) {
    if (result.CommonMsg.Status < 300) {
      // 清空自己
      uploadRef.value!.clearFiles()
      ElMessage.success('上传 ' + uploadFileName + ' 文件成功')
      // 上传成功
      emit('updatedCallBack')
    }
  });
}

const motoVersion = ref('v1.34')

</script>

<style>


</style>
