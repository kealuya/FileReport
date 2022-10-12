<template>
  <el-row :gutter="40">
    <el-col :span="6">
      <el-card class="card-style2" shadow="hover">
        <el-row class="card-content-height" justify="center" align="middle">
          <el-col :span="8">
            <el-row justify="center">
              <div class="image-chang-wrapper1"></div>
            </el-row>
          </el-col>

          <el-col :span="14">
            <el-row justify="start">
              <div style="font-size: 32px">
                <count-to
                    ref="countToRef1"
                    :start-val="startVal1"
                    :end-val="endVal1"
                    :autoplay="false"
                    suffix=""
                    :decimals=0
                />
                <span style="font-size: 14px; "> 个</span></div>
            </el-row>
            <el-row style="height: 10px;"></el-row>
            <el-row justify="start">
              <div>
                已建项目
              </div>
            </el-row>
          </el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="card-style2" shadow="hover">
        <el-row class="card-content-height" justify="center" align="middle">
          <el-col :span="8">
            <el-row justify="center">
              <div class="image-chang-wrapper2"></div>
            </el-row>
          </el-col>
          <el-col :span="14">
            <el-row justify="start">
              <div style="font-size: 32px">
                <count-to
                    ref="countToRef2"
                    :start-val="startVal2"
                    :end-val="endVal2"
                    :autoplay="false"
                    suffix=""
                    :duration="2000"
                    :decimals=0
                />
                <span style="font-size: 14px; "> 个</span></div>
            </el-row>
            <el-row style="height: 10px;"></el-row>
            <el-row justify="start">
              <div>
                文档总数
              </div>
            </el-row>
          </el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="card-style2" shadow="hover">
        <el-row class="card-content-height" justify="center" align="middle">
          <el-col :span="8">
            <el-row justify="center">
              <div class="image-chang-wrapper3"></div>
            </el-row>
          </el-col>
          <el-col :span="14">
            <el-row justify="start">
              <div style="font-size: 32px">
                <count-to
                    ref="countToRef3"
                    :start-val="startVal3"
                    :end-val="endVal3"
                    :autoplay="false"
                    suffix=""
                    :decimals=0
                />
                <span style="font-size: 14px; "> 位</span></div>
            </el-row>
            <el-row style="height: 10px;"></el-row>
            <el-row justify="start">
              <div>
                维护人员
              </div>
            </el-row>
          </el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card class="card-style2" shadow="hover">
        <el-row class="card-content-height" justify="center" align="middle">
          <el-col :span="8">
            <el-row justify="center">
              <div class="image-chang-wrapper4"></div>
            </el-row>
          </el-col>
          <el-col :span="14">
            <el-row justify="start">
              <div style="font-size: 32px">
                <count-to
                    ref="countToRef4"
                    :start-val="startVal4"
                    :end-val="endVal4"
                    :autoplay="false"
                    suffix=""
                    :duration="3000"
                    :decimals=0
                />
                <span style="font-size: 14px; "> 次</span></div>
            </el-row>
            <el-row style="height: 10px;"></el-row>
            <el-row justify="start">
              <div>
                发布次数
              </div>
            </el-row>
          </el-col>
        </el-row>
      </el-card>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {CountTo} from "~/components/CountTo";
//已建项目
const startVal1 = ref(0)
const endVal1 = ref(5)
const countToRef1 = ref<InstanceType<typeof CountTo>>()
//文档总数
const startVal2 = ref(0)
const endVal2 = ref(214)
const countToRef2 = ref<InstanceType<typeof CountTo>>()
//维护人员
const startVal3 = ref(0)
const endVal3 = ref(15)
const countToRef3 = ref<InstanceType<typeof CountTo>>()
//发布次数
const startVal4 = ref(0)
const endVal4 = ref(1435)
const countToRef4 = ref<InstanceType<typeof CountTo>>()
onMounted(() => {
  setTimeout(() => {
    countToRef1.value.start()
    countToRef2.value.start()
    countToRef3.value.start()
    countToRef4.value.start()
  }, 600)
})

</script>

<style scoped>


.card-style2 {
  height: 120px;
  /*background: -webkit-linear-gradient(*/
  /*    left, #72b5f6, #4ba1f5*/
  /*);*/
}

.card-content-height {
  height: 80px;
}


.image-chang-wrapper1 {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/project_icon_n.png);
  transition: all 0.48s ease-out;
}

.image-chang-wrapper1:hover {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/project_icon_s.png);
}


.image-chang-wrapper2 {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/document_icon_n.png);
  transition: all 0.48s ease-out;
}

.image-chang-wrapper2:hover {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/document_icon_s.png);
}


.image-chang-wrapper3 {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/personnel_icon_n.png);
  transition: all 0.48s ease-out;
}

.image-chang-wrapper3:hover {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/personnel_icon_s.png);
}


.image-chang-wrapper4 {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/issue_icon_n.png);
  transition: all 0.48s ease-out;
}

.image-chang-wrapper4:hover {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-image: url(/issue_icon_s.png);
}


</style>
