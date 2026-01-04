<template>
  <div class="chat-wrap">
    <div
      class="chat-window"
      :style="{
        boxShadow: `var(${'--el-box-shadow-dark'})`,
      }"
    >
      <el-container class="chat-window-container">
        <el-aside class="aside-container">
          <NavigationModal></NavigationModal>
          <ContactListModal></ContactListModal>
        </el-aside>
        <el-container class="chat-container">
          <el-header>
            <h2 class="chat-name"></h2>
          </el-header>
          <el-main class="main-container">
            <div class="chat-screen">
              <div class="chat-empty-state">
                <svg viewBox="0 0 64 64" width="80" height="80">
                  <path fill="#E0E0E0" d="M32 8C18.7 8 8 18.7 8 32s10.7 24 24 24 24-10.7 24-24S45.3 8 32 8zm0 42c-9.9 0-18-8.1-18-18s8.1-18 18-18 18 8.1 18 18-8.1 18-18 18z"/>
                  <circle fill="#E0E0E0" cx="24" cy="28" r="3"/>
                  <circle fill="#E0E0E0" cx="40" cy="28" r="3"/>
                  <path fill="#E0E0E0" d="M32 38c-3.3 0-6.3 1.7-8 4.5.8.3 1.6.5 2.5.5 2.5 0 4.8-1 6.5-2.6 1.7 1.6 4 2.6 6.5 2.6.9 0 1.7-.2 2.5-.5-1.7-2.8-4.7-4.5-8-4.5z"/>
                </svg>
                <p class="empty-text">选择一个聊天开始对话</p>
              </div>
            </div>
          </el-main>
        </el-container>
      </el-container>
    </div>
  </div>
</template>

