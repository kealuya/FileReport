<template>
  <div>
    <el-input
        v-model="input"
        placeholder="检索文档"
        class="input-with-select"
        clearable
        @input="searchHandler"
    >
      <template #prepend>
        <el-select v-model="select" placeholder="全部" size="large" style="width: 100px;">
          <el-option label="PPT" value="PPT"/>
          <el-option label="WORD" value="WORD"/>
          <el-option label="全部" value="ALL"/>
        </el-select>
      </template>
      <template #append>
        <el-button :icon="Search" style="width: 100px;"/>
      </template>
    </el-input>
    <div style="height: 500px;  overflow-y: auto;">
      <div v-show="loading" v-loading="loading" style="width:  100%;height: 100%;  ">
      </div>
      <div v-show="isShowResult" class="search-result">
        <template v-for="item in searchResultItems">
          <div class="search-result-card">
            <a>
          <span class="search-result-card-content">
          🍗[{{ item.project }}]🍗‍： {{ item.fileName }} 💻 {{ item.version }} 📆 {{
              item.updateDate
            }} 👨‍💻‍ {{ item.updateMan }}
        </span>
            </a>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import {Search} from "@element-plus/icons-vue"

const input = ref<string>('')
const select = ref<string>('ALL')
const isShowResult = ref<boolean>(false)
const searchResultItems = reactive<Array<any>>([])
const searchResultItemsTemp = [
  {
    project: "差旅",
    fileName: "神州浩天差旅服务平台-推广最新版（2021.3.22）.pptx",
    version: "v1.21",
    updateDate: "2022年06月12日",
    updateMan: "李宝库",
    obsUrl: ""
  },

  {
    project: "差旅",
    fileName: "差旅平台--技术方案（详版） - 中南财经政法大学.doc",
    version: "v2.11",
    updateDate: "2021年06月12日",
    updateMan: "李报表",
    obsUrl: ""
  },

  {
    project: "差旅",
    fileName: "浩天聚合差旅报销一体化云服务解决方案_20210425.doc",
    version: "v1.53",
    updateDate: "2022年06月28日",
    updateMan: "李宝库",
    obsUrl: ""
  },

  {
    project: "采购",
    fileName: "神州浩天采购服务平台产品介绍-推广修改(20210323).pptx",
    version: "v1.12",
    updateDate: "2021年06月12日",
    updateMan: "李哥哥",
    obsUrl: ""
  },

  {
    project: "医疗",
    fileName: "浙江大学医学院邵逸夫医院全供应链管理系统技术方案20210614.docx",
    version: "v1.01",
    updateDate: "2021年06月12日",
    updateMan: "李基金",
    obsUrl: ""
  },

  {
    project: "差旅",
    fileName: "差旅平台--技术方案（详版） - 中南财经政法大学.doc",
    version: "v2.11",
    updateDate: "2021年06月12日",
    updateMan: "李报表",
    obsUrl: ""
  },

  {
    project: "差旅",
    fileName: "浩天聚合差旅报销一体化云服务解决方案_20210425.doc",
    version: "v1.53",
    updateDate: "2022年06月28日",
    updateMan: "李宝库",
    obsUrl: ""
  },

  {
    project: "采购",
    fileName: "神州浩天采购服务平台产品介绍-推广修改(20210323).pptx",
    version: "v1.12",
    updateDate: "2021年06月12日",
    updateMan: "李哥哥",
    obsUrl: ""
  },

  {
    project: "医疗",
    fileName: "浙江大学医学院邵逸夫医院全供应链管理系统技术方案20210614.docx",
    version: "v1.01",
    updateDate: "2021年06月12日",
    updateMan: "李基金",
    obsUrl: ""
  },

  {
    project: "差旅",
    fileName: "差旅平台--技术方案（详版） - 中南财经政法大学.doc",
    version: "v2.11",
    updateDate: "2021年06月12日",
    updateMan: "李报表",
    obsUrl: ""
  },

  {
    project: "差旅",
    fileName: "浩天聚合差旅报销一体化云服务解决方案_20210425.doc",
    version: "v1.53",
    updateDate: "2022年06月28日",
    updateMan: "李宝库",
    obsUrl: ""
  },

  {
    project: "采购",
    fileName: "神州浩天采购服务平台产品介绍-推广修改(20210323).pptx",
    version: "v1.12",
    updateDate: "2021年06月12日",
    updateMan: "李哥哥",
    obsUrl: ""
  },

  {
    project: "医疗",
    fileName: "浙江大学医学院邵逸夫医院全供应链管理系统技术方案20210614.docx",
    version: "v1.01",
    updateDate: "2021年06月12日",
    updateMan: "李基金",
    obsUrl: ""
  },

]
const loading = ref(false)

let over: any = undefined
const searchHandler = (value: string | number) => {

  if (value != "") {
    loading.value = true
    isShowResult.value = true
    clearTimeout(over)
    over = setTimeout((() => {
      searchResultItems.length = 0
      searchResultItems.push(...searchResultItemsTemp)
      loading.value = false
    }), 1000)

  }
  if (value == "") {
    loading.value = false
    isShowResult.value = false
    clearTimeout(over)
  }
}

</script>

<style>
.input-with-select .el-input-group__prepend {
  background-color: var(--el-fill-color-blank);
}

.search-result {
  width: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
}

.search-result-card {
  width: 90%;
  height: 50px;
  margin: 10px;
  display: flex;
  align-items: center;
  border-radius: 6px;
  box-shadow: 2px 2px 3px #b4b3b3;
  padding: 0px 20px 0px 20px;
  border: 1px solid #e3e2e2;
}

.search-result-card-content {
  font-size: 18px;
  border-bottom: 1px solid;
}

</style>
