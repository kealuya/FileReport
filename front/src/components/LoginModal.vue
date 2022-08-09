<template>
  <div>
    <el-dialog
        v-model="isDialogShow"
        title="社内登录"
        width="40%"
        @close="resetForm(undefined)"
    >
      <el-form
          ref="loginFormRef"
          :model="loginForm"
          status-icon
          :rules="loginRules"
          label-width="120px"
      >
        <el-form-item label="手机号" prop="inputPhoneNumber">
          <el-input v-model="loginForm.inputPhoneNumber" placeholder="请输入手机号">
            <template #prepend>
              <img src="/folder/icon-phone-number.png" style="width: 22px;height: 22px;">
            </template>
            <template #append>
              <template v-if="wait60sForCaptcha==false">
                <el-button @click="getCaptcha(loginFormRef)">获取验证码</el-button>
              </template>
              <template v-else>
                <div style="width: 130px;text-align: center">剩余{{ wait60s }}s</div>
              </template>
            </template>
          </el-input>
        </el-form-item>
        <div style="height: 10px;"></div>
        <el-form-item label="验证码" prop="captcha">
          <el-input style="width: 200px" v-model="loginForm.captcha" placeholder="请输入验证码">
            <template #prepend>
              <img src="/folder/icon-captcha.png" style="width: 22px;height: 22px;">
            </template>
          </el-input>
        </el-form-item>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="isDialogShow = false;resetForm(loginFormRef)">取消</el-button>
          <el-button type="primary" @click="submitForm(loginFormRef)">登录</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {computed, reactive, ref, WritableComputedRef} from "vue";
import {useUserStore} from "~/stores";

// ==============================================================================
// 组件 v-modal 模式
const props = defineProps<{ modelValue: boolean }>();
const emit = defineEmits<{
  (e: 'update:modelValue', v: boolean): void,
  (e: 'loginSuccess'): void
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
import type {FormInstance} from 'element-plus'

const loginFormRef = ref<FormInstance>()
const loginForm = reactive({
  inputPhoneNumber: '',
  captcha: '',
})

const loginRules = reactive({
  inputPhoneNumber: [{
    validator: (rule: any, value: any, callback: any) => {
      if (value === '') {
        callback(new Error('手机号填了吗？'))
      } else if (!(value as string).match("^1\\d{10}$")) {
        callback(new Error("填的是手机号吗？"))
      } else {
        callback()
      }
    }, trigger: 'blur'
  }],
  captcha: [{
    validator: (rule: any, value: any, callback: any) => {
      if (value === '') {
        callback(new Error('验证码为空!'))
      } else if (!(value as string).match("^\\d{6}$")) {
        callback(new Error("验证码不正确!"))
      } else {
        callback()
      }
    }, trigger: 'blur'
  }],
})

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      // console.log('submit!')
      console.log(loginForm.inputPhoneNumber, loginForm.captcha)
    } else {

      return false
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

//
const userStore = useUserStore()
const getCaptcha = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validateField("inputPhoneNumber", (valid) => {
    if (valid) {
      userStore.getCaptcha(loginForm.inputPhoneNumber)
      wait60sCountDown()
    } else {
      return false
    }
  })
}

const wait60sForCaptcha = ref(false)
const wait60s = ref(60)
const wait60sCountDown = (id?: NodeJS.Timer) => {

  if (wait60s.value > 0) {
    if (!id) {
      wait60sForCaptcha.value = true
      let innerId = setInterval(() => {
        wait60sCountDown(innerId)
      }, 1000)
    } else {
      wait60s.value -= 1
    }
  } else {
    clearInterval(id)
    wait60sForCaptcha.value = false
    wait60s.value = 60
  }
}

</script>

<style>


</style>
