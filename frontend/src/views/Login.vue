<template>
  <div class="message">
    <el-card shadow="hover" class="login-card">
      <el-form :model="loginForm" :rules="loginRules" ref="loginForm" label-position="top" label-width="20rem">
        <el-form-item label="Username" prop="username">
          <el-input v-model="loginForm.username" class="loginInput" />
        </el-form-item>
        <el-form-item label="Password" prop="pass">
          <el-input type="password" v-model="loginForm.pass" class="loginInput" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('loginForm')">Login</el-button>
          <el-button @click="resetForm('loginForm')">Reset</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
/* eslint-disable */
export default {
  data() {
    var validateUser = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Empty username'))
      } else {
        callback()
      }
    }
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Empty password'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        username: '',
        pass: ''
      },
      loginRules: {
        username: [{ validator: validateUser, trigger: 'blur' }],
        pass: [{ validator: validatePass, trigger: 'blur' }]
      }
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          alert('submit!')
        } else {
          console.log('error submit!')
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    }
  }
}
</script>

<style scoped>
.message {
  padding: 5rem 10rem;
  display: flex;
  align-items: center;
  justify-content: center;
}
.login-card {
  width: 30rem;
  font-family: Raleway;
}
.loginInput {
  margin: 0;
  padding: 0;
}
</style>
