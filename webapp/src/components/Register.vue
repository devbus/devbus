<template>
  <div>
    <el-form :model="form" :rules="rules" ref="form" label-position="top">
      <el-form-item label="Name" prop="name">
        <el-input type="text" v-model="form.name" auto-complete="on"></el-input>
      </el-form-item>
      <el-form-item label="Email" prop="email">
        <el-input type="email" v-model="form.email" auto-complete="on"></el-input>
      </el-form-item>
      <el-form-item label="Password" prop="password">
        <el-input type="password" v-model="form.password" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item label="Confirm Password" prop="confirmPassword">
        <el-input type="password" v-model="form.confirmPassword" auto-complete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button class="center-block" type="primary" @click="onSubmit">Sign Up</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
  import Vue from 'vue'
  import Axios from 'axios'
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
      let self = this
      return {
        form: {
          name: '',
          email: '',
          password: '',
          confirmPassword: ''
        },
        rules: {
          name: {required: true, trigger: 'blur'},
          email: {type: 'email', required: true, message: 'Invalid email', trigger: 'blur'},
          password: [
            {required: true, message: 'Password is required', trigger: 'blur'}
          ],
          confirmPassword: {
            required: true,
            trigger: 'blur',
            validator (rule, value, callback, source, options) {
              let errors = []
              if (self.form.password !== value) {
                errors.push(new Error("Passwords don't match"))
              }
              callback(errors)
            }
          }
        }
      }
    },
    methods: {
      onSubmit () {
        let self = this
        Axios.post('/api/session/register', self.form).then((res) => {
          self.$message('Register successfully!')
        }).catch((res) => {
          self.$message.error('Failed to register!')
        })
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
