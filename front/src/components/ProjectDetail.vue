<template>
  <div style="padding: 20px">
    <el-card style="width: 100%;">
      <div style="display: flex;flex-direction: row;justify-content: start;align-items: center">
        <div style="font-size: 18px">近期更新</div>
      </div>
      <div style="display: flex;flex-direction: row;align-items: center;justify-content: space-between">
        <div style="  flex: 1">
          <el-table :data="headerData">
            <el-table-column prop="docName" label="名称" width="480"/>
            <el-table-column prop="updateDate" label="更新时间" width="180"/>
            <el-table-column prop="updateUser" label="更新人"/>
          </el-table>
        </div>
        <div style="flex: 1">
          <el-image fit="none" class="imageShow"
                    src="/folder/background-image-files1.jpg"/>
        </div>
      </div>
    </el-card>

    <div style="height: 20px"></div>
    <div style="width:100%;display: flex;justify-content: space-between">
      <div>
        <el-button type="primary" @click="fileUpdate(UPLOAD_MODAL_MODE.NEW)">
          新建文档
          <div style="width: 10px;"></div>
          <el-image style="width: 20px; height: 20px" src="/folder/icon-new-file.png"></el-image>
        </el-button>
      </div>
      <el-input style="width:600px" v-model="searchContent" placeholder="检索" class="input-with-select"
                @change="searchContentChange">
        <template #prepend>
          <el-button :icon="Search"/>
        </template>
      </el-input>
    </div>
    <div style="height: 40px"></div>
    <el-table
        :data="tableData"
        @sort-change="sortChange"
        :default-sort="{prop:'updateDate',order:'descending'}"
        style="width: 100%"
    >
      <el-table-column width="60">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-image style="width: 40px; height: 40px" src="/folder/icon-doc.png"></el-image>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="名称" min-width="400">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.docName }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="版本" min-width="100">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.versionShow }}</span>
            <template v-if="scope.row.isRelease===true">
              <div class="release">
                发布
              </div>
            </template>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="updateDate" label="更新时间" sortable="custom" min-width="140">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ timeChange(scope.row.updateDate) }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="更新人" min-width="100">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.updateUser }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="createDate" label="追加时间" sortable="custom" min-width="140">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ timeChange(scope.row.createDate) }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="所有者" min-width="50">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.owner }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="操作" min-width="120">
        <template #default="scope">
          <el-dropdown>
            <el-button type="primary">
              操作
              <el-icon class="el-icon--right">
                <arrow-down/>
              </el-icon>
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="fileUpdate(UPLOAD_MODAL_MODE.UPLOAD,scope.row)">更新上传</el-dropdown-item>
                <el-dropdown-item @click="fileAuthority(scope.row)">权限调整</el-dropdown-item>
                <el-dropdown-item @click="fileDiscard(scope.row)">废弃文档</el-dropdown-item>
                <el-dropdown-item @click="fileHistory">版本历史</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>
    <div style="height: 20px"></div>
    <el-pagination background layout="prev, pager, next" v-model:currentPage="currentPage" :page-count="pageCount"/>
  </div>
  <UploadModal v-model="uploadModalDialogVisible" :mode="uploadModalMode" :item="selectFile"
               @updateSuccess="updateSuccess"></UploadModal>

  <DiscardModal v-model="discardModalDialogVisible" :item="selectFile"
                @discardSuccess="discardSuccess"></DiscardModal>

  <AuthorityModal v-model="authorityModalDialogVisible" :item="selectFile"
                  @authoritySuccess="authoritySuccess"></AuthorityModal>
</template>
<script lang="ts" setup>

// @ts-ignore
import {Search, ArrowDown} from "@element-plus/icons-vue"
import {onMounted, reactive, ref, watchEffect} from "vue";
import {ElLoading} from "element-plus/es";
import {ElNotification} from "element-plus/es/components/notification/index";
/*
  [Bug Report] ElNotification样式错乱 #5968
  https://github.com/element-plus/element-plus/issues/5968
  https://github.com/element-plus/element-plus/issues/5667
  ElNotification 样式错乱问题
  需要手动引入对应组件的样式，才可以。暂时发现【ElNotification】【ElMessageBox】
 */
import "element-plus/es/components/notification/style/index";
import {UPLOAD_MODAL_MODE} from "~/enum";
import {useRoute} from "vue-router";
import {callDocFileList} from "~/utils/doc";
import {timeChange} from "~/utils/common";

type  ParamObject = { [key: string]: string }

