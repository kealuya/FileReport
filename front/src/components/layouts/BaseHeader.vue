<template>
  <el-menu @select="handleSelect" class="el-menu-demo" :ellipsis="false" mode="horizontal">
    <el-row justify="center" align="middle">
      <el-col>
        <el-image style="width: 150px; height: 40px" src="https://yyk-app.obs.cn-north-4.myhuaweicloud.com/szht.png"/>
      </el-col>
    </el-row>
    <div class="flex-grow"/>
    <el-row align="middle" @click="searchHandle">
      <el-col>
        <el-row>
          <div style="position: relative;display: flex;justify-content: center;align-items: center">
            <input class="input-rounded" disabled type="text" value="检索"
                   style="width: 130px;height: 30px">
            <img src="/search_icon.png" style="width: 22px;height: 22px;position: absolute;left: 10px">
            <img src="/document.png" style="width: 22px;height: 22px;position: absolute;right: 10px">
          </div>
        </el-row>
      </el-col>
    </el-row>
    <div style="width: 20px"></div>
    <el-sub-menu index="1">
      <template #title>
        <el-row>
          <el-col>
            <el-row align="middle">
              <img style="width: 40px;height: 40px"
                   src="https://yyk-app.obs.cn-north-4.myhuaweicloud.com/15922124562-20220527165037315.gif">
              <div style="width: 10px"></div>
              <template v-if="userName==''">
                <div>欢迎访问文档发布系统</div>
              </template>
              <template v-else>
                <div>欢迎🍎🍓🥝🍅🍇{{ userName }}🍎🍓🥝🍅🍇</div>
              </template>
            </el-row>
          </el-col>
        </el-row>
      </template>
      <template v-if="userName==''">
        <el-menu-item style="height: 60px;" @click="login" index="login">社内人员登录</el-menu-item>
      </template>
      <template v-else>
        <el-menu-item style="height: 60px;" @click="logout">退出登录</el-menu-item>
      </template>
    </el-sub-menu>
    <el-menu-item @click="toggleDark()">
      <button class="border-none w-full bg-transparent cursor-pointer" style="height: 60px;">
        <i inline-flex i="dark:ep-moon ep-sunny"/>
      </button>
    </el-menu-item>
  </el-menu>

  <!--  <div class="full-faker" v-show="isShowFaker">-->
  <!--      <div>fdfdf</div>-->
  <!--  </div>-->

  <el-dialog v-model="isShowFaker" title="检索" width="50%" center>
    <div style="padding: 20px 20px 20px 20px">
      <SearchBar/>
    </div>
  </el-dialog>

  <LoginModal v-model="isShowLoginModal"/>

</template>

<script lang="ts" setup>
import {toggleDark} from '~/composables';
import {ref, watchEffect} from 'vue'
import {useUserStore} from "~/stores";
import {storeToRefs} from "pinia";

const {userName} = storeToRefs(useUserStore())

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  console.log(activeIndex.value)
}
const isShowFaker = ref(false)
watchEffect(
    () => {
      console.log(activeIndex.value)
    }
)

const searchHandle = () => {
  isShowFaker.value = !isShowFaker.value
}

const isShowLoginModal = ref(false)
const login = () => {
  isShowLoginModal.value = !isShowLoginModal.value
}

const logout = () => {
  useUserStore().logout()
  window.location.reload()
}


</script>
<style>
.input-rounded {
  font-size: medium;
  color: #696868;
  border-radius: 14px;
  border: 1px solid #ccc;
  text-align: center;
  transition: all 0.48s ease-out;
}

.input-rounded:hover {
  color: #2369e7;
  border: 3px solid #2369e7;
}

.full-faker {
  position: fixed;
  z-index: 9000;
  top: 0;
  left: 0;
  background: rgba(183, 183, 190, 0.8);
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

}
</style>
