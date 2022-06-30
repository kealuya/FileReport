<template>


  <div style="margin: 15px 0px 15px 0px">
    <StateCard/>
  </div>

  <div style="margin: 15px 0px 15px 0px;width: 100%;display: flex;justify-content: center;flex-flow: wrap;">
    <template v-for=" item in projectList ">
      <el-card :body-style="{ padding: '20px',width:'300px' }">
        <img style="width: 100%"
            src="https://yyk-app.obs.cn-north-4.myhuaweicloud.com/yyk_szht1949_1653543339403024.jpg"

        />
        <div style="padding: 14px">
          <span>Yummy hamburger</span>
          <div class="bottom">
            <time class="time">2222</time>
            <el-button text class="button">Operating</el-button>
          </div>
        </div>
      </el-card>
    </template>
  </div>


  <el-pagination background layout="prev, pager, next" :total="1000"/>
  <!-- example components -->
  <el-button @click="toast">El Message</el-button>
  <el-button type="primary" @click="count++">primary count is: {{ count }}</el-button>
  <el-button type="success" @click="count++">count is: {{ count }}</el-button>
  <el-button type="warning" @click="count++">count is: {{ count }}</el-button>
  <el-button type="danger" @click="count++">count is: {{ count }}</el-button>
  <el-button type="info" @click="count++">count is: {{ count }}</el-button>
  <br/>
  <el-input v-model="input" style="width: 200px; margin: 20px"/>
  <el-tag>Tag 1</el-tag>

  <br/>
  <el-date-picker v-model="curDate" type="date" placeholder="Pick a day"></el-date-picker>

</template>


<script setup lang="ts">
import {reactive, ref} from "vue";
import {ElMessage} from 'element-plus'

defineProps<{ msg: string }>();

const count = ref(0);
const input = ref("element-plus");

const curDate = ref('')

const toast = () => {
  ElMessage.success('Hello')
}


const projectList = reactive<Array<{ projectName: string, projectIcon: string }>>([
  {projectName: "差旅", projectIcon: ""},
  {projectName: "采购", projectIcon: ""},
  {projectName: "物流", projectIcon: ""},
  {projectName: "资产", projectIcon: ""},
  {projectName: "网约车", projectIcon: ""},
  {projectName: "业财融合", projectIcon: ""},
  {projectName: "医疗", projectIcon: ""},
  {projectName: "代码规约类", projectIcon: ""},
  {projectName: "公司管理类", projectIcon: ""},
  {projectName: "技能提升类", projectIcon: ""},
])


</script>


<style>
.ep-button {
  margin: 4px;
}
</style>
