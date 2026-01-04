<template>
  <div class="login-wrap">
    <div
      class="login-window"
      :style="{
        boxShadow: `var(${'--el-box-shadow-dark'})`,
      }"
    >
      <h2 class="login-item">登录</h2>
      <el-form
        ref="formRef"
        :model="loginData"
        label-width="70px"
        class="demo-dynamic"
      >
        <el-form-item
          prop="telephone"
          label="账号"
          :rules="[
            {
              required: true,
              message: '此项为必填项',
              trigger: 'blur',
            },
          ]"
        >
          <el-input v-model="loginData.telephone" />
        </el-form-item>
        <el-form-item
          prop="sms_code"
          label="验证码"
          :rules="[
            {
              required: true,
              message: '此项为必填项',
              trigger: 'blur',
            },
          ]"
        >
          <el-input v-model="loginData.sms_code" style="max-width: 200px">
            <template #append>
              <el-button
                @click="sendSmsCode"
                style="background-color: #07C160; color: #ffffff"
                >点击发送</el-button
              >
            </template>
          </el-input>
        </el-form-item>
      </el-form>
      <div class="login-button-container">
        <el-button type="primary" class="login-btn" @click="handleSmsLogin"
          >登录</el-button
        >
      </div>

      <div class="go-register-button-container">
        <button class="go-register-btn" @click="handleRegister">注册</button>
        <button class="go-sms-btn" @click="handleLogin">密码登录</button>
      </div>
    </div>
  </div>
</template>

<script>
import { reactive, toRefs } from "vue";
import axios from "axios";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { useStore } from "vuex";
export default {
  name: "smsLogin",
  setup() {
    const data = reactive({
      loginData: {
        telephone: "",
        sms_code: "",
      },
    });
    const router = useRouter();
    const store = useStore();
    const handleSmsLogin = async () => {
      try {
        if (!data.loginData.telephone || !data.loginData.sms_code) {
          ElMessage.error("请填写完整登录信息。");
          return;
        }
        if (!checkTelephoneValid()) {
          ElMessage.error("请输入有效的手机号码。");
          return;
        }
        const response = await axios.post(
          store.state.backendUrl + "/user/smsLogin",
          data.loginData
        );
        console.log(response);
        if (response.data.code == 200) {
          if (response.data.data.status == 1) {
            ElMessage.error("该账号已被封禁，请联系管理员。");
            return;
          }
          try {
            ElMessage.success(response.data.message);
            if (!response.data.data.avatar.startsWith("http")) {
              response.data.data.avatar =
                store.state.backendUrl + response.data.data.avatar;
            }
            store.commit("setUserInfo", response.data.data);
            // 准备创建websocket连接
            const wsUrl =
              store.state.wsUrl + "/wss?client_id=" + response.data.data.uuid;
            console.log(wsUrl);
            store.state.socket = new WebSocket(wsUrl);
            store.state.socket.onopen = () => {
              console.log("WebSocket连接已打开");
            };
            store.state.socket.onmessage = (message) => {
              console.log("收到消息：", message.data);
            };
            store.state.socket.onclose = () => {
              console.log("WebSocket连接已关闭");
            };
            store.state.socket.onerror = () => {
              console.log("WebSocket连接发生错误");
            };
            router.push("/chat/contactlist");
          } catch (error) {
            console.log(error);
          }
        } else {
          ElMessage.error(response.data.message);
        }
      } catch (error) {
        ElMessage.error(error);
      }
    };
    const checkTelephoneValid = () => {
      const regex = /^1[3456789]\d{9}$/;
      return regex.test(data.loginData.telephone);
    };
    const handleRegister = () => {
      router.push("/register");
    };
    const handleLogin = () => {
      router.push("/login");
    };
    const sendSmsCode = async () => {
      if (!data.loginData.telephone) {
        ElMessage.error("请输入手机号码。");
        return;
      }
      if (!checkTelephoneValid()) {
        ElMessage.error("请输入有效的手机号码。");
        return;
      }
      try {
        const req = {
          telephone: data.loginData.telephone,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/user/sendSmsCode",
          req
        );
        console.log(rsp);
        if (rsp.data.code == 200) {
          ElMessage.success(rsp.data.message);
        } else if (rsp.data.code == 400) {
          ElMessage.warning(rsp.data.message);
        } else {
          ElMessage.error(rsp.data.message);
        }
      } catch (error) {
        console.error(error);
      }
    };

    return {
      ...toRefs(data),
      router,
      handleSmsLogin,
      handleLogin,
      handleRegister,
      sendSmsCode,
    };
  },
};
</script>

<style>
.login-wrap {
  height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e8eb 50%, #dfe6e9 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
}

.login-window {
  background-color: rgba(255, 255, 255, 0.95);
  position: relative;
  padding: 40px 50px;
  border-radius: 16px;
  backdrop-filter: blur(20px);
  min-width: 380px;
}

.login-item {
  text-align: center;
  margin-bottom: 30px;
  color: #1c1c1e;
  font-size: 24px;
  font-weight: 600;
}

.login-button-container {
  display: flex;
  justify-content: center;
  margin-top: 24px;
  width: 100%;
}

.login-btn,
.login-btn:hover {
  background-color: #07C160;
  border: none;
  color: #ffffff;
  font-weight: 500;
  height: 44px;
  width: 100%;
  border-radius: 22px;
  font-size: 16px;
  transition: all 0.2s ease;
}

.login-btn:hover {
  background-color: #06ad56;
  transform: translateY(-1px);
}

.go-register-button-container {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}

.go-register-btn,
.go-sms-btn {
  background-color: transparent;
  border: none;
  cursor: pointer;
  color: #07C160;
  font-weight: 500;
  font-size: 14px;
  text-decoration: none;
  transition: all 0.2s ease;
}

.go-register-btn:hover,
.go-sms-btn:hover {
  color: #06ad56;
  text-decoration: underline;
}

.el-form-item {
  margin-bottom: 20px;
}

.el-input__wrapper {
  border-radius: 10px;
  box-shadow: 0 0 0 1px #e5e5e5;
}

.el-input__wrapper:hover,
.el-input__wrapper.is-focus {
  box-shadow: 0 0 0 1px #07C160;
}

.el-alert {
  margin-top: 20px;
}
</style>