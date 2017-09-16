<template>
  <div>
    <el-form :model="form" :rules="rules" ref="form">
      <el-form-item prop="email">
        <el-input type="email" v-model="form.email" auto-complete="on"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="form.password" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item class="text-center">
        <el-checkbox label="Remember me" v-model="form.remember"></el-checkbox>
      </el-form-item>
      <el-form-item>
        <el-button class="center-block" type="primary" @click="onSubmit">Sign In</el-button>
      </el-form-item>
      <el-form-item class="text-center">
        <a class="text-center col-xs-12" href="/resetPassword.html">Forgot Password?</a>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import Axios from 'axios'
  import Vue from 'vue'
  import {
    Form,
    FormItem,
    Input,
    Checkbox,
    Button,
    Message
  } from 'element-ui'

  Vue.use(Form)
  Vue.use(FormItem)
  Vue.use(Input)
  Vue.use(Checkbox)
  Vue.use(Button)
  Vue.prototype.$message = Message

  export default {
    name: 'login',
    data () {
      return {
        form: {
          email: '',
          password: '',
          remember: false
        },
        rules: {
          email: [
            {type: 'email', required: true, message: 'Invalid email', trigger: 'blur'}
          ],
          password: [
            {required: true, message: 'Password is required', trigger: 'blur'}
          ]
        }
      }
    },
    methods: {
      onSubmit () {
        let self = this
        Axios.post('/api/session/login', self.form).then((res) => {
          window.location = '/app/main.html'
        }).catch((res) => {
          self.$message.error('Sorry, Email or password error!')
        })
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
