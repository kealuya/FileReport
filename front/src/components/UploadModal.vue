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
          :auto-upload="false"
          :on-exceed="handleExceed"
          :limit="1"
          :http-request="httpRequest"
          :on-change="onChange"
      >
        <template #trigger>
          <el-button type="primary">选择文件</el-button>
        </template>

        <el-button ref="uploadHandler" :disabled="isUploadBtnDisabled" class="ml-3" type="success"
                   @click="submitUpload">
          上传
        </el-button>

      </el-upload>
      <el-divider/>
      <div style="display: flex;align-items: center;justify-content: start">
        <div>文档名称：</div>
        <el-input :disabled="mode==UPLOAD_MODAL_MODE.UPLOAD" v-model="item.docName" size="small" style=";width: 450px"/>
      </div>
      <div style="height: 10px;"></div>
      <div style="display: flex;align-items: center;justify-content: start">
        <div>当前版本：</div>
        <el-input disabled v-model="item.versionShow" size="small" style="margin-right: 20px;width: 80px"
                  maxlength="5"/>
        <template v-if="mode==UPLOAD_MODAL_MODE.UPLOAD">
          <div>更新后版本：</div>
          <el-input v-model="updatedVersion" size="small" style=";width: 80px" maxlength="5"/>
        </template>
      </div>
      <div style="height: 10px;"></div>
      <div style="display: flex;align-items: center;justify-content: start">
        <div>文件类型：</div>
        <el-input disabled v-model="item.docType" size="small" style="width: 80px" maxlength="5"/>
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
import {computed, ref} from "vue";
import {
  ElLoading,
  ElMessage,
  genFileId,
  UploadFile,
  UploadInstance,
  UploadProps,
  UploadRawFile,
  UploadRequestOptions
} from "element-plus";
import {WritableComputedRef} from "@vue/reactivity";
import {UPLOAD_MODAL_MODE} from "~/enum";
import {useUserStore} from "~/stores";
import {callNewDoc} from "~/utils/doc";
import {ElMessageBox} from "element-plus/es";

const ak = 'SNZGBWTDEF0IRJKXJGJF';
const sk = 'W3H3nbgxHU3zDAblqwvTjO18V6X9ZeIexyn7Ter1';
const server = 'obs.cn-north-4.myhuaweicloud.com';
const bucket = 'file-report-store';
// ==============================================================================
// 组件 v-modal 模式
const props = defineProps<{ modelValue: boolean, mode: UPLOAD_MODAL_MODE, item: DocFile }>();
const emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void,
  (e: 'updateSuccess'): void
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
const handleClose = () => {

  // 清空自己
  uploadRef.value!.clearFiles()

  isUploadBtnDisabled.value = true
  updatedVersion.value = ""
  // 关闭自己
  isDialogShow.value = false
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
  // check版本号
  if (props.mode === UPLOAD_MODAL_MODE.UPLOAD && !updatedVersion.value.match("v\\d\\.\\d\\d")) {
    ElMessage.error("版本号不正确(v9.99)")
    return
  }
  // submit方法出发httpRequest方法执行
  uploadRef.value!.submit()
}

// 实际的上传工作
const httpRequest = (options: UploadRequestOptions): any => {
  // @ts-ignore : 确保有该对象
  let obsClient: any = window.document.getObsClient(ak, sk, server);
  console.log(options.file.name)
  const loadingFlg = ElLoading.service({
    lock: true,
    text: '正在上传',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  // todo
  // 1.获取扩展名，与元扩展名进行比对，不同报错
  // 2.后台获取文件当前版本号等内容
  /*
     3.上传文件id规则（ownerid+时间戳+文件类型）
       文档/文件的区别：文档是首次上传时，创建的一个位置，名称可以随时改，但id不变
       之后上传的文件都是这个文档的版本，在obs上保存文件id，在数据库中保存文档id与文件id的关联关系
   */
  // 4.文件上传
  // 5.成功后，将文件信息进行后台同步

  const userStore = useUserStore()

  if (props.mode === UPLOAD_MODAL_MODE.NEW) {
    // 新建文档的场合
    // 按照规则拼接filename
    let fileName = userStore.phoneNumber + "-" + Date.now() + "." + props.item.docType
    //上传
    obsClient.putObject({
      Bucket: bucket,
      Key: fileName,
      SourceFile: options.file
    }).then(function (result: any) {
      if (result.CommonMsg.Status < 300) {
        // 上传成功，后台处理
        let myFile: DocFile = props.item
        myFile.fileName = fileName
        myFile.ownerId = userStore.phoneNumber

        // 其他字段，后端补齐
        callNewDoc(myFile).then((res: HttpResponse) => {
          if (res.success) {
            // 关闭加载
            loadingFlg.close()
            // 清空自己
            handleClose()
            ElMessage.success('上传 ' + props.item.docName + ' 文件成功')
            // 上传成功,考虑只是通知上一个页面刷新...
            emit('updateSuccess')
          } else {
            ElMessageBox.alert(res.msg, '提示', {
              confirmButtonText: '好的',
              callback: () => {
              }
            })
          }
        })
      }
    });
  } else {

  }


}

const isUploadBtnDisabled = ref(true)
const updatedVersion = ref("")
const onChange = (file: UploadFile) => {

  let tempFileName = file.name
  let docType = tempFileName.substring(tempFileName.lastIndexOf(".") + 1)
  // 判断是否是文档文件
  if (docType.match("doc|docx|ppt|pptx|xls|xlsx|pdf")) {
    // 判断是 new 还是 update
    // new
    if (props.mode == UPLOAD_MODAL_MODE.NEW) {
      //  新建的场合
      props.item.docType = docType
      props.item.versionShow = "v1.00"
      props.item.docName = tempFileName.substring(0, tempFileName.lastIndexOf("."))
      isUploadBtnDisabled.value = false
    } else {
      // update 更新的场合需要检查上传文件与源文件类型是否一致
      if (props.item.docType !== docType) {
        isUploadBtnDisabled.value = true
        updatedVersion.value = ""
        ElMessage.error('上传文件类型需与源文档类型一致')
        return
      }
      // v1.00 + 0.01 计算
      updatedVersion.value = "v" + (parseFloat(props.item
          .versionShow.replace("v", "")) + 0.01).toString()
      isUploadBtnDisabled.value = false
      return
    }

  } else {
    // 不是文档文件，报错
    // 清空操作
    uploadRef.value!.clearFiles()
    isUploadBtnDisabled.value = true
    updatedVersion.value = ""
    ElMessage.error('上传文件类型不支持')
    return
  }


}

</script>

<style>


</style>
