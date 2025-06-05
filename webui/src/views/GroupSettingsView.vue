<template>
  <div class="group-settings">
    <h1 class="group-settings-title">Group Settings</h1>
    <div class="group-settings-card">
      <img class="group-photo" :src="groupPhotoURL" alt="Group Avatar">
      <input type="file" accept="image/*" @change="handleFileUpload">
      <button class="upload-button" @click="uploadGroupPhoto">Upload Group Photo</button>
      <input v-model="groupname" class="group-name" type="text">
      <div class="button-group">
        <button class="save-button" @click="saveGroupName">Save</button>
        <button class="add-user-button" @click="addUserToGroup">Add User</button>
        <button class="leave-button" @click="leaveGroup">Leave Group</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "@/services/axios";

export default {
  data() {
    return {
      groupname: this.$route.query.groupname || "Unnamed Group",
      groupPhotoURL: "https://ui-avatars.com/api/?name=Group&size=100",
      imageFile: null,
    };
  },
  async mounted() {
    await this.fetchGroupData();
  },
  methods: {
    async fetchGroupData() {
      try {
        const response = await axios.get(`/groups/${this.groupname}`);
        if (response.data) {
          this.groupPhotoURL = response.data.group_photo_url || this.groupPhotoURL;
        }
      } catch (error) {
        console.error("Error fetching group data:", error);
      }
    },
    handleFileUpload(event) {
      this.imageFile = event.target.files[0];
    },
    async uploadGroupPhoto() {
      if (!this.imageFile) {
        alert("Please select an image first.");
        return;
      }
      try {
        const formData = new FormData();
        formData.append("image", this.imageFile);

        const uploadResponse = await axios.post("/upload", formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });

        const imageUrl = uploadResponse.data.imageUrl;
        console.log("Image uploaded successfully:", imageUrl);
        await axios.put(`/groups/${this.groupname}/group_photo`, {newPhotoURL: imageUrl });

        this.groupPhotoURL = imageUrl;
        alert("Group photo updated!");
      } catch (error) {
        console.error("Error uploading group photo:", error);
        alert("Failed to upload group photo.");
      }
    },
    async saveGroupName() {
      try {
        await axios.put(`/groups/${this.groupname}`, { NewGroupName: this.groupname });
        alert("Group name updated!");
      } catch (error) {
        console.error("Error updating group name:", error);
        alert("Failed to update group name.");
      }
    },
    async addUserToGroup() {
      const newUser = prompt("Enter username to add:");
      if (!newUser) return;
      try {
        await axios.post(`/groups/${this.groupname}`, { Name: newUser });
        alert(`${newUser} added to the group!`);
      } catch (error) {
        console.error("Error adding user to group:", error);
        alert("Failed to add user.");
      }
    },
    async leaveGroup() {
      if (!confirm("Are you sure you want to leave this group?")) return;
      try {
        await axios.delete(`/groups/${this.groupname}`);
        alert("You have left the group.");
        this.$router.push("/chats");
      } catch (error) {
        console.error("Error leaving group:", error);
        alert("Failed to leave group.");
      }
    },
  },
};
</script>

<style scoped>
.group-settings {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  min-height: 100vh;
  background: linear-gradient(135deg, #6a82fb 0%, #fc5c7d 100%);
  font-family: "Inter", "Roboto", Arial, sans-serif;
  color: #22223b;
  padding-top: 60px;
}

.group-settings-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 32px;
  letter-spacing: 1px;
  color: #fff;
  text-shadow: 0 2px 16px rgba(0,0,0,0.12);
}

.group-settings-card {
  background: rgba(255,255,255,0.95);
  border-radius: 18px;
  padding: 48px 40px 36px 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: 0 8px 32px rgba(80, 80, 120, 0.18);
  min-width: 340px;
  max-width: 95vw;
}

.group-photo {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  margin-bottom: 18px;
  object-fit: cover;
  border: 4px solid #6a82fb;
  box-shadow: 0 2px 12px rgba(106,130,251,0.18);
  background: #f3f3f3;
}

input[type="file"] {
  margin-bottom: 12px;
  font-size: 1rem;
  color: #444;
}

.upload-button {
  background: linear-gradient(90deg, #6a82fb 0%, #fc5c7d 100%);
  color: #fff;
  border: none;
  border-radius: 8px;
  padding: 8px 20px;
  font-size: 1rem;
  font-weight: 500;
  margin-bottom: 18px;
  margin-top: 2px;
  cursor: pointer;
  transition: background 0.2s;
}
.upload-button:hover {
  background: linear-gradient(90deg, #fc5c7d 0%, #6a82fb 100%);
}

.group-name {
  font-size: 1.1rem;
  padding: 12px;
  border: 1.5px solid #d1d1e0;
  border-radius: 10px;
  text-align: center;
  width: 260px;
  margin-bottom: 24px;
  background: #f7f7fa;
  color: #22223b;
  transition: border 0.2s;
}
.group-name:focus {
  border: 1.5px solid #6a82fb;
  outline: none;
}

.button-group {
  display: flex;
  gap: 14px;
  margin-top: 10px;
}

button {
  padding: 10px 22px;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.18s, box-shadow 0.18s;
  box-shadow: 0 2px 8px rgba(106,130,251,0.07);
}

.save-button {
  background: linear-gradient(90deg, #21005d 0%, #6a82fb 100%);
  color: #fff;
}
.save-button:hover {
  background: linear-gradient(90deg, #6a82fb 0%, #21005d 100%);
}

.add-user-button {
  background: linear-gradient(90deg, #43e97b 0%, #38f9d7 100%);
  color: #fff;
}
.add-user-button:hover {
  background: linear-gradient(90deg, #38f9d7 0%, #43e97b 100%);
}

.leave-button {
  background: linear-gradient(90deg, #ff5858 0%, #f09819 100%);
  color: #fff;
}
.leave-button:hover {
  background: linear-gradient(90deg, #f09819 0%, #ff5858 100%);
}
</style>
