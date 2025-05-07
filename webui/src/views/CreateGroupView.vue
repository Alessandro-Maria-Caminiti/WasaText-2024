<template>
  <div class="form-wrapper">
    <div class="form-card">
      <h2 class="form-title">Create a Group</h2>

      <div class="form-group">
        <label for="groupname">Group Name</label>
        <input type="text" id="groupname" v-model="groupname" placeholder="Enter group name" />
      </div>

      <div class="form-group">
        <label for="members">Add Members</label>
        <div class="add-members">
          <input type="text" id="members" v-model="newMember" placeholder="Enter a username" />
          <button class="btn" @click="addMember">Add</button>
        </div>
      </div>

      <div class="form-group">
        <label>Added Members</label>
        <div class="member-list">
          <span class="member-tag" v-for="(member, index) in addedMembers" :key="index">{{ member }}</span>
        </div>
      </div>

      <ErrorMsg v-if="errorMessage" :message="errorMessage" />

      <button class="btn create-btn" @click="createGroup">Create Group</button>
    </div>
  </div>
</template>

<script>
import axios from "@/services/axios";
import ErrorMsg from '../components/ErrorMsg.vue';

export default {
  components: {ErrorMsg},
  data() {
    return {
      addedMembers: [],
      groupname: '',
      newMember:"",
      errorMessage: ''
    };
  },
  methods: {
    async createGroup() {
      try {

        if (!this.groupname.trim()) {
          this.errorMessage = "Group name cannot be empty!";
          return;
        }

        this.errorMessage = "";
        console.log('Groupname:', this.groupname);

        const response = await axios.post("/groups", {
          groupname: this.groupname,
          names: this.addedMembers
        })


        this.$router.push({ path: "/chats", query: { groupname: this.groupname.trim() } });

      } catch(error) {
        
        if (error.response && error.response.data.error) {
          this.errorMessage = error.response.data.error;
        } else {
          this.errorMessage = "An unexpected error occurred.";
        }

      }
    },
    addMember() {
    if (!this.newMember.trim()) {
      this.errorMessage = "Member name cannot be empty!";
      return;
    }
    this.errorMessage = "";
    this.addedMembers.push(this.newMember.trim()); // add member
    this.newMember = ""; // clear input field
  }

  }
};
</script>

<style scoped>
.form-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f9fafb;
  font-family: 'Inter', sans-serif;
}

.form-card {
  background: white;
  padding: 32px;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

.form-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 24px;
  text-align: center;
  color: #1f2937;
}

.form-group {
  margin-bottom: 16px;
  display: flex;
  flex-direction: column;
}

label {
  font-size: 0.9rem;
  margin-bottom: 6px;
  color: #374151;
}

input {
  padding: 10px;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  outline: none;
  font-size: 1rem;
  transition: border-color 0.2s;
}

input:focus {
  border-color: #3b82f6;
}

.add-members {
  display: flex;
  gap: 10px;
  align-items: center;
}

.member-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.member-tag {
  background-color: #e0f2fe;
  color: #0369a1;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 0.85rem;
}

.btn {
  padding: 10px 16px;
  background-color: #3b82f6;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 0.95rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.btn:hover {
  background-color: #2563eb;
}

.create-btn {
  width: 100%;
  margin-top: 20px;
}

</style>