const PAGE_SIZE: number = 5
const tableData: DocFile[] = reactive([])
const pageCount = ref<number>(1)
const currentPage = ref<number>(1)
const projectId: string = useRoute().query.projectId as string
const sortCol: ParamObject = {"updateDate": "descending"}
const searchContent = ref("")
const searchContentObj = reactive<ParamObject>({})

const getDocFileList = async (p: PagingInfo) => {

  let res: HttpResponse = await callDocFileList(p)
  if (res.success) {
    tableData.length = 0
    let data = res.data as { docFiles: Array<DocFile>, count: number }
    tableData.push(...data.docFiles)
    pageCount.value = Math.ceil(data.count / PAGE_SIZE)
  }

}

watchEffect(
    async () => {
      let p: PagingInfo = {
        proId: parseInt(projectId),
        page: currentPage.value,
        pageSize: PAGE_SIZE,
        sortCol: sortCol,
        // search: {}
        search: searchContentObj
      }
      console.log("watchEffect", p)
      await getDocFileList(p)
    }
)

const searchContentChange = async (value: string) => {

  Object.keys(searchContentObj).forEach(key => delete searchContentObj[key])
  if (value !== "") {
    searchContentObj["docName"] = value
    searchContentObj["updateUser"] = value
    searchContentObj["owner"] = value
  }
  currentPage.value = 1
  let p: PagingInfo = {
    proId: parseInt(projectId),
    page: currentPage.value,
    pageSize: PAGE_SIZE,
    sortCol: sortCol,
    search: searchContentObj
  }
  await getDocFileList(p)
}


const sortChange = async (column: any) => {

  Object.keys(sortCol).forEach(key => delete sortCol[key])
  if (column.prop !== null) {
    sortCol[column.prop] = column.order
  }
  currentPage.value = 1
  let p: PagingInfo = {
    proId: parseInt(projectId),
    page: currentPage.value,
    pageSize: PAGE_SIZE,
    sortCol: sortCol,
    search: searchContentObj
  }
  await getDocFileList(p)
}


const headerData: DocFile[] = reactive(
    []
)


const uploadModalDialogVisible = ref(false)
const selectFile = ref<DocFile>()
const uploadModalMode = ref<UPLOAD_MODAL_MODE>()
const fileUpdate = (type: UPLOAD_MODAL_MODE, item?: DocFile) => {
  if (type == UPLOAD_MODAL_MODE.NEW) {
    selectFile.value = {proId: projectId} as DocFile
    uploadModalMode.value = UPLOAD_MODAL_MODE.NEW
    uploadModalDialogVisible.value = true
  } else if (type == UPLOAD_MODAL_MODE.UPLOAD) {
    selectFile.value = {...item!, "proId": projectId}
    uploadModalMode.value = UPLOAD_MODAL_MODE.UPLOAD
    uploadModalDialogVisible.value = true
  }
}
const updateSuccess = () => {
  console.log("更新list")
}

const discardSuccess = () => {
  console.log("删除")
}
const discardModalDialogVisible = ref(false)
const fileDiscard = (item: DocFile) => {
  selectFile.value = {...item, "proId": projectId}
  discardModalDialogVisible.value = true
}

const authoritySuccess = () => {
  console.log("权限")
}
const authorityModalDialogVisible = ref(false)
const fileAuthority = (item: DocFile) => {
  selectFile.value = {...item, "proId": projectId}
  authorityModalDialogVisible.value = true
}

const fileHistory = () => {
  ElNotification({
    title: '版本历史',
    message: '等等，下次再说',
    type: 'warning',
    offset: 300
  })
}

onMounted(async () => {
  const loadingInstance = ElLoading.service({fullscreen: true})

  //
  // // 页面初始化
  // let p: PagingInfo = {
  //   proId: parseInt(projectId),
  //   page: 1,
  //   pageSize: PAGE_SIZE,
  //   sortCol: sortCol,
  //   // search: {}
  //   search: {}
  // }
  // await getDocFileList(p)

  loadingInstance.close()
})

</script>
<style>

.record {
  box-shadow: 0px 0px 0px #b4b3b3;
  padding: 0px 20px 0px 20px;
  border: 1px solid #e3e2e2;
  transition: all 0.48s ease-out;
}

.record:hover {
  box-shadow: 2px 2px 3px #b4b3b3;
  padding: 0px 20px 0px 20px;
  border: 1px solid #e3e2e2;
}

.release {
  width: 32px;
  height: 24px;
  background-color: var(--ep-color-primary);
  color: white;
  border-radius: 4px;
  text-align: center;
  font-size: 12px;
  margin-left: 4px;
}

.imageShow {
  padding-left: 20px;
  height: 200px;
  width: 98%;
  border-radius: 8px;
  filter: blur(4px);
}


</style>
