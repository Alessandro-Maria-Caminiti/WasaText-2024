<template>
  <div class="container">
    <div class="input-group">
      <label for="groupname">Groupname:</label>
      <input id="groupname" v-model="groupname" type="text" placeholder="Enter group name">
      <label for="members">Add Members:</label>

      <div class="add-members">
        <input id="members" v-model="newMember" type="text" placeholder="Enter an username">
        <button @click="addMember"> Add </button>
      </div>

      <label for="memberlist">Added Members:</label>
      <p v-for="(member, index) in addedMembers" id="memberlist" :key="index">
        {{ member }}
      </p>

      <ErrorMsg v-if="errorMessage" :message="errorMessage" />
    </div>
    <button @click="createGroup">Create</button>
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
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 40px;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 4px 24px rgba(0,0,0,0.08);
  padding: 40px 32px;
  max-width: 400px;
  margin-left: auto;
  margin-right: auto;
}

.input-group {
  margin-bottom: 24px;
  display: flex;
  flex-direction: column;
  gap: 18px;
  width: 100%;
}

.add-members {
  display: flex;
  flex-direction: row;
  gap: 12px;
  align-items: center;
}

label {
  font-weight: 600;
  color: #222;
  margin-bottom: 2px;
  letter-spacing: 0.01em;
}

input {
  padding: 10px 14px;
  width: 100%;
  border: 1.5px solid #e0e0e0;
  border-radius: 8px;
  font-size: 16px;
  transition: border 0.2s;
  background: #fafbfc;
  outline: none;
}

input:focus {
  border: 1.5px solid #4f8cff;
  background: #fff;
}

button {
  padding: 10px 24px;
  background: linear-gradient(90deg, #4f8cff 0%, #38d39f 100%);
  color: #fff;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  font-size: 16px;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(79,140,255,0.08);
  transition: background 0.2s, transform 0.1s;
}

button:hover {
  background: linear-gradient(90deg, #38d39f 0%, #4f8cff 100%);
  transform: translateY(-2px) scale(1.03);
}

#memberlist {
  background: #f1f6fa;
  color: #333;
  border-radius: 6px;
  padding: 6px 12px;
  margin: 2px 0;
  font-size: 15px;
  box-shadow: 0 1px 3px rgba(79,140,255,0.04);
}

.error-message {
  color: #e74c3c;
  font-size: 15px;
  margin-top: 6px;
  font-weight: 500;
  letter-spacing: 0.01em;
}
</style>
