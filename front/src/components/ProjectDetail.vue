<template>
  <div style="padding: 20px">

    <el-card style="width: 100%;  ">
      <div style="display: flex;flex-direction: row;justify-content: start;align-items: center">
        <div style="font-size: 18px">近期更新</div>
      </div>
      <div style="display: flex;flex-direction: row;align-items: center;justify-content: space-between">
        <div style="  flex: 1">
          <el-table :data="headerData">
            <el-table-column prop="name" label="名称" width="480"/>
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
    <div style="width:600px;float: right">
      <el-input v-model="searchContent" placeholder="检索" class="input-with-select">
        <template #prepend>
          <el-button :icon="Search"/>
        </template>
      </el-input>
    </div>
    <div style="height: 40px"></div>
    <el-table
        :data="tableData"
        :default-sort="{ prop: 'updateDate', order: 'descending' }"
        style="width: 100%"
    >
      <el-table-column width="60">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-image style="width: 40px; height: 40px" src="/folder/icon-doc.png"></el-image>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="名称" sortable min-width="400">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.name }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="版本" sortable min-width="100">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.version }}</span>
            <template v-if="scope.row.isRelease===false">
              <div class="release">
                发布
              </div>
            </template>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="updateDate" label="更新时间" sortable min-width="180">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.updateDate }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="更新人" sortable min-width="100">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.updateUser }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="追加时间" sortable min-width="180">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <span style="margin-left: 10px">{{ scope.row.createDate }}</span>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="操作" sortable min-width="100">
        <template #default="scope">
          <el-popover trigger="click" placement="right">
            <div style="width: 100%;height: 100%;">
              <div
                  style="display: flex;flex-direction: column;justify-content: space-around;align-items: center;height: 140px">
                <div>
                  <el-button size="small" type="primary" @click="release">发布</el-button>
                </div>
                <div>
                  <el-button size="small" type="primary" @click="fileUpdate">更新</el-button>
                </div>
                <div>
                  <el-button size="small" type="primary">废弃</el-button>
                </div>
                <div>
                  <el-button size="small" type="primary">版本</el-button>
                </div>

              </div>
            </div>
            <template #reference>
              <div>
                <el-button style="width: 40px">操作</el-button>
              </div>
            </template>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <div style="height: 20px"></div>
    <el-pagination background layout="prev, pager, next" :total="1000"/>
  </div>
  <UploadModal @updatedCallBack="updatedCallBack" @closeCallBack="closeCallBack" :is-show="dialogVisible"></UploadModal>
</template>

<script lang="ts" setup>

import {reactive, ref} from "vue";
import {isDark} from "~/composables";
import {Search} from "@element-plus/icons-vue"

interface File {
  updateDate: string
  createDate: string
  name: string
  updateUser: string
  version: string
  isHandlerButtonShow?: boolean
  fileType: string
  isRelease: false
}

const searchContent = ref('')

const updatedCallBack = () => {
  console.log("上传成功")
  dialogVisible.value = false
//todo 更新页面
}
const closeCallBack = () => {
  dialogVisible.value = false
}

const release = () => {


}

const tableData: File[] = reactive(
    []
)
for (let i = 0; i < 15; i++) {
  tableData.push({
    updateDate: '2016-05-03 12:32:55',
    name: '浩天业财融合结算平台接口文档-v17(2)(1)111.docx',
    createDate: '2017-08-03 12:32:51',
    updateUser: '边宇辰',
    version: 'v1.22',
    fileType: "111",
    isRelease: false
  })
}

const headerData: File[] = reactive(
    []
)

for (let i = 0; i < 3; i++) {
  headerData.push({
    updateDate: '2016-05-03 12:32:55',
    name: '浩天业财融合结算平台接口文档-v17(2)(1)111.docx',
    createDate: '2017-08-03 12:32:51',
    updateUser: '边宇辰',
    version: 'v1.22',
    fileType: 'word',
    isRelease: false
  })
}

const dialogVisible = ref(false)
const fileUpdate = () => {
  dialogVisible.value = true
}


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
