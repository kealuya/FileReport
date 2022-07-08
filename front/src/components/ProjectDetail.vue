<template>
  <div style="padding: 20px">

    <el-card class="box-card">

      <div style="width: 100%;display: flex;justify-content: start;align-items: center">
        <div style="font-size: 18px">近期更新</div>
      </div>
      <el-table :data="headerData" style="width: 50%">
        <el-table-column prop="name" label="名称" width="480"/>
        <el-table-column prop="updateDate" label="更新时间" width="180"/>
        <el-table-column prop="updateUser" label="更新人"/>
      </el-table>

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

      <el-table-column label="名称" sortable min-width="480">
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
                  <el-button @click="release" size="small" type="primary">发布</el-button>
                </div>
                <div>
                  <el-button size="small" type="primary">更新</el-button>
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

const release = () => {
  console.log(isDark.value)
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
    fileType: 'word',
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


</style>
