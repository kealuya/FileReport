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
import {computed, onMounted, reactive, ref, WritableComputedRef} from "vue";
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
import {ElMessageBox, FormInstance} from 'element-plus'
import {callGetCaptcha, callLoginWithCaptcha} from "~/utils/user";
import "element-plus/es/components/message-box/style/index";

const loginFormRef = ref<FormInstance>()
const loginForm = reactive({
  inputPhoneNumber: '',
  captcha: '',
})
// 表单验证规则
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

// 登录
const userStore = useUserStore()
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      callLoginWithCaptcha(loginForm.inputPhoneNumber, loginForm.captcha).then((res: HttpResponse) => {
        if (res.success) {
          let user = res.data as User
          userStore.loginSuccess(user)
          isDialogShow.value = false;
          resetForm(loginFormRef.value)
        } else {
          ElMessageBox.alert(res.msg, '提示', {
            confirmButtonText: '好的',
            callback: () => {
              // resetForm(loginFormRef.value)
            }
          })
        }
      })
    } else {
      return false
    }
  })
}
// 重置表单
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

// 获取验证码
const getCaptcha = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validateField("inputPhoneNumber", (valid) => {
    if (valid) {
      callGetCaptcha(loginForm.inputPhoneNumber).then((res: HttpResponse) => {
        if (res.success) {
          // 倒计时
          wait60sCountDown()
        } else {
          ElMessageBox.alert(res.msg, '提示', {
            confirmButtonText: '好的',
          })
        }
      })
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