<script>
import { reactive, toRefs, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import Modal from "@/components/Modal.vue";
import SmallModal from "@/components/SmallModal.vue";
import ContactListModal from "@/components/ContactListModal.vue";
import NavigationModal from "@/components/NavigationModal.vue";
import axios from "axios";
import { ElMessage } from "element-plus";
export default {
  name: "ContactList",
  components: {
    Modal,
    SmallModal,
    ContactListModal,
    NavigationModal,
  },
  setup() {
    const router = useRouter();
    const store = useStore();
    const data = reactive({
      chatMessage: "",
      chatName: "",
      userInfo: store.state.userInfo,
      contactSearch: "",
      createGroupReq: {
        owner_id: "",
        name: "",
        notice: "",
        add_mode: null,
        avatar: "",
      },
      isCreateGroupModalVisible: false,
      isApplyContactModalVisible: false,
      isNewContactModalVisible: false,
      contactUserList: [],
      myGroupList: [],
      myJoinedGroupList: [],
      applyContactReq: {
        owner_id: "",
        contact_id: "",
        message: "",
      },
      ownListReq: {
        owner_id: "",
      },
      newContactList: [],
      applyContent: "",
    });

    const handleCreateGroup = async () => {
      try {
        data.createGroupReq.owner_id = data.userInfo.uuid;
        const response = await axios.post(
          store.state.backendUrl + "/group/createGroup",
          data.createGroupReq
        );
      } catch (error) {
        console.error(error);
      }
    };
    const showCreateGroupModal = () => {
      data.isCreateGroupModalVisible = true;
    };
    const quitCreateGroupModal = () => {
      data.isCreateGroupModalVisible = false;
    };
    const closeCreateGroupModal = () => {
      if (data.createGroupReq.name == "") {
        ElMessage("请输入群聊名称");
        return;
      }
      if (data.createGroupReq.add_mode == null) {
        ElMessage("请选择加群方式");
        return;
      }
      data.isCreateGroupModalVisible = false;
      handleCreateGroup();
    };
    const showApplyContactModal = () => {
      data.isApplyContactModalVisible = true;
    };
    const quitApplyContactModal = () => {
      data.isApplyContactModalVisible = false;
    };
    const closeApplyContactModal = () => {
      if (data.applyContactReq.contact_id == "") {
        ElMessage.error("请输入申请用户/群组id");
        return;
      }
      if (data.applyContactReq.contact_id[0] == 'G') {
        handleApplyGroup();
      } else {
        handleApplyContact();
      }
    };

    const showNewContactModal = () => {
      handleNewContactList();
    };

    const quitNewContactModal = () => {
      data.isNewContactModalVisible = false;
      data.newContactList = [];
    };
    const handleApplyGroup = async () => {
      try {
        let req = {
          group_id: data.applyContactReq.contact_id,
        }
        let rsp = await axios.post(store.state.backendUrl + "/group/checkGroupAddMode", req);
        if (rsp.data.code == 200) {
          if (rsp.data.data == 0) { // 直接加入
            handleEnterDirectly(data.applyContactReq.contact_id);
            return;
          }
        } else {
          ElMessage.error("申请失败");
          return;
        }
        data.applyContactReq.owner_id = data.userInfo.uuid;
        rsp = await axios.post(
          store.state.backendUrl + "/contact/applyContact",
          data.applyContactReq
        );
        console.log(rsp);
        if (rsp.data.code == 200) {
          if (rsp.data.message == "申请成功") {
            data.isApplyContactModalVisible = false;
            ElMessage.success("申请成功");
            return;
          } else {
            ElMessage.error(rsp.data.message);
          }
        } else {
          ElMessage.error("申请失败");
        }
      } catch (error) {
        console.error(error);
      }
    };
    const handleApplyContact = async () => {
      try {
        data.applyContactReq.owner_id = data.userInfo.uuid;
        const rsp = await axios.post(
          store.state.backendUrl + "/contact/applyContact",
          data.applyContactReq
        );
        console.log(rsp);
        if (rsp.data.code == 200) {
          if (rsp.data.message == "申请成功") {
            data.isApplyContactModalVisible = false;
            ElMessage.success("申请成功");
            return;
          }
        } 
        ElMessage.error(rsp.data.message);
      } catch (error) {
        console.error(error);
      }
    };
    const handleShowUserList = async () => {
      try {
        data.ownListReq.owner_id = data.userInfo.uuid;
        const getUserListRsp = await axios.post(
          store.state.backendUrl + "/contact/getUserList",
          data.ownListReq
        );
        data.contactUserList = getUserListRsp.data.data;
      } catch (error) {
        console.error(error);
      }
    };
    const handleHideUserList = () => {
      data.contactUserList = [];
    };

    const handleShowMyGroupList = async () => {
      try {
        data.ownListReq.owner_id = data.userInfo.uuid;
        const loadMyGroupRsp = await axios.post(
          store.state.backendUrl + "/group/loadMyGroup",
          data.ownListReq
        );
        data.myGroupList = loadMyGroupRsp.data.data;
      } catch (error) {
        console.error(error);
      }
    };
    const handleHideMyGroupList = () => {
      data.myGroupList = [];
    };
    const handleShowMyJoinedGroupList = async () => {
      try {
        data.ownListReq.owner_id = data.userInfo.uuid;
        const loadMyJoinedGroupRsp = await axios.post(
          store.state.backendUrl + "/contact/loadMyJoinedGroup",
          data.ownListReq
        );
        data.myJoinedGroupList = loadMyJoinedGroupRsp.data.data;
      } catch (error) {
        console.error(error);
      }
    };
    const handleHideMyJoinedGroupList = () => {
      data.myJoinedGroupList = [];
    };


    const handleToChatUser = async (user) => {
      try {
        const req = {
          send_id: data.userInfo.uuid,
          receive_id: user.user_id,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/session/checkOpenSessionAllowed",
          req
        );
        if (rsp.data.code == 200) {
          if (rsp.data.data == true) {
            router.push("/chat/" + user.user_id);
          } else {
            ElMessage.warning(rsp.data.message);
            console.error(rsp.data.message);
          }
        } else {
          ElMessage.error(rsp.data.message);
          console.error(rsp.data.message);
        }
      } catch (error) {
        ElMessage.error(error);
        console.error(error);
      }
    };

    const handleToChatGroup = async (group) => {
      try {
        const req = {
          send_id: data.userInfo.uuid,
          receive_id: group.group_id,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/session/checkOpenSessionAllowed",
          req
        );
        if (rsp.data.code == 200) {
          if (rsp.data.data == true) {
            router.push("/chat/" + group.group_id);
          } else {
            ElMessage.warning(rsp.data.message);
            console.error(rsp.data.message);
          }
        } else {
          if (rsp.data.code == 400) {
            ElMessage.warning(rsp.data.message);
            console.error(rsp.data.message);
          } else {
            ElMessage.error(rsp.data.message);
            console.error(rsp.data.message);
          }
        }
      } catch (error) {
        console.error(error);
      }
    };

    const handleNewContactList = async () => {
      try {
        data.ownListReq.owner_id = data.userInfo.uuid;
        const rsp = await axios.post(
          store.state.backendUrl + "/contact/getNewContactList",
          data.ownListReq
        );
        console.log(rsp);
        data.newContactList = rsp.data.data;
        if (data.newContactList == null) {
          ElMessage.warning("没有新的好友申请");
          return;
        }
        data.isNewContactModalVisible = true;
        console.log(rsp);
      } catch (error) {
        console.error(error);
      }
    };

    const handleAgree = async (contactId) => {
      try {
        const req = {
          owner_id: data.userInfo.uuid,
          contact_id: contactId,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/contact/passContactApply",
          req
        );
        console.log(rsp);
        if (rsp.data.code == 200) {
          ElMessage.success(rsp.data.message);
          data.newContactList = data.newContactList.filter(
            (c) => c.contact_id !== contactId
          );
        } else {
          ElMessage.error(rsp.data.message);
        }
      } catch (error) {
        console.error(error);
      }
    };
    const handleEnterDirectly = async (groupId) => {
      try {
        const req = {
          owner_id:  groupId,
          contact_id: data.userInfo.uuid,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/group/enterGroupDirectly",
          req
        );
        console.log(rsp);
        if (rsp.data.code == 200) {
          ElMessage.success(rsp.data.message);
          data.isApplyContactModalVisible = false;
        } else {
          ElMessage.error(rsp.data.message);
        }
      } catch (error) {
        console.error(error);
      }
    };
    const handleReject = async (contactId) => {
      try {
        const req = {
          owner_id: data.userInfo.uuid,
          contact_id: contactId,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/contact/refuseContactApply",
          req
        );
        console.log(rsp);
        if (rsp.data.code == 200) {
          ElMessage.success(rsp.data.message);
          console.log(rsp.data.message);
          data.newContactList = data.newContactList.filter(
            (c) => c.contact_id !== contactId
          );
        } else if (rsp.data.code == 400) {
          ElMessage.warning(rsp.data.message);
          console.log(rsp.data.message);
        } else if (rsp.data.code == 500) {
          ElMessage.error(rsp.data.message);
          console.log(rsp.data.message);
        }
      } catch (error) {
        console.error(error);
      }
    }
    const handleBlack = async (contactId) => {
      try {
        const req = {
          owner_id: data.userInfo.uuid,
          contact_id: contactId,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/contact/blackApply",
          req
        );
        if (rsp.data.code == 200) {
          ElMessage.success(rsp.data.message);
          console.log(rsp.data.message);
          data.newContactList = data.newContactList.filter(
            (c) => c.contact_id !== contactId
          );
        } else if (rsp.data.code == 400) {
          ElMessage.warning(rsp.data.message);
          console.log(rsp.data.message);
        } else if (rsp.data.code == 500) {
          ElMessage.error(rsp.data.message);
          console.log(rsp.data.message);
        }
      } catch (error) {
        ElMessage.error(error);
        console.error(error);
      }
    }
    const handleCancelBlack = async (user) => {
      try {
        const req = {
          owner_id: data.userInfo.uuid,
          contact_id: user.user_id,
        };
        const rsp = await axios.post(
          store.state.backendUrl + "/contact/cancelBlackContact",
          req
        );
        if (rsp.data.code == 200) {
          ElMessage.success(rsp.data.message);
          console.log(rsp.data.message);
        } else if (rsp.data.code == 400) {
          ElMessage.warning(rsp.data.message);
          console.log(rsp.data.message);
        } else if (rsp.data.code == 500) {
          ElMessage.error(rsp.data.message);
          console.log(rsp.data.message);
        }
      } catch (error) {
        ElMessage.error(error);
        console.error(error);
      }
    };


    return {
      ...toRefs(data),
      router,
      handleCreateGroup,
      showCreateGroupModal,
      closeCreateGroupModal,
      quitCreateGroupModal,
      showApplyContactModal,
      quitApplyContactModal,
      closeApplyContactModal,
      showNewContactModal,
      quitNewContactModal,
      handleShowUserList,
      handleHideUserList,
      handleShowMyGroupList,
      handleHideMyGroupList,
      handleShowMyJoinedGroupList,
      handleHideMyJoinedGroupList,
      handleToChatUser,
      handleToChatGroup,
      handleNewContactList,
      handleAgree,
      handleReject,
      handleCancelBlack,
      handleBlack,
    };
  },
};
</script>

<style scoped>
.contactlist-header {
  display: flex;
  flex-direction: row;
  margin-top: 10px;
  margin-bottom: 10px;
}

.contact-search-input {
  width: 185px;
  height: 30px;
  margin-left: 5px;
  margin-right: 5px;
}

.contactlist-header-right {
  width: 40px;
  height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.create-group-btn {
  background-color: rgb(252, 210.9, 210.9);
  cursor: pointer;
  border: none;
  height: 100%;
  width: 30px;
  height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 10px;
}

.create-group-icon {
  width: 15px;
  height: 15px;
}

.el-menu {
  background-color: rgb(252, 210.9, 210.9);
  width: 101%;
}

.el-menu-item {
  background-color: rgb(255, 255, 255);
  height: 45px;
}

.contactlist-user-title {
  font-family: Arial, Helvetica, sans-serif;
}

h3 {
  font-family: Arial, Helvetica, sans-serif;
  color: rgb(69, 69, 68);
}

.modal-quit-btn-container {
  height: 30%;
  width: 100%;
  display: flex;
  flex-direction: row-reverse;
}

.modal-quit-btn {
  background-color: rgba(255, 255, 255, 0);
  color: rgb(229, 25, 25);
  padding: 15px;
  border: none;
  cursor: pointer;
  position: fixed;
  justify-content: center;
  align-items: center;
}

.modal-header {
  height: 20%;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  /*background-color:aqua;*/
}

.modal-body {
  height: 55%;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.newcontact-modal-body {
  height: 70%;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.newcontact-modal-footer {
  height: 10%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-footer {
  height: 25%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-header-title {
  height: 70%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.contactlist-avatar {
  width: 30px;
  height: 30px;
  margin-right: 20px;
}

.newcontact-list {
  width: 280px;
  display: flex;
  flex-direction: column;
  align-items: center;
  font-family: Arial, Helvetica, sans-serif;
}

.newcontact-item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  height: 40px;
}

.action-btn {
  background-color: rgb(252, 210.9, 210.9);
  border: none;
  cursor: pointer;
  justify-content: center;
  align-items: center;
  font-family: Arial, Helvetica, sans-serif;
}

.contactlist-user-menu-item {
  justify-content: center;
  align-items: center;
}

.contactlist-user-item {
  width: 221px;
  height: 45px;
  display: flex;
  align-items: center;
  color: rgba(43, 42, 42, 0.893);
}

.contactlist-user-avatar {
  width: 30px;
  height: 30px;
  margin-left: 20px;
  margin-right: 20px;
}

.chat-screen {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chat-empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #8e8e93;
}

.empty-text {
  margin-top: 16px;
  font-size: 17px;
  color: #8e8e93;
}
</style>