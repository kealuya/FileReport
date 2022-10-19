<template>


  <div style="margin: 15px 0px 15px 0px">
    <StateCard/>
  </div>
  <div style="display: flex;justify-content: center">
    <div class="project">
      <template v-for="item in projectList">
        <el-card style="margin: 10px" :body-style="{padding:'20px',width:'200px',height:'240px'}" shadow="hover"
                 @click="gotoDetail(item)">
          <div class="pro-card-display">
            <div style="flex: 3">
              <img style="width: 100%" :src="item.projectIcon"/>
            </div>
            <div style="flex: 1">
              <span style="font-size: 24px">{{ item.projectName }}</span>
              <!--            <div class="bottom">-->
              <!--              <div style="font-size: 12px;color: gray;">-->
              <!--                {{ item.updateDate }}-->
              <!--              </div>-->
              <!--            </div>-->
            </div>
          </div>

        </el-card>
      </template>
    </div>
  </div>
  <div class="table">
    <div style="flex:10;">
      <h3>状态一览</h3>
      <el-table :data="projectNewStateData" size="large" style="width: 100%;height: 100%">
        <el-table-column prop="docName" label="文档"/>
        <el-table-column prop="updateContent" label="更新"/>
        <el-table-column prop="updateDate" label="时间"/>
        <el-table-column prop="fileVersion" label="版本"/>
      </el-table>
    </div>

    <div style="flex:1; "></div>
    <div style="flex:10;">
      <h3>活跃度一览</h3>
      <el-table :data="projectActiveData" size="large" style=" width: 100%;height: 100%">
        <el-table-column prop="projectName" label="项目名称"/>
        <el-table-column prop="releaseCount" label="发布总数"/>
        <el-table-column prop="fileCount" label="文档总数"/>
      </el-table>
    </div>
  </div>
  <div style="height: 150px"></div>

</template>


<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import {ElMessage, ElLoading, ElMessageBox} from 'element-plus'
import {useRouter} from "vue-router";
import {useUserStore} from "~/stores";
import {http} from "~/http";
import {callGetProjectList} from "~/utils/project";

const count = ref(0);
const input = ref("element-plus");


const toast = () => {
  ElMessage.success('Hello')
}

const router = useRouter()

const gotoDetail = (item: DisplayPro) => {
  ElLoading.service({fullscreen: true})
  // router.push({name: 'Detail', params: {projectId: item.projectId}})
  router.push({path: '/detail', query: {projectId: item.projectId}})
}

type DisplayPro = { projectName: string, projectIcon: string, updateDate?: string, projectId?: string }
const projectList = reactive<Array<DisplayPro>>([])
const getProjectList = async () => {

  let res: HttpResponse = await callGetProjectList()
  if (res.success) {
    let proList = res.data as Array<Project>
    proList.map((pro: Project) => {
      projectList.push({projectName: pro.proName, projectIcon: pro.proLogo, projectId: pro.proId})
    })

  } else {
    await ElMessageBox.alert(res.msg, '提示', {
      confirmButtonText: '好的',
      callback: () => {
        // resetForm(loginFormRef.value)
      }
    })
  }

}


onMounted(async () => {

  await getProjectList()

})


const projectNewStateData = reactive<Array<{ updateContent: string, fileName: string, updateDate: string, fileVersion: string, }>>(
    [
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
      {
        updateContent: "追加资产接口",
        fileName: "浩天业财融合结算平台接口文档-v17(2)(1).docx",
        updateDate: "2022-06-22 13:31:12",
        fileVersion: "v1.42"
      },
    ]
)


const projectActiveData = reactive<Array<{ projectName: string, releaseCount: string, fileCount: string }>>(
    [
      {
        projectName: "差旅",
        releaseCount: "1243",
        fileCount: "103",
      },
      {
        projectName: "采购",
        releaseCount: "534",
        fileCount: "39",
      },
      {
        projectName: "业财融合",
        releaseCount: "55",
        fileCount: "93",
      },
      {
        projectName: "物流",
        releaseCount: "34",
        fileCount: "13",
      },
      {
        projectName: "资产",
        releaseCount: "34",
        fileCount: "13",
      },
      {
        projectName: "医疗",
        releaseCount: "34",
        fileCount: "13",
      },
      {
        projectName: "网约车",
        releaseCount: "34",
        fileCount: "13",
      },
      {
        projectName: "公司管理类",
        releaseCount: "34",
        fileCount: "13",
      },
    ]
)


</script>


<style>
.project {
  margin: 15px 0px 15px 0px;
  width: 1400px;
  display: flex;
  justify-content: start;
  flex-flow: wrap;
}

.table {
  width: 100%;
  display: flex;
  justify-content: space-between;
}

.pro-card-display {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>
